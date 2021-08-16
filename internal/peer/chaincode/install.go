/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chaincode

import (
	"context"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	cb "github.com/SmartBFT-Go/fabric-protos-go/v2/common"
	pb "github.com/SmartBFT-Go/fabric-protos-go/v2/peer"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/core/common/ccpackage"
	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/internal/peer/common"
	"github.com/hyperledger/fabric/internal/pkg/identity"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var chaincodeInstallCmd *cobra.Command

const (
	installCmdName = "install"
)

// Installer holds the dependencies needed to install
// a chaincode
type Installer struct {
	Command         *cobra.Command
	EndorserClients []pb.EndorserClient
	Input           *InstallInput
	Signer          identity.SignerSerializer
	CryptoProvider  bccsp.BCCSP
}

// InstallInput holds the input parameters for installing
// a chaincode
type InstallInput struct {
	Name        string
	Version     string
	Language    string
	PackageFile string
	Path        string
}

// installCmd returns the cobra command for chaincode install
func installCmd(cf *ChaincodeCmdFactory, i *Installer, cryptoProvider bccsp.BCCSP) *cobra.Command {
	chaincodeInstallCmd = &cobra.Command{
		Use:       "install",
		Short:     "Install a chaincode.",
		Long:      "Install a chaincode on a peer. This installs a chaincode deployment spec package (if provided) or packages the specified chaincode before subsequently installing it.",
		ValidArgs: []string{"1"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if i == nil {
				var err error
				if cf == nil {
					cf, err = InitCmdFactory(cmd.Name(), true, false, cryptoProvider)
					if err != nil {
						return err
					}
				}
				i = &Installer{
					Command:         cmd,
					EndorserClients: cf.EndorserClients,
					Signer:          cf.Signer,
					CryptoProvider:  cryptoProvider,
				}
			}
			return i.installChaincode(args)
		},
	}
	flagList := []string{
		"lang",
		"ctor",
		"path",
		"name",
		"version",
		"peerAddresses",
		"tlsRootCertFiles",
		"connectionProfile",
	}
	attachFlags(chaincodeInstallCmd, flagList)

	return chaincodeInstallCmd
}

// installChaincode installs the chaincode
func (i *Installer) installChaincode(args []string) error {
	if i.Command != nil {
		// Parsing of the command line is done so silence cmd usage
		i.Command.SilenceUsage = true
	}

	i.setInput(args)

	// LSCC install
	return i.install()
}

func (i *Installer) setInput(args []string) {
	i.Input = &InstallInput{
		Name:    chaincodeName,
		Version: chaincodeVersion,
		Path:    chaincodePath,
	}

	if len(args) > 0 {
		i.Input.PackageFile = args[0]
	}
}

// install installs a chaincode deployment spec to "peer.address"
// for use with lscc
func (i *Installer) install() error {
	ccPkgMsg, err := i.getChaincodePackageMessage()
	if err != nil {
		return err
	}

	proposal, err := i.createInstallProposal(ccPkgMsg)
	if err != nil {
		return err
	}

	signedProposal, err := protoutil.GetSignedProposal(proposal, i.Signer)
	if err != nil {
		return errors.WithMessagef(err, "error creating signed proposal for %s", chainFuncName)
	}

	return i.submitInstallProposal(signedProposal)
}

func (i *Installer) submitInstallProposal(signedProposal *pb.SignedProposal) error {
	// install is currently only supported for one peer
	proposalResponse, err := i.EndorserClients[0].ProcessProposal(context.Background(), signedProposal)
	if err != nil {
		return errors.WithMessage(err, "error endorsing chaincode install")
	}

	if proposalResponse == nil {
		return errors.New("error during install: received nil proposal response")
	}

	if proposalResponse.Response == nil {
		return errors.New("error during install: received proposal response with nil response")
	}

	if proposalResponse.Response.Status != int32(cb.Status_SUCCESS) {
		return errors.Errorf("install failed with status: %d - %s", proposalResponse.Response.Status, proposalResponse.Response.Message)
	}
	logger.Infof("Installed remotely: %v", proposalResponse)

	return nil
}

func (i *Installer) getChaincodePackageMessage() (proto.Message, error) {
	// if no package provided, create one
	if i.Input.PackageFile == "" {
		if i.Input.Path == common.UndefinedParamValue || i.Input.Version == common.UndefinedParamValue || i.Input.Name == common.UndefinedParamValue {
			return nil, errors.Errorf("must supply value for %s name, path and version parameters", chainFuncName)
		}
		// generate a raw ChaincodeDeploymentSpec
		ccPkgMsg, err := genChaincodeDeploymentSpec(i.Command, i.Input.Name, i.Input.Version)
		if err != nil {
			return nil, err
		}
		return ccPkgMsg, nil
	}

	// read in a package generated by the "package" sub-command (and perhaps signed
	// by multiple owners with the "signpackage" sub-command)
	// var cds *pb.ChaincodeDeploymentSpec
	ccPkgMsg, cds, err := getPackageFromFile(i.Input.PackageFile, i.CryptoProvider)
	if err != nil {
		return nil, err
	}

	// get the chaincode details from cds
	cName := cds.ChaincodeSpec.ChaincodeId.Name
	cVersion := cds.ChaincodeSpec.ChaincodeId.Version

	// if user provided chaincodeName, use it for validation
	if i.Input.Name != "" && i.Input.Name != cName {
		return nil, errors.Errorf("chaincode name %s does not match name %s in package", i.Input.Name, cName)
	}

	// if user provided chaincodeVersion, use it for validation
	if i.Input.Version != "" && i.Input.Version != cVersion {
		return nil, errors.Errorf("chaincode version %s does not match version %s in packages", i.Input.Version, cVersion)
	}

	return ccPkgMsg, nil
}

func (i *Installer) createInstallProposal(msg proto.Message) (*pb.Proposal, error) {
	creator, err := i.Signer.Serialize()
	if err != nil {
		return nil, errors.WithMessage(err, "error serializing identity")
	}

	prop, _, err := protoutil.CreateInstallProposalFromCDS(msg, creator)
	if err != nil {
		return nil, errors.WithMessagef(err, "error creating proposal for %s", chainFuncName)
	}

	return prop, nil
}

// genChaincodeDeploymentSpec creates ChaincodeDeploymentSpec as the package to install
func genChaincodeDeploymentSpec(cmd *cobra.Command, chaincodeName, chaincodeVersion string) (*pb.ChaincodeDeploymentSpec, error) {
	if existed, _ := ccprovider.ChaincodePackageExists(chaincodeName, chaincodeVersion); existed {
		return nil, errors.Errorf("chaincode %s:%s already exists", chaincodeName, chaincodeVersion)
	}

	spec, err := getChaincodeSpec(cmd)
	if err != nil {
		return nil, err
	}

	cds, err := getChaincodeDeploymentSpec(spec, true)
	if err != nil {
		return nil, errors.WithMessagef(err, "error getting chaincode deployment spec for %s", chaincodeName)
	}

	return cds, nil
}

// getPackageFromFile get the chaincode package from file and the extracted ChaincodeDeploymentSpec
func getPackageFromFile(ccPkgFile string, cryptoProvider bccsp.BCCSP) (proto.Message, *pb.ChaincodeDeploymentSpec, error) {
	ccPkgBytes, err := ioutil.ReadFile(ccPkgFile)
	if err != nil {
		return nil, nil, err
	}

	// the bytes should be a valid package (CDS or SignedCDS)
	ccpack, err := ccprovider.GetCCPackage(ccPkgBytes, cryptoProvider)
	if err != nil {
		return nil, nil, err
	}

	// either CDS or Envelope
	o := ccpack.GetPackageObject()

	// try CDS first
	cds, ok := o.(*pb.ChaincodeDeploymentSpec)
	if !ok || cds == nil {
		// try Envelope next
		env, ok := o.(*cb.Envelope)
		if !ok || env == nil {
			return nil, nil, errors.New("error extracting valid chaincode package")
		}

		// this will check for a valid package Envelope
		_, sCDS, err := ccpackage.ExtractSignedCCDepSpec(env)
		if err != nil {
			return nil, nil, errors.WithMessage(err, "error extracting valid signed chaincode package")
		}

		// ...and get the CDS at last
		cds, err = protoutil.UnmarshalChaincodeDeploymentSpec(sCDS.ChaincodeDeploymentSpec)
		if err != nil {
			return nil, nil, errors.WithMessage(err, "error extracting chaincode deployment spec")
		}

		err = platformRegistry.ValidateDeploymentSpec(cds.ChaincodeSpec.Type.String(), cds.CodePackage)
		if err != nil {
			return nil, nil, errors.WithMessage(err, "chaincode deployment spec validation failed")
		}
	}

	return o, cds, nil
}

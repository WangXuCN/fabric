find . -type f \
    -name '*.go' \
    -exec sed -i -e 's,github.com/hyperledger/fabric-protos-go,github.com/SmartBFT-Go/fabric-protos-go/v2,g' {} \;
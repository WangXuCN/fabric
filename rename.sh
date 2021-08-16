find . -type f \
    -name '*.go' \
    -exec sed -i -e 's,github.com/hyperledger/fabric/,github.com/SmartBFT-Go/fabric/,g' {} \;
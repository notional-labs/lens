#!/usr/bin/env bash

set -eox pipefail

echo "Generating gogo proto code"
cd proto

buf generate --template buf.gen.gogo.yaml $file

# move proto files to the right places
#
# Note: Proto files are suffixed with the current binary version.
cp -r ../github.com/strangelove-ventures/lens/* ../
rm -rf ../github.com

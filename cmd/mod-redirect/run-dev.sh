#!/bin/bash

set -euo pipefail

go clean
go generate
go build

./mod-redirect --config ../../hack/example-config.json


#!/bin/bash

set -euo pipefail

go generate
go build

./mod-redirect --config ../../hack/example-config.json


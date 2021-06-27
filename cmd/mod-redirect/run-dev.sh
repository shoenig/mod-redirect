#!/bin/bash

set -euo pipefail

go generate

go run main.go --config ../../hack/example-config.json


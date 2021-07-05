#!/bin/bash

set -euo pipefail

go generate

PORT=9800 go run main.go --config ../../hack/example-config.json

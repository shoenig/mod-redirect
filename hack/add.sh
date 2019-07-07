#!/bin/bash

set -euo pipefail

# e.g. '{"kind":"pkgs", "namespace":"baz/b", "vcs":"git", "destination":"https://a.com/baz/b"}'

if [[ ${#} -ne 4 ]] ; then
	echo "usage: ${0} [kind] [path] [vcs] [source]"
	exit 1
fi

kind="${1}"
ns="${2}"
vcs="${3}"
dest="${4}"

if [[ "${kind}" != "pkgs" && "${kind}" != "cmds" ]] ; then
	echo "kind must be pkgs or cmds"
	exit 1
fi

if [[ "${vcs}" != "git" ]] ; then
	echo "vcs must be git"
	exit 1
fi

echo "creating cURL command to add"
echo "  kind: ${kind}"
echo "  namespace: ${ns}"
echo "  vcs: ${vcs}"
echo "  source: ${dest}"
echo ""
echo ""

header="X-MR-Key"
key=$(cat example-config.json | jq -r .authentication.keys[0])

echo "run this command"
echo ""
echo "  curl -H \"${header}: ${key}\" -XPOST \"localhost:8080/v1/set\" -d '{\"kind\":\"${kind}\", \"namespace\":\"${ns}\", \"vcs\":\"${vcs}\", \"destination\":\"${dest}\"}'"
echo ""
echo ""

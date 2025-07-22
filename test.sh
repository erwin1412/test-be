#!/bin/bash

echo "==> Running unit tests in ./service"
echo

set -e

pushd service > /dev/null

go test -v

popd > /dev/null

echo
echo "==> Service tests PASSED"

#!/usr/bin/env bash

go get -v -d ./...
rm -Rf $GOPATH/src/github.com/hashicorp/hcl
exit 0
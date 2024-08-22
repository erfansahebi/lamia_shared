#!/bin/bash

rm -rf go/protos
mkdir go/protos

protoc --go_out=go/protos --go_opt=Mprotos/auth/auth.proto=github.com/erfansahebi/lamia_auth \
  --go-grpc_out=go/protos --go-grpc_opt=Mprotos/auth/auth.proto=github.com/erfansahebi/lamia_auth \
  protos/auth/auth.proto

# shellcheck disable=SC2164
cd "$(dirname "$0")"

# shellcheck disable=SC2162
find go/protos/github.com/erfansahebi/** -name "*.pb.go" | while read filepath; do
  dir=$(dirname "$filepath")
  parent_dir=$(basename "$dir")
  new_dir=${parent_dir#lamia_}
  mkdir -p go/protos/"$new_dir"
  cp "$filepath" go/protos/"$new_dir"/
done

rm -rf go/protos/github.com
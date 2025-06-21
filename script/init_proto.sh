#!/bin/bash

# Exit immediately if a command fails
set -e

# Optional: Echo each command for debugging
# set -x

# Run protoc to generate Go and gRPC code
protoc -I="../schedule-proto" \
  --go_out=. \
  --go-grpc_out=. \
  "../schedule-proto/auth_service/auth.v1.proto"

protoc -I="../schedule-proto" \
  --go_out=. \
  --go-grpc_out=. \
  "../schedule-proto/auth_service/permission.v1.proto"

protoc -I="../schedule-proto" \
  --go_out=. \
  --go-grpc_out=. \
  "../schedule-proto/auth_service/role.v1.proto"

protoc -I="../schedule-proto" \
  --go_out=. \
  --go-grpc_out=. \
  "../schedule-proto/auth_service/token.v1.proto"

echo "Protobuf files generated successfully."

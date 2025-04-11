#!/bin/bash

# Install dependencies
# brew install fswatch (Mac)
# apt-get install fswatch (Linux)

fswatch -o api/proto/*.proto | while read _; do
  echo "Proto files changed. Regenerating..."
  protoc --proto_path=api/proto \
         --go_out=api/proto/gen \
         --go_gprc_out=api/proto/gen \
         api/proto/*.proto
  #docker-compose restart auth-service product-service api-gateway
done
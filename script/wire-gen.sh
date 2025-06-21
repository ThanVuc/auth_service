#!/bin/bash

set -e

cd $workspace/auth_service/internal/wire
wire

# Check if the wire command was successful
echo "Wire command executed successfully."
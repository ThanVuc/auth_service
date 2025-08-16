#!/usr/bin/env bash

set -ex

GO=go
GOOSE=goose
GOOSE_DIR=sql/schema
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgres://auth_user:auth1234@localhost:5432/auth_db?sslmode=disable"

MODE=$1
ARG=$2

if [[ -z "$MODE" ]]; then
  echo "Usage: $0 {run|create|up|down|reset|status|down-to|up-to} [args]"
  exit 1
fi

case "$MODE" in
  run)
    $GO run cmd/server/main.go
    ;;
  
  create)
    if [[ -z "$ARG" ]]; then
      echo "Error: migration name is required. Usage: $0 create migration_name"
      exit 1
    fi
    $GOOSE -dir $GOOSE_DIR create "$ARG" sql
    ;;
  
  up)
    GOOSE_DRIVER=$GOOSE_DRIVER GOOSE_DBSTRING=$GOOSE_DBSTRING $GOOSE -dir=$GOOSE_DIR up
    ;;
  
  down)
    GOOSE_DRIVER=$GOOSE_DRIVER GOOSE_DBSTRING=$GOOSE_DBSTRING $GOOSE -dir=$GOOSE_DIR down
    ;;
  
  reset)
    GOOSE_DRIVER=$GOOSE_DRIVER GOOSE_DBSTRING=$GOOSE_DBSTRING $GOOSE -dir=$GOOSE_DIR reset
    ;;
  
  status)
    GOOSE_DRIVER=$GOOSE_DRIVER GOOSE_DBSTRING=$GOOSE_DBSTRING $GOOSE -dir=$GOOSE_DIR status
    ;;
  
  down-to)
    if [[ -z "$ARG" ]]; then
      echo "Error: version is required. Usage: $0 down-to VERSION"
      exit 1
    fi
    GOOSE_DRIVER=$GOOSE_DRIVER GOOSE_DBSTRING=$GOOSE_DBSTRING $GOOSE -dir=$GOOSE_DIR down-to "$ARG"
    ;;
  
  up-to)
    if [[ -z "$ARG" ]]; then
      echo "Error: version is required. Usage: $0 up-to VERSION"
      exit 1
    fi
    GOOSE_DRIVER=$GOOSE_DRIVER GOOSE_DBSTRING=$GOOSE_DBSTRING $GOOSE -dir=$GOOSE_DIR up-to "$ARG"
    ;;
  
  *)
    echo "Unknown command: $MODE"
    echo "Usage: $0 {run|create|up|down|reset|status|down-to|up-to} [args]"
    exit 1
    ;;
esac

read -p "Press enter to exit..."
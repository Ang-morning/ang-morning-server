#!/bin/bash

# env
if [ -f .env ]; then
  source .env
fi

# Run the migrate command in local environment
migrate -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" -path ./internal/migrations up
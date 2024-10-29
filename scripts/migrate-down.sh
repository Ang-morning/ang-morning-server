#!/bin/bash

# env local에서는 .env를 통해 환경변수를 설정할 수 있다.
if [ -f .env ]; then
  source .env
fi

# Run the migrate command in local environment
migrate -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" -path ./internal/migrations down

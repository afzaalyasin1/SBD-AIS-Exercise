#!/bin/sh

# 1. Stop and remove any old containers
docker stop pg-db orderservice_app > /dev/null 2>&1
docker rm pg-db orderservice_app > /dev/null 2>&1

# 2. Run the database
docker run --name pg-db \
  -e POSTGRES_DB=order \
  -e POSTGRES_USER=docker \
  -e POSTGRES_PASSWORD=docker \
  -p 5432:5432 \
  -v pg_data:/var/lib/postgresql/data \
  --rm -d postgres:16

# Wait a few seconds for the DB to be ready
echo "Waiting 5s for DB to start..."
sleep 5

docker build -t orderservice .


docker run --name orderservice_app \
  --network host \
  --env-file .env \
  --rm -d orderservice

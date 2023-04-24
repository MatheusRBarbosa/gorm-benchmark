#!/bin/bash
echo "Running Docker"

docker run -p 8080:8080 -d -e PORT=8080 --name=gorm-benchmark \
 -e DB_USER=postgres \
 -e DB_PASSWORD=senha123 \
 -e DB_HOST=0.0.0.0 \
 -e DB_PORT=5432 \
 -e DB_NAME=postgres \
 gorm-benchmark
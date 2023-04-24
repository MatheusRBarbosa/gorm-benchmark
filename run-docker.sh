#!/bin/bash
echo "Running Docker"

docker run -p 8080:8080 -d -e PORT=8080 --name=gorm-benchmark \
 --network=host \
 -e POSTGRES_DSN="postgresql://postgres:senha123@localhost:5432/postgres?sslmode=disable" \
 -e MYSQL_DSN="root:senha123@tcp(localhost:3306)/teste?parseTime=true" \
 -e TARGET=postgres \
 gorm-benchmark
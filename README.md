# Learning go

## Description
Simple app to learn golang. Everything new knowledge can be docummented here. 

## Development Guide

### Prequisite

- [Go 1.24 or later](https://golang.org/doc/install)

### Guide

1. Copy `env.sample` to `.env`, change env as needed.

   ```sh
   cp env.sample .env
   ```

1. Install dependencies.

   ```sh
   go mod tidy
   ```

1. Run docker under `dev` directory
   
   ```sh
   docker-compose up -d 
   ```
   This will start mysql, redis, and kafka.

   If you want to stop docker you can use this command.
   ```
   docker-compose down
   ```

1. Run app.

   ```sh
   go run app/server/main.go
   ```

1. Web service is available on port 7172.
1. Run consumer
   ```sh
   go run app/consumer/userverification/main.go
   ```

### DB Migrate
1. Install DB migrate tools (optional)
   
   ```sh
   brew install golang-migrate
   ```
1. Create migration file (example)

   ```sh
   migrate create -ext sql -dir db/migrations -tz Local create_users_table
   ```
   This will generate migration file in directory db/migrations with local timezone

1. Edit generated migration file for up and down file.

1. Run DB migration

   ```sh
   migrate -path db/migrations -database "mysql://user:password@tcp(localhost:3306)/dbname" up
   ```
   Replace user, password, and dbname following env config for mysql

### Kafka Topics
1. Check container name for kafka
   ```sh
   docker ps
   ```
1. Exec kafka container
   ```sh
   docker exec -it kafka bash
   ```
1. Create topic for user.verify-user-job
   ```sh
   kafka-topics.sh \
   --create \
   --bootstrap-server localhost:9092 \
   --replication-factor 1 \
   --partitions 1 \
   --topic learning.verify-user-job
   ```
1. Check list topics
   ```
   kafka-topics.sh --list --bootstrap-server localhost:9092
   ```
1. Check consumer lag
   ```sh
   kafka-consumer-groups.sh \
   --bootstrap-server localhost:9092 \
   --group users-group \
   --describe
   ```

### Dependencies
- [Kafka](https://kafka.apache.org/)
- [MySQL](https://www.mysql.com)
- [Cassandra](https://cassandra.apache.org)
- [Redis](https://redis.io)

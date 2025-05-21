#!/bin/sh

# Wait for Postgres to be ready
until pg_isready -h db -p 5432 -U postgres; do
    echo "Waiting for postgres..."
    sleep 2
done

echo "Listing /app directory:"
ls -al /app

echo "Listing /app/migrations directory:"
ls -al /app/migrations

# Run migration
echo "Running database migrations..."
migrate -path /app/migrations -database "postgres://postgres:postgres1234@db:5432/postgres?sslmode=disable" up

# Run seeder
echo "Running database seeders..."
# psql "postgres://postgres:postgres1234@db:5432/postgres?sslmode=disable" -f /app/migrations/dbseed/000001_periode_table_seed.up.sql

# Run the bot
echo "Starting activity-bot..."
./activity-bot

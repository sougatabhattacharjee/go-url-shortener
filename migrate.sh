#!/bin/sh

# Set the PGPASSWORD environment variable for the psql command
export PGPASSWORD=$POSTGRES_PASSWORD

# Wait until PostgreSQL is ready
until pg_isready -h db -U $POSTGRES_USER; do
  echo "Waiting for postgres..."
  sleep 2
done

# Run the migration script
psql -h db -U $POSTGRES_USER -d $POSTGRES_DB -f /docker-entrypoint-initdb.d/init.sql

# Keep the container running after the migration
tail -f /dev/null

#!/bin/bash
# Apply Up/Down migrations using Docker
# Usage:
# dmgr.sh <sequence_op> <steps>
# e.g.
# dmgr.sh up 1
# dmgh.sh down 2

# Use DB secrets from external file
# should contain export DB_USER=... export DB_PASSWORD=.. and export DB_NAME=...

if [ $# -ne 2 ]; then
  echo "Usage: $0 <sequence_op> <steps>"
  exit 1
fi

# Get script dir
DIR="$( cd "$( dirname "$0" )" && pwd )"
source $DIR/db.secret

# Get migration files dir
MIGRATIONS=""$( cd "$DIR/../migrations" && pwd )""

echo Migrations Path: $MIGRATIONS
echo Database: $DB_NAME

sequence_op="$1"
steps="$2"

docker run -v $MIGRATIONS:/migrations --network $DOCKER_NETWORK migrate/migrate -path=/migrations/ -database postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable $sequence_op $steps


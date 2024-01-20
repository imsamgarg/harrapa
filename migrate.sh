#!/bin/bash

if [ "$1" != "up" ] && [ "$1" != "down" ]; then
    echo "Usage: $(basename $0) <mode>[up,down]"
    exit 1
fi

source .env
cd ./sql/schema
goose postgres "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$DB_HOST:$DB_PORT/$POSTGRES_DB?sslmode=disable" $1
cd -
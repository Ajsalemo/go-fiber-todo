#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "todos" <<-EOSQL
    CREATE TABLE tasks(
        ID SERIAL PRIMARY KEY NOT NULL,
        name VARCHAR(500) NOT NULL,
        completed BOOLEAN NOT NULL DEFAULT false
    );
EOSQL


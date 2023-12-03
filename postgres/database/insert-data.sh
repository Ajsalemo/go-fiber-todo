#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "todos" <<-EOSQL
    INSERT INTO tasks (
        name,
        completed
    ) VALUES
    ('replace door hinges', false),
    ('wash the car', true);
EOSQL


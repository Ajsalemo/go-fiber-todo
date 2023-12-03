#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "postgres" <<-EOSQL
	CREATE DATABASE todos;
	GRANT ALL PRIVILEGES ON DATABASE todos TO postgres;
EOSQL


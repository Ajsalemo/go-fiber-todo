FROM postgres:latest

COPY ./database/create-db.sh /docker-entrypoint-initdb.d
COPY ./database/create-table.sh /docker-entrypoint-initdb.d
COPY ./database/insert-data.sh /docker-entrypoint-initdb.d
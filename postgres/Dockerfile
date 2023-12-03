FROM postgres:latest

ARG POSTGRES_PASSWORD mysecretpassword

COPY ./database/create-db.sh /docker-entrypoint-initdb.d
COPY ./database/create-table.sh /docker-entrypoint-initdb.d
COPY ./database/insert-data.sh /docker-entrypoint-initdb.d
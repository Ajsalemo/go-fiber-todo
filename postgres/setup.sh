#!/bin/bash

docker pull postgres:latest
docker run -d -p 5432:5432 -it --rm --network postgres-network --name postgres -e POSTGRES_PASSWORD=mysecretpassword postgres:latest
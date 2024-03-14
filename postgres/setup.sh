#!/bin/bash

docker pull postgres:latest
docker run -d -p 5432:5432 -it --rm --network postgres-network --name postgres -e POSTGRES_PASSWORD=mysecretpassword postgres:latest
# This runs the 'custom' postgres image used with the Dockerfile under postgres/Dockerfile 
# This image runs the init scripts defined in the Dockerfile
docker run -d -p 5432:5432 -it --rm --network postgres-network --name go-fiber-todo-postgres -e POSTGRES_PASSWORD=mysecretpassword go-fiber-todo-postgres:latest

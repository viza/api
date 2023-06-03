#!/bin/bash
SERVER_ADDRESS="localhost" 
SERVER_PORT="8282"
DB_USER="root" 
DB_PASSWD="!vetalr00t!"
DB_ADDR="localhost"
DB_PORT="3306"
DB_NAME="banking"

export SERVER_ADDRESS=$SERVER_ADDRESS
export SERVER_PORT=$SERVER_PORT 
export DB_USER=$DB_USER
export DB_PASSWD=$DB_PASSWD
export DB_ADDR=$DB_ADDR
export DB_PORT=$DB_PORT
export DB_NAME=$DB_NAME

go run main.go
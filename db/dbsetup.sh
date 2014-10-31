#!/bin/bash

export PGUSER="${PGUSER:-trivia}"
export PGPASSWORD="${PGPASSWORD:-theansweris}"
export PGDB="${PGDB:-trivia}"

gosu postgres postgres --single <<- EOSQL
create database $PGDB;
create user $PGUSER password '$PGPASSWORD';
grant all privileges on database $PGDB to $PGUSER;
EOSQL

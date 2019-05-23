#!/bin/bash
for i in `seq 1 20`;
do
    pg_isready -h localhost -p 5432 && \
    echo 'Success' && exit 0
    echo -n .
    sleep 1
done
echo 'Failed waiting for Postgres' && exit 1

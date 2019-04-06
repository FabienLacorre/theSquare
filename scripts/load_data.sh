#!/bin/sh

# File: load_data.sh
# File Created: 11 Mar 2019 14:32
# By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

docker rm -f `docker ps -aq`
docker-compose up -d neo4j
docker-compose exec neo4j /bin/bash -c 'bin/neo4j-admin load --from=data/preset/2019-03-25.dump --force && exit'
docker rm -f `docker ps -aq`

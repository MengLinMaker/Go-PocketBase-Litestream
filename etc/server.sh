#!/bin/ash
set +e

echo "Restoring database"
litestream restore -if-db-not-exists -if-replica-exists /db/data.db
litestream restore -if-db-not-exists -if-replica-exists /db/logs.db

if [ "$STAGE" == "PROD" ]
then
  echo "Replicate PROD database"
  litestream replicate
else
  echo "Don't replicate DEV database"
  /server.bin serve --http="0.0.0.0:8080"
fi

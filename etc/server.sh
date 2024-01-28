#!/bin/ash
set +e

echo "Restoring database"
rm -r /db
litestream restore -if-replica-exists /db/data.db
litestream restore -if-replica-exists /db/logs.db

if [ "$STAGE" == "prod" ]
then
  echo "Replicate PROD database"
  litestream replicate
else
  echo "Don't replicate DEV database"
  /server.bin serve --http="127.0.0.1:8080"
fi
echo "Starting server"

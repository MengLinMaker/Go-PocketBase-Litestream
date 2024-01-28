#!/bin/ash
set +e

echo "Restoring database"
litestream restore /db/data.db
litestream restore /db/logs.db

if [ "$STAGE" == "prod" ]
then
  echo "Replicate PROD database"
  litestream replicate
else
  echo "Don't replicate DEV database"
  /server.bin serve --http="127.0.0.1:8080"
fi
echo "Starting server"

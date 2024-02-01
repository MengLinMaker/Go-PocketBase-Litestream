#!/bin/ash
set +e

echo "Restoring database"
litestream restore -if-db-not-exists -if-replica-exists ./pb_data/data.db
litestream restore -if-db-not-exists -if-replica-exists ./pb_data/logs.db

if [ "$STAGE" == "PROD" ]
then
  echo "Replicate PROD database"
  litestream replicate &
else
  echo "Don't replicate DEV database"
fi

/server.bin serve --http="0.0.0.0:8080" --dev

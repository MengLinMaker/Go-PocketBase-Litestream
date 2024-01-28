#!/bin/ash
set +e

echo "Restoring database"
litestream restore /db/data.db
litestream restore /db/logs.db

if [ $STAGE==prod ]; then
  echo "Replicating production database"
  litestream replicate
fi
echo "Starting server"

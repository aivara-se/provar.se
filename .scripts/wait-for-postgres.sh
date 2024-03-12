#!/bin/sh

failcounter=0

until pg_isready --username=postgres --host="127.0.0.1" --quiet &> /dev/null ; do
    let "failcounter += 1"
    echo "Waiting for database connection..."
     if [[ $failcounter -gt 60 ]]; then
      echo "Failed to connect to database"
      exit 1
    fi
    sleep 2
done

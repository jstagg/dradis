#!/bin/bash

# Start process
redis-server --bind 0.0.0.0 &

# Wait for pid 
until pids=$(pidof redis-server)
do   
    sleep 1
done

# Import the data
sleep 5
cat /data/data.txt | redis-cli --pipe 
sleep 5

# Let the container run indefinitely
mkfifo /tmp/mypipe

while read SIGNAL; do
    case "$SIGNAL" in
        *EXIT*)break;;
        *)echo "signal  $SIGNAL  is unsupported" >/dev/stderr;;
    esac
done < /tmp/mypipe

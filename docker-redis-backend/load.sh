#!/bin/bash

# Start process
redis-server &

# Wait for pid 
until pids=$(pidof redis-server)
do   
    sleep 1
done

# Import the data
sleep 5
cat data.txt | redis-cli --pipe 

# Let the container run indefinitely
mkfifo /tmp/mypipe

while read SIGNAL; do
    case "$SIGNAL" in
        *EXIT*)break;;
        *)echo "signal  $SIGNAL  is unsupported" >/dev/stderr;;
    esac
done < /tmp/mypipe

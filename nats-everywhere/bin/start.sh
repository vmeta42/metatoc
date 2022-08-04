#!/bin/bash

if [ $1 ]; then
    nats_host=$1
else
    nats_host="172.22.50.211"
fi

nats_port="4222"

#echo "nats_address: $nats_host:$nats_port"
#exit 0
 
echo "Check docker..."
docker -v
if [ $? -eq 0 ]; then
    echo "Docker installed!"
else
    echo "Docker uninstalled! Pleaset install it first. see https://www.docker.com/"
    exit 0
fi
#exit 0

echo "Check nats exists..."
nats_exists=$(docker image ls "nats:latest" | wc -l)
if [ $nats_exists -eq 2 ]; then
    echo "Nats exists!"
else
    echo "Pull nats..."
    docker pull nats:latest
fi
#exit 0

echo "Check nats running..."
nats_running=$(docker ps --filter=name=metatoc-nats | wc -l)
if [ $nats_running -eq 2 ]; then
    echo "Nats is running!"
else
    echo "Run nats..."
    docker run --name metatoc-nats --rm --network host --rm -d -p 4222:4222 nats:latest -js
fi
#exit 0

echo "Check nats-box exists..."
nats0box_exists=$(docker image ls "natsio/nats-box:latest" | wc -l)
if [ $nats0box_exists -eq 2 ]; then
    echo "Nats-box exists!"
else
    echo "Pull nats-box..."
    docker pull natsio/nats-box:latest
fi
#exit 0

echo "Add nats stream..."
docker run --rm -ti natsio/nats-box:latest -c "nats stream add STREAM_METATOC --subjects "CONSUMER_METATOC.*" --storage file --retention limits --discard=old --max-msgs=-1 --max-msgs-per-subject=-1 --max-bytes=-1 --max-age=-1 --max-msg-size=-1 --dupe-window=2m --replicas=1 -s $nats_host:$nats_port"

echo "Add nats consumer..."
docker run --rm -ti natsio/nats-box:latest -c "nats consumer add STREAM_METATOC CONSUMER_METATOC --pull --deliver=all --replay=instant --filter="" --max-deliver=-1 --max-pending=0 -s $nats_host:$nats_port"
#exit 0

echo "Check nats-everywhere exists..."
nats0everywhere_exists=$(docker image ls "nats-everywhere:latest" | wc -l)
if [ $nats0everywhere_exists -eq 2 ]; then
    echo "Nats-everywhere exists!"
else
    echo "Make nats-everywhere..."
    docker build --tag nats-everywhere:latest --build-arg NATS_URL=nats://$nats_host:$nats_posrt --build-arg NATS_SUBJECT=CONSUMER_METATOC.* --build-arg NATS_DURABLE=CONSUMER_METATOC ../.
fi
#exit 0

echo "Check nats-everywhere running..."
nats0everywhere_running=$(docker ps --filter=name=nats-everywhere | wc -l)
if [ $nats0everywhere_running -eq 2 ]; then
    echo "Nats-everywhere running!"
else
    echo "Run nats-everywhere..."
    docker run --name nats-everywhere --rm -ti -d -v $PWD/../logs:/opt/logs nats-everywhere:latest
fi
#exit 0

#!/bin/bash

./diem-swarm -c ./etc --diem-node ./node -n 4
echo "Copy config"

for i in 0 1 2 3;
do
# sed -i "74s/.*/    listen_address: \/ip4\/0.0.0.0\/tcp\/8001/g" /opt/diem/etc/$i/node.yaml;
sed -i "109s/.*/  address: \"0.0.0.0:50001\"/g" /opt/diem/etc/$i/node.yaml;
# sed -i "158s/.*/  listen_address: \/ip4\/0.0.0.0\/tcp\/8000/g" /opt/diem/etc/$i/node.yaml;
done

cp -r ./etc/* /opt/diem/out_etc/

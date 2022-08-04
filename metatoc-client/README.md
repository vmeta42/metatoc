# metatoc-client

## Getting started

1. Run MetaTOC client.

```
cd ./bin/
chmod 0777 ./start.sh
./start.sh
```

2. Create a new wallet, address and private key will be returned after successful execution.

```
MetaTOC >> signup
SUCCESSFUL
address is [c90822febb489091c62f77dcca68xxxx]
private key is [44a336f7cffc0f38071c41d9fa487489e551ca0f7d4a2a7618aa80250fecxxxx]
```

3. Create a new block on chain with new HDFS path.

```
MetaTOC >> create -a c90822febb489091c62f77dcca68xxxx -k 44a336f7cffc0f38071c41d9fa487489e551ca0f7d4a2a7618aa80250fecxxxx -p "/meta42/metatoc/brkjC2F6" -c "Welcome to MetaTOC meetup!"
SUCCESSFUL
```

4. Return the detail of data related to HDFS path.

```
MetaTOC >> detail -a c90822febb489091c62f77dcca68xxxx -k 44a336f7cffc0f38071c41d9fa487489e551ca0f7d4a2a7618aa80250fecxxxx -p "/meta42/metatoc/brkjC2F6"
SUCCESSFUL
data is [Welcome to MetaTOC meetup!]
```

5. Create another wallet.

```
MetaTOC » signup
SUCCESSFUL
address is [16442a1791ac9a2ced99d719a197xxxx]
private key is [890ffe5e387198f8adaf5d9617b7fffa3ed55bac698e4828b030653e68b9xxxx]
```

6. List HDFS resource paths.

```
MetaTOC » list -a c90822febb489091c62f77dcca68xxxx
SUCCESSFUL
There are 1 pieces of data
The 1st data is [/meta42/metatoc/brkjC2F6]
MetaTOC » list -a 16442a1791ac9a2ced99d719a197xxxx
SUCCESSFUL
no data
```

7. Share block with other wallet.

```
MetaTOC » share -f c90822febb489091c62f77dcca68xxxx -o 16442a1791ac9a2ced99d719a197xxxx -k 44a336f7cffc0f38071c41d9fa487489e551ca0f7d4a2a7618aa80250fecxxxx -t "/meta42/metatoc/brkjC2F6"
SUCCESSFUL
```

8. List HDFS resource paths again.

```
MetaTOC » list -a c90822febb489091c62f77dcca68xxxx
SUCCESSFUL
There are 1 pieces of data
The 1st data is [/meta42/metatoc/brkjC2F6]
MetaTOC » list -a 16442a1791ac9a2ced99d719a197xxxx
SUCCESSFUL
There are 1 pieces of data
The 1st data is [/meta42/metatoc/brkjC2F6]
```
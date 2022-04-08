# CEFCO

Cloud-Edge File Coordination Operator, can transfer file from edge node to cloud by NATS.

## Prepare

```bash
# clone repo
$ git clone https://github.com/inspursoft/cefco
```

if not deploy nats2.2, deploy nats-jetstream for cloud node(refer https://docs.nats.io/nats-server/installation and https://docs.nats.io/jetstream/getting_started).

Build reader image for edge node and writer image for cloud, please refer the **[readme.md](filesync/readme.md)**



## Build run

```bash
# create crd
$ kubectl apply -f artifacts/crd.yaml

# run for test
$ make filesync-controller VERSION=0.1.0 noRace=T
$ ./filesync-controller --kubeconfig /root/.kube/config

# run for docker image
$ make filesync-controller-docker filesyncVersion=0.1.0
# push it to your registry
# run
$ kubectl apply -f artifacts/rbac.yaml
# deploy the operator(assign the hostname and image version)
$ kubectl apply -f artifacts/deployment.yaml

# create an example
# assign hostname and image for reader and writer, include nats ip and port
$ kubectl apply -f artifacts/examples/example-mini.yaml
# delete the example if undeploy
$ kubectl delete -f artifacts/examples/example-mini.yaml
```

## tip for nats-jetstream

```bash
# Create a stream
$ nats --user s1 --password s1 str add example --subjects="example.>" --storage=memory --retention=workq --discard=new --max-msgs=-1 --max-bytes=-1 --max-age=-1 --max-msg-size=-1 --dupe-window=2s --replicas=1
# Create a consumer
nats --user s1 --password s1 con add example mini --deliver=all --replay=instant --filter="example.mini" --max-deliver=-1 --max-pending=1 --pull
```
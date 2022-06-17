# Deploy the provider

## build images & push

```shell
cd crossplane-provider
make

docker tag build-f77e8e61/provider-poc-controller-amd64 localhost:5001/provider-poc-controller:0.1
docker push localhost:5001/provider-poc-controller:0.1


## Install controller & CRD
```shell
kubectl apply -k https://github.com/crossplane/crossplane/cluster\?ref\=master
kubectl apply -f package/crossplane_local.yaml
kubectl apply -f package/crds/*.yaml
```

## Deploy example
```shell
kubectl apply -f examples/provider/*
kubectl apply -f examples/sample/*
```
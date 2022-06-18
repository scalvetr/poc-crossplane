# Deploy the provider

## Build the provider & push

```shell
git clone git@github.com:scalvetr/poc-crossplane-provider.git
cd poc-crossplane-provider

docker login docker.io --username scalvetr

export PLATFORMS="linux_amd64";
export DOCKER_REGISTRY="scalvetr";
make build publish
```

## Install the provider to the K8s cluster

```shell
kubectl crossplane install provider scalvetr/poc-crossplane-provider:latest
kubectl get crossplane
```

## Deploy example
```shell
kubectl apply -f examples/provider/*
kubectl apply -f examples/sample/*
```

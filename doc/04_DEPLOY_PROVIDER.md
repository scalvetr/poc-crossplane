# Deploy the provider

## build images & push

```shell

git clone git@github.com:scalvetr/poc-crossplane-provider.git
cd poc-crossplane-provider
make

docker tag build-f77e8e61/poc-crossplane-provider-controller-amd64 localhost:5001/poc-crossplane-provider-controller:0.1
docker push localhost:5001/poc-crossplane-provider-controller:0.1


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


Useful commands
```shell
# generate crds
make generate

# test
make test

# build the image
make build
make build.all

docker tag build-f77e8e61/poc-crossplane-provider-controller-amd64 localhost:5001/poc-crossplane-provider-controller:0.1

cd package
# build the xpkg file
kubectl crossplane build provider

# build the xpkg file
kubectl crossplane push provider localhost:5001/poc-crossplane-provider-controller:0.1

cd ..
# deploy examples
```
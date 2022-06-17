# Deploy the provider

## build images & push

```shell
git clone git@github.com:scalvetr/poc-crossplane-provider.git
cd poc-crossplane-provider
make

docker tag build-e2c376fb/poc-crossplane-provider-amd64 localhost:5001/poc-crossplane-provider:0.1
docker tag build-e2c376fb/poc-crossplane-provider-controller-amd64 localhost:5001/poc-crossplane-provider-controller:0.1

docker push localhost:5001/poc-crossplane-provider:0.1
docker push localhost:5001/poc-crossplane-provider-controller:0.1

```

##Build the package

```shell
cd package
mv crossplane.yaml ../crossplane.yaml.bkp

cat > crossplane.yaml << EOF
apiVersion: meta.pkg.crossplane.io/v1alpha1
kind: Provider
metadata:
  name: poc-crossplane-provider
  annotations:
    meta.crossplane.io/maintainer: Crossplane Maintainers <info@crossplane.io>
    meta.crossplane.io/source: github.com/scalvetr/poc-crossplane-provider-provider
    meta.crossplane.io/license: Apache-2.0
    meta.crossplane.io/description: |
      A template that can be used to create Crossplane providers.

spec:
  controller:
    image: localhost:5001/localhost:5001/poc-crossplane-provider-controller:0.1

EOF

kubectl crossplane build provider

kubectl crossplane push provider localhost:5001/poc-crossplane-provider-controller:0.1
kubectl crossplane install provider localhost:5001/poc-crossplane-provider-controller:0.1
mv ../crossplane.yaml.bkp crossplane.yaml
cd ..

#Deploy the CRDs 
kubectl apply -f package/crds/objects.poc.crossplane.io_topics.yaml
kubectl apply -f package/crds/poc.crossplane.io_providerconfigs.yaml
kubectl apply -f package/crds/poc.crossplane.io_providerconfigusages.yaml
kubectl apply -f package/crds/poc.crossplane.io_storeconfigs.yaml
```

## Deploy example
```shell
kubectl apply -f examples/provider/*
kubectl apply -f examples/sample/*
```

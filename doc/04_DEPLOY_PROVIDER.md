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
# kubectl crossplane install provider scalvetr/poc-crossplane-provider:v0.0.0-5.g8112ee1

cat <<EOF | kubectl create -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: poc-crossplane-provider
spec:
  package: "scalvetr/poc-crossplane-provider:v0.0.0-12.g6452e7e"
EOF

kubectl get crossplane
```

## Undeploy the provider
```shell
kubectl delete providers.pkg.crossplane.io poc-crossplane-provider
```
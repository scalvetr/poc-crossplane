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

## Deploy example
```shell
kubectl apply -f examples/provider/*
cat <<EOF | kubectl create -f -
apiVersion: objects.poc.crossplane.io/v1alpha1
kind: Topic
metadata:
  name: topic1
spec:
  forProvider:
    name: topic1
    partitions: 4
  providerConfigRef:
    name: tenant1
EOF

curl -k --header "x-tenant: tenant1" https://service-api.localtest.me/v1/topics
# [{"name":"topic1","partitions":5,"status":null}]

cat <<EOF | kubectl apply -f -
apiVersion: objects.poc.crossplane.io/v1alpha1
kind: Topic
metadata:
  name: topic1
spec:
  forProvider:
    name: topic1
    partitions: 4
  providerConfigRef:
    name: tenant1
EOF
# topic.objects.poc.crossplane.io/topic1 configured

curl -k --header "x-tenant: tenant1" https://service-api.localtest.me/v1/topics/topic1
# {"name":"topic1","partitions":4,"status":null}
```

Useful commands
```shell
# get the provider config
kubectl get providerconfig tenant1 -o yaml

# get desired status for topic1
kubectl get topic topic1 -o yaml

# get current status of topic1
curl -k --header "x-tenant: tenant1" https://service-api.localtest.me/v1/topics/topic1

# check service logs
kubectl -n service-api logs -l app=service-api 
```

## Undeploy everything
```shell
kubectl delete -f examples/sample/* --grace-period=0 --force
kubectl delete -f examples/provider/* --grace-period=0 --force
kubectl delete providers.pkg.crossplane.io poc-crossplane-provider
```
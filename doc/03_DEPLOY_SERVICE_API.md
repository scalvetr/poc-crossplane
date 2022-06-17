# Deploy the Service API in ArgoCD


## Deploy service to local registry
```shell
cd service-api
./gradlew bootBuildImage --imageName=localhost:5001/service-api:0.1

# push image to registry
docker push localhost:5001/service-api:0.1

# test
docker run -p 8080:8080 localhost:5001/service-api:0.1
```

## Install service to the Kubernetes cluster

Init variables
```shell
export TARGET_NAMESPACE="service-api"
```

```shell
helm install servie-api servie-api-helm --values servie-api-helm/values.yaml
```
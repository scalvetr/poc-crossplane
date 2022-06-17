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

Init the namespace
```shell
export TARGET_NAMESPACE="service-api"

cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Namespace
metadata:
  name: ${TARGET_NAMESPACE}
EOF
```

Install the service
```shell
helm install service-api service-api-helm --values service-api-helm/values.yaml
```
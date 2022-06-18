# Deploy the Service API in ArgoCD


## Deploy service to local registry
```shell
cd service-api
./gradlew bootBuildImage --imageName=localhost:5001/service-api:0.1

# push image to registry
docker push localhost:5001/service-api:0.1

# test
docker run -p 8080:8080 localhost:5001/service-api:0.1

# Ctrl+C to stop
cd ..
```

## Install service to the Kubernetes cluster


Install the service
```shell
export TARGET_NAMESPACE="service-api"
helm install service-api service-api-helm -n ${TARGET_NAMESPACE} --create-namespace --values service-api-helm/values.yaml

kubectl -n ${TARGET_NAMESPACE} get pods
#NAME                           READY   STATUS    RESTARTS   AGE
#service-api-6ddcc85df6-m6n5h   1/1     Running   0          25s

# test
curl -k https://service-api.localtest.me/actuator/health

echo "open https://service-api.localtest.me/actuator/swagger-ui.html"
```

Uninstall
```shell
helm uninstall service-api -n ${TARGET_NAMESPACE}

```
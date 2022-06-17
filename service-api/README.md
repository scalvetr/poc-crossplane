# Sample Application

## Test 

```shell
# Start locally
./gradlew bootRun

curl http://localhost:8080/actuator/health


http://localhost:8080/actuator/swagger-ui.html
```

## Deploy lo local registry
```shell
./gradlew bootBuildImage --imageName=localhost:5001/service-api:0.1
# test
docker run -p 8080:8080 localhost:5001/service-api:0.1
docker push localhost:5001/service-api:0.1
```


## Test Helm Chart
```shell
helm uninstall service-api

helm install service-api service-api-helm \
--set image.pullPolicy=Always \
--values service-api-helm/values.yaml
```
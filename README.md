# POC CROSSPLANE

## INSTALLATION

* [Install K8S local cluster with KinD](doc/01_K8S_LOCAL_CLUSTER_KIND.md)
* [Install ArgoCD and Crossplane](doc/02_INSTALL_CROSSPLANE.md)
* [Deploy Sample Application to Local Registry](doc/03_SAMPLE_APPLICATION_BUILD_LOCAL_REGISTRY.md)
* [Configure Sample Application in ArgoCD](doc/04_SAMPLE_APPLICATION_CONFIGURE_ARGOCD.md)

## Argo CD: Access UI dashboard
```shell
export admin_password=$(kubectl -n argocd get secret argocd-initial-admin-secret --template={{.data.password}} | base64 -d; echo)
export admin_username="admin"

echo """
ArgoCD Installed
URL: https://argocd.localtest.me/
USERNAME: ${admin_username}
PASSWORD: ${admin_password}
"""
```

## Deploy the service API

Follow the instructions [here](doc/03_DEPLOY_SERVICE_API.md)

See the API here

https://service-api.localtest.me/actuator/swagger-ui.html

# Install Crossplane

Follow instructions [here](https://crossplane.io/docs/v1.6/getting-started/install-configure.html)

```shell
kubectl create namespace crossplane-system

helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
```

Check Crossplane Status
```shell
helm list -n crossplane-system

kubectl get all -n crossplane-system
```

Insatall the crossplane CLI

```shell
curl -sL https://raw.githubusercontent.com/crossplane/crossplane/master/install.sh | sh
```
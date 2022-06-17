# provider-poc

`provider-poc` is a minimal [Crossplane](https://crossplane.io/) Provider
that is meant to be used as a template for implementing new Providers. It comes
with the following features that are meant to be refactored:

- A `ProviderConfig` type that only points to a credentials `Secret`.
- A `MyType` resource type that serves as an example managed resource.
- A managed resource controller that reconciles `MyType` objects and simply
  prints their configuration in its `Observe` method.

## Developing

1. Use this repository as a template to create a new one.
1. Find-and-replace `provider-poc` with your provider's name.
1. Run `make` to initialize the "build" Make submodule we use for CI/CD.
1. Run `make reviewable` to run code generation, linters, and tests.
1. Replace `MyType` with your own managed resource implementation(s).

Refer to Crossplane's [CONTRIBUTING.md] file for more information on how the
Crossplane community prefers to work. The [Provider Development][provider-dev]
guide may also be of use.

[CONTRIBUTING.md]: https://github.com/crossplane/crossplane/blob/master/CONTRIBUTING.md
[provider-dev]: https://github.com/crossplane/crossplane/blob/master/docs/contributing/provider_development_guide.md


Useful commands
```shell
# generate crds
make generate

# test
make test

# build the image
make build
make build.all

docker tag build-f77e8e61/provider-poc-controller-amd64 localhost:5001/provider-poc-controller:0.1

cd package
# build the xpkg file
kubectl crossplane build provider

# build the xpkg file
kubectl crossplane push provider localhost:5001/provider-poc-controller:0.1

cd ..
# deploy examples
```
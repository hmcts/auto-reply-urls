# k8s-operator-auto-reply-urls

This is an operator for automatically adding and removing reply urls to an application in azure ad based on ingress events

It was built on [operator-sdk](https://github.com/operator-framework/operator-sdk)

## Building this project

* Install the operator-sdk [CLI](https://github.com/operator-framework/operator-sdk/blob/master/doc/user/install-operator-sdk.md)

Export this value to your shell:
```bash
$ export GO111MODULE=on
```

Run build:
```bash
$ operator-sdk build hmctspublic.azurecr.io/rpe/auto-reply-urls:version
```

Deploying it:
```bash
$ az ad sp create-for-rbac --skip-assignment
$ kubectl create secret generic --from-literal azure_client_id= --from-literal azure_client_secret= --from-literal azure_tenant_id= --from-literal object_id= 
$ az acr login --name hmctspublic --subscription DCD-DNP-DEV
$ docker push hmctspublic.azurecr.io/rpe/auto-reply-urls:version
```

Note the SP needs permissions as detailed here:
https://stackoverflow.com/a/53014616/4951015

Update the version in deploy/operator.yaml and then:
```bash
$ kubectl apply -R -f deploy/
```

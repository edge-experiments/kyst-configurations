## Argo CD Scalability Experiment

### Create an Argo CD application from command line
For example:
```shell
argocd app create argocd-scalability-00001 \
--config-management-plugin wrap4kyst-flotta \
--repo https://github.com/edge-experiments/kyst-configurations.git \
--path examples/kubernetes/nginx/deploy-flotta \
--dest-server https://$sharedkystmaster:6443  \
--dest-namespace argocd-scalability
```

### Create Argo CD applications using ApplicationSet
Argo CD's [ApplicationSet](https://argo-cd.readthedocs.io/en/stable/user-guide/application-set/) controller can automate the creation of multiple Argo CD applications, using 'generators' and 'templates'.

There are [various generators](https://argo-cd.readthedocs.io/en/stable/operator-manual/applicationset/Generators/) provided by Argo.
Looks like the [Git Generator](https://argo-cd.readthedocs.io/en/stable/operator-manual/applicationset/Generators-Git/) fits into our use case.
Here is an example:
```shell
kubectl -n argocd apply -f argocd-scalability/applicationset.yaml
```

### Create Argo CD applications using Cluster Loader
[Cluster Loader](https://github.com/kubernetes/perf-tests/tree/master/clusterloader2) can automate the creation of multiple Argo CD applications, because an Argo CD application is a Kubernetes Custom Resource.

However, a little hack is necessary to make it work. The intial trial failed, because Cluster Loader today only allows objects to be created in its managed (meaning created at runtime with randomized names) namespaces. But Argo CD only recogonizes Applications in `argocd` namespace. We can hack Cluster Loader's [code](https://github.com/kubernetes/perf-tests/blob/8a0c339a42a6f0419a10fd3a701a8284c37511f3/clusterloader2/pkg/test/simple_test_executor.go#L181) as a work around. More discussions available [here](https://kubernetes.slack.com/archives/C09QZTRH7/p1657089711061019).

We also need the `unique-configspec-name` option for the `wrap4kyst` plugin. Using this option, even multiple Argo CD applications are created with one single set of files in git, the output ConfigSpecs are still unique. Therefore, the multiple Argo CD applications won't conflict with each other. With this option, we eliminate the need of creating a separate set of files for each Argo CD application.

### Sync an Argo CD application using local directory
```shell
argocd app sync argocd-scalability-00001 --local ./examples/kubernetes/nginx/deploy-flotta/
```

#### Prerequisites
- `kustomize` installed if using `--local` option for `argocd app sync`.
- `wrap4kyst` binary in `PATH` if using the `wrap4kyst` plugin and using the `--local` option for `argocd app sync`. The binary can be made available by compliling the `wrap4kyst` plugin locally, or by copying a complied one from a container:
```shell
docker cp 23058b2b81fb:/wrap4kyst .
```

### Questions
- Telemetry?
- Apps v.s. clusters?

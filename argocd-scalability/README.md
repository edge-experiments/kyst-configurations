### Argo CD Scalability Experiment

#### Manually create an Argo CD application
For example:
```shell
argocd app create argocd-scalability-00001 \
--config-management-plugin wrap4kyst-flotta \
--repo https://github.com/edge-experiments/kyst-configurations.git \
--path examples/kubernetes/nginx/deploy-flotta \
--dest-server https://$sharedkystmaster:6443  \
--dest-namespace argocd-scalability
```

#### Automation using Argo CD ApplicationSet
Argo CD's [ApplicationSet](https://argo-cd.readthedocs.io/en/stable/user-guide/application-set/) controller can automate the creation of multiple Argo CD applications, using 'generators' and 'templates'.

There are [various generators](https://argo-cd.readthedocs.io/en/stable/operator-manual/applicationset/Generators/) provided by Argo.
Looks like the [Git Generator](https://argo-cd.readthedocs.io/en/stable/operator-manual/applicationset/Generators-Git/) fits into our use case.

#### Sync an Argo CD application using local directory
```shell
argocd app sync argocd-scalability-00001 --local ./examples/kubernetes/nginx/deploy-flotta/
```

##### Prerequisites
- `kustomize` installed if using `--local` option for `argocd app sync`.
- `wrap4kyst` binary in `PATH` if using the `wrap4kyst` plugin and using the `--local` option for `argocd app sync`. The binary can be made available by compliling the `wrap4kyst` plugin locally, or by copying a complied one from a container:
```shell
docker cp 23058b2b81fb:/wrap4kyst .
```

#### Questions
- Telemetry?
- (Pure) configurations v.s. workloads?
- Apps v.s. clusters?

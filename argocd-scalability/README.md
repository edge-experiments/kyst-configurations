### ArgoCD Scalability Experiment

#### Create an Argo CD application
For example:
```shell
argocd app create argocd-scalability-0001 \
--config-management-plugin wrap4kyst-flotta \
--repo https://github.com/edge-experiments/kyst-configurations.git \
--path examples/kubernetes/nginx/deploy-flotta \
--dest-server https://$sharedkystmaster:6443  \
--dest-namespace argocd-scalability
```

#### Sync an Argo CD application using local directory
```shell
argocd app sync argocd-scalability-0001 --local ./examples/kubernetes/nginx/deploy-flotta/
```

#### Prerequisites
- `kustomize` installed if using `--local` option for `argocd app sync`.
- `wrap4kyst` binary in `PATH` if using the `wrap4kyst` plugin and using the `--local` option for `argocd app sync`. The binary be made available by compliling the `wrap4kyst` plugin locally, or by copying a complied one from a container:
```shell
docker cp 23058b2b81fb:/wrap4kyst .
```

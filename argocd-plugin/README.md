Register the `wrap4kyst` Argo CD plugin:
```shell
kubectl patch -n argocd configmap argocd-cm --patch-file argocd-plugin/argocd-cm-patch.yaml
```

An example to create an Argo CD application using the plugin:
```shell
argocd app create kyst-configuration-demo \
    --config-management-plugin wrap4kyst \
    --repo https://github.com/edge-experiments/kyst-configurations.git \
    --path simulated/kubernetes/guestbook/deploy \
    --dest-server https://12.34.56.78:6443 \
    --dest-namespace default
```

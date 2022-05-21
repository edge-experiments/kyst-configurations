Make the custom tool available in argocd-repo-server pod.
Replace the dummy `image` of `argocd-plugin/argocd-repo-server-patch.yaml` with real image, then run:
```shell
kubectl patch -n argocd deployment argocd-repo-server --patch-file=argocd-plugin/argocd-repo-server-patch.yaml
```

Register the `wrap4kyst` Argo CD plugin:
```shell
kubectl patch -n argocd configmap argocd-cm --patch-file argocd-plugin/argocd-cm-patch.yaml
```

An example to create an Argo CD application using the plugin:
```shell
argocd app create kyst-configuration-demo \
    --config-management-plugin wrap4kyst \
    --repo https://github.com/edge-experiments/kyst-configurations.git \
    --path examples/kubernetes/guestbook/deploy \
    --dest-server https://12.34.56.78:6443 \
    --dest-namespace default
```

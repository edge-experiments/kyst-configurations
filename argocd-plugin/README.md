#### How to use the `wrap4kyst` Argo CD plugin

`wrap4kyst` is an Argo CD [Config Management Plugin](https://argo-cd.readthedocs.io/en/stable/user-guide/config-management-plugins/) that allows kyst users to translate their manifests into kyst Custom Resources.

Build the `wrap4kyst` Argo CD plugin.
```shell
docker build -t <my-image> argocd-plugin/wrap4kyst/
```

Make the `wrap4kyst` binary available in argocd-repo-server pod.
Replace the dummy `image` of `argocd-plugin/argocd-repo-server-patch.yaml` with the built image, then run the following command against the cluster hosting Argo CD.
```shell
kubectl patch -n argocd deployment argocd-repo-server --patch-file=argocd-plugin/argocd-repo-server-patch.yaml
```

Register the `wrap4kyst` Argo CD plugin.
```shell
kubectl patch -n argocd configmap argocd-cm --patch-file argocd-plugin/argocd-cm-patch.yaml
```

Create an Argo CD application using the plugin. For example:
```shell
argocd app create kyst-configuration-demo \
    --config-management-plugin wrap4kyst \
    --repo https://github.com/edge-experiments/kyst-configurations.git \
    --path examples/kubernetes/guestbook/deploy \
    --dest-server https://12.34.56.78:6443 \
    --dest-namespace default
```

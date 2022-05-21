This directory serves as an example to manually access the kyst device server.

1. In the kyst backend cluster's `default` namespace, create a deviceGroup, a configSpec, and an inventoryDevice, all named `bar`.
```shell
kubectl apply -f examples/bar/
```

2. Generate client key and client certificate for inventoryDevice `bar`, as described in the second half of this [section](https://github.com/edge-experiments/kyst#demo-with-one-shot-agent) (regarding PKI identity) in kyst documentation.

3. Send a request to the kyst device server via curl. For example:
```shell
curl https://<device-server-ip>:<device-server-port>/apis/edge.limani.kube/v1alpha1/namespaces/default/devicespecs/bar \
    --cacert /root/limani/hack/pki/ca/ca.crt \
    --cert /root/limani/hack/pki/ca/issued/bar.crt \
    --key /root/limani/hack/pki/clients/bar/private/bar.key
```

The response should be something like this:
```json
{
  "kind": "DeviceSpec",
  "apiVersion": "edge.limani.kube/v1alpha1",
  "metadata": {
    "name": "bar",
    "namespace": "default",
    "creationTimestamp": null
  },
  "content": [
    "apiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: guestbook-ui\nspec:\n  replicas: 1\n  revisionHistoryLimit: 3\n  selector:\n    matchLabels:\n      app: guestbook-ui\n  template:\n    metadata:\n      labels:\n        app: guestbook-ui\n    spec:\n      containers:\n      - image: gcr.io/heptio-images/ks-guestbook-demo:0.2\n        name: guestbook-ui\n        ports:\n        - containerPort: 80\n        resources:\n          limits:\n            cpu: 100m\n            memory: 64Mi\n",
    "apiVersion: v1\nkind: Service\nmetadata:\n  name: guestbook-ui\nspec:\n  ports:\n  - port: 80\n    targetPort: 80\n  selector:\n    app: guestbook-ui\n"
  ],
  "specName": "default/bar",
  "specVersion": "2222552",
  "lastInputTime": "2022-05-17T20:29:59Z"
}
```

This directory serves as an example to manually access the kyst device server.

1. In the kyst backend cluster's `default` namespace, create a deviceGroup, a configSpec, and an inventoryDevice, all named `bar`.
```shell
kubectl apply -f hack/bar/
```

2. Send a request to the kyst device server via curl. For example:
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
    "YXBpVmVyc2lvbjogYXBwcy92MQpraW5kOiBEZXBsb3ltZW50Cm1ldGFkYXRhOgogIG5hbWU6IGd1ZXN0Ym9vay11aQpzcGVjOgogIHJlcGxpY2FzOiAxCiAgcmV2aXNpb25IaXN0b3J5TGltaXQ6IDMKICBzZWxlY3RvcjoKICAgIG1hdGNoTGFiZWxzOgogICAgICBhcHA6IGd1ZXN0Ym9vay11aQogIHRlbXBsYXRlOgogICAgbWV0YWRhdGE6CiAgICAgIGxhYmVsczoKICAgICAgICBhcHA6IGd1ZXN0Ym9vay11aQogICAgc3BlYzoKICAgICAgY29udGFpbmVyczoKICAgICAgLSBpbWFnZTogZ2NyLmlvL2hlcHRpby1pbWFnZXMva3MtZ3Vlc3Rib29rLWRlbW86MC4yCiAgICAgICAgbmFtZTogZ3Vlc3Rib29rLXVpCiAgICAgICAgcG9ydHM6CiAgICAgICAgLSBjb250YWluZXJQb3J0OiA4MAogICAgICAgIHJlc291cmNlczoKICAgICAgICAgIGxpbWl0czoKICAgICAgICAgICAgY3B1OiAxMDBtCiAgICAgICAgICAgIG1lbW9yeTogNjRNaQo=",
    "YXBpVmVyc2lvbjogdjEKa2luZDogU2VydmljZQptZXRhZGF0YToKICBuYW1lOiBndWVzdGJvb2stdWkKc3BlYzoKICBwb3J0czoKICAtIHBvcnQ6IDgwCiAgICB0YXJnZXRQb3J0OiA4MAogIHNlbGVjdG9yOgogICAgYXBwOiBndWVzdGJvb2stdWkK"
  ],
  "specName": "default/bar",
  "specVersion": "179810",
  "lastInputTime": "2022-05-11T17:06:03Z"
}
```

3. The `content` of the returned deviceSpec is base64 encoded:
```shell
ubuntu@ip-192-168-1-38:~$ echo YXBpVmVyc2lvbjogYXBwcy92MQpraW5kOiBEZXBsb3ltZW50Cm1ldGFkYXRhOgogIG5hbWU6IGd1ZXN0Ym9vay11aQpzcGVjOgogIHJlcGxpY2FzOiAxCiAgcmV2aXNpb25IaXN0b3J5TGltaXQ6IDMKICBzZWxlY3RvcjoKICAgIG1hdGNoTGFiZWxzOgogICAgICBhcHA6IGd1ZXN0Ym9vay11aQogIHRlbXBsYXRlOgogICAgbWV0YWRhdGE6CiAgICAgIGxhYmVsczoKICAgICAgICBhcHA6IGd1ZXN0Ym9vay11aQogICAgc3BlYzoKICAgICAgY29udGFpbmVyczoKICAgICAgLSBpbWFnZTogZ2NyLmlvL2hlcHRpby1pbWFnZXMva3MtZ3Vlc3Rib29rLWRlbW86MC4yCiAgICAgICAgbmFtZTogZ3Vlc3Rib29rLXVpCiAgICAgICAgcG9ydHM6CiAgICAgICAgLSBjb250YWluZXJQb3J0OiA4MAogICAgICAgIHJlc291cmNlczoKICAgICAgICAgIGxpbWl0czoKICAgICAgICAgICAgY3B1OiAxMDBtCiAgICAgICAgICAgIG1lbW9yeTogNjRNaQo= | base64 -d
apiVersion: apps/v1
kind: Deployment
metadata:
  name: guestbook-ui
spec:
  replicas: 1
  revisionHistoryLimit: 3
  selector:
    matchLabels:
      app: guestbook-ui
  template:
    metadata:
      labels:
        app: guestbook-ui
    spec:
      containers:
      - image: gcr.io/heptio-images/ks-guestbook-demo:0.2
        name: guestbook-ui
        ports:
        - containerPort: 80
        resources:
          limits:
            cpu: 100m
            memory: 64Mi
```

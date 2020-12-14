```bash
helm install grafeas grafeas-charts/ --set persistence.enabled=true --set persistence.storageType=postgres --set image.repository=quay.io/kjanania --set image.name=grafeas --set image.tag=resource-details
```

```
oc apply -f https://raw.githubusercontent.com/open-policy-agent/gatekeeper/v3.2.2/deploy/gatekeeper.yaml
```

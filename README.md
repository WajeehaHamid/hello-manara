# hello-manara

Real push-to-deploy demo on Knative via Argo CD.

## Flow

```
git push (this repo)
   └─► GitHub Actions: docker build → push to ghcr.io
        └─► Actions: update k8s/service.yaml image tag → commit back
             └─► Argo CD (in cluster) detects change → kubectl apply
                  └─► Knative creates new revision → pod pulls image → live
```

## Layout

- `main.go` — tiny HTTP server
- `Dockerfile` — distroless Go image
- `k8s/service.yaml` — Knative Service manifest (image tag auto-updated by CI)
- `.github/workflows/build-and-deploy.yaml` — build, push, pin tag

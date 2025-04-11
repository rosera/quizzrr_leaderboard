# Google Compute Engine Reverse Proxy

Nginx redirect to backend compute service

- [x] Amend HOST_IP
- [x] Amend HOST_PORT
- [x] Build + Tag Semantic version
- [x] Push Container Image to Container Image Registry


CloudBuild will pass the HOST_IP and HOST_PORT to Docker build.
Docker build will then replace the NGINX config with the values.
```
gcloud builds submit --config cloudbuild.yaml
```

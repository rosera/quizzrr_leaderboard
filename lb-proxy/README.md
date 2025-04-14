# Google Compute Engine Reverse Proxy

Overview of how to create a Reverse Proxy using Cloud Run VPC Direct.
The backend resource is a Compute Engine instance.

## VPC Direct

Using Cloud Run and Compute Engine, both need to be in the same __REGION__.

-  [ ] Region


## CloudBuild

For the Reverse Proxy example, NGINX is used:

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

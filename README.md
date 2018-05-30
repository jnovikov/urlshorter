# UrlShorter
UrlShorter -- prototype of service for creating short urls from originals
## Description
Prototype of service on Go lang.
# Usage
1) Compile the source code
2) Copy the default_config.yaml to config.yaml and edit it as you need.
3) Run it
## Docker
You can easily run it on docker.
1) Copy the default_config.yaml to config.yaml and edit it as you need.
2) Run `docker build -t "urlshorter" .`
3) Run `docker run -p 8000:8000 urlshorter `
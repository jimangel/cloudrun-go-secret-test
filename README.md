# Cloud Run Secret Manager Tester [WIP]

This is a [~30 line](main.go) HTTP golang app that reads a file. Great for testing and debugging (secret) volume mounts; not great for production.

In the demo below, _the link is publicly available._

**DO NOT PUT SENSITIVE INFORMATION IN THIS SECRET (yet)**

The main goal is ensuring files exist where you intend them to. Once configuration is validated, then look further into securing the environment with IAM and various other layers of protection.

```
# TODO:
# gcloud create secret

# gcloud create secret iam for cloud run (run as service account?)

# cloud create using my public image

```

[TODO: PICTURE]

> Note: The container is small (~8MB) built on `gcr.io/distroless/static:nonroot`

## Build your own image

Install [google/ko](https://github.com/google/ko)r. `ko` makes it easy to build multi-arch images with no major dependancies or configuration.

```
# Set to your image repo that you have configured with `docker push` access
export KO_DOCKER_REPO="your-repo.docker.io/username/path/image"

# Build and push image (--bare ensures that the image name is not adjusted from the set REPO path)
ko publish . --platform=all --bare
```
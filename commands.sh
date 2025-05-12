docker buildx build --platform linux/amd64,linux/arm64 -t dimasenko2000/db-worker:v0.0.4 --push . -f DW-Dockerfile
docker buildx build --platform linux/amd64,linux/arm64 -t dimasenko2000/kw-worker:v0.0.4 --push . -f KW-Dockerfile

docker buildx build --platform linux/amd64 . -f DW-Dockerfile
docker buildx build --platform linux/arm64 . -f DW-Dockerfile
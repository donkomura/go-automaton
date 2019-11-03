FROM alpine:latest
WORKDIR /usr/local/bin/
COPY gopath/bin/$REPO_NAME .
ENTRYPOINT ["/usr/local/bin/$REPO_NAME"]

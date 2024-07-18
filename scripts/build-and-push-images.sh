#!/bin/bash

# Build servers
#
docker build -f httpsserver.Dockerfile -t ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpsserver:latest ./
docker build -f httpserver.Dockerfile -t ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpserver:latest ./

# Build clients
#
docker build -f httpclient.Dockerfile -t ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpclient:latest ./
docker build -f httpsclient.Dockerfile -t ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpsclient:latest ./


# Push servers
#
docker push ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpsserver:latest
docker push ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpserver:latest

# Push clients
#
docker push ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpclient:latest
docker push ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpsclient:latest
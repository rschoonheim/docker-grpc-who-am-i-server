# Docker Grpc - Who am I?

Simple gRPC server and client in GoLang that can be used to test the connection between a client and a server in a
Docker environment.

## How to use

### Server - HTTP Mode
To run the server in HTTP mode, you can use the following example:
```yaml
networks:
  default:
    name: docker-grpc-who-am-i-network
    driver: bridge

services:

  # HTTP Server
  # ------------------------------------------
  #
  http-server:
    image: ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpserver:latest
    container_name: http-server
    ports:
      - "50051:50051"
    networks:
      - default
    restart: always
```

### Server - HTTPs Mode

#### Generating server and client certificates
To generate the certificates, you can use the following script
```bash
```

#### Running the server
```yaml
networks:
  default:
    name: docker-grpc-who-am-i-network
    driver: bridge


services:

  # HTTPS Server
  # ------------------------------------------
  #
  https-server:
    image: ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpsserver:latest
    container_name: https-server
    ports:
      - "50052:50052"
    networks:
      - default
    restart: always
```
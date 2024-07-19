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
      - "8080:8080"
    networks:
      - default
    restart: always


  # HTTP Client
  # ------------------------------------------
  #
  http-client:
    image: ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpclient:latest
    container_name: http-client
    networks:
      - default
    restart: always
    environment:
      - WHOAMI_SERVER_ADDRESS=http-server:8080
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
      - "8080:8080"
    networks:
      - default
    restart: always

  # HTTPS Client
  # ------------------------------------------
  #
  https-client:
    image: ghcr.io/rschoonheim/docker-grpc-who-am-i-server/httpsclient:latest
    container_name: https-client
    networks:
      - default
    restart: always
    environment:
      - WHOAMI_SERVER_ADDRESS=https-server:8080
```
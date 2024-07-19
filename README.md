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
#!/bin/bash

# Check if `.tls` directory exists, if not create it.
#
if [ ! -d ./.tls ]; then
  mkdir .tls
fi

cd .tls

mkdir certs crl newcerts private
chmod 700 private
touch index.txt
echo 1000 > serial

cp ../config/root.cnf openssl.cnf

# Generate private key for Certificate Authority (CA)
#
openssl genrsa -out private/ca.key.pem 4096
chmod 400 private/ca.key.pem

# Generate root certificate for Certificate Authority (CA)
#
openssl req -config openssl.cnf -key private/ca.key.pem -new -x509 \
    -days 7300 -sha256 -extensions v3_ca -out certs/ca.cert.pem
chmod 444 certs/ca.cert.pem

# Verify root certificate
#
openssl verify -CAfile certs/ca.cert.pem certs/ca.cert.pem

# Intermediate CA
#
mkdir intermediate
cd intermediate

mkdir certs crl csr newcerts private
chmod 700 private
touch index.txt
echo 1000 > serial
echo 1000 > crlnumber

pwd

cp ../../config/intermediate.cnf openssl.cnf

# Generate private key for Intermediate Certificate Authority (CA)
#
openssl genrsa -out private/intermediate.key.pem 4096
chmod 400 private/intermediate.key.pem

# Generate CSR for Intermediate Certificate Authority (CA)
#
cd ../
openssl req -config intermediate/openssl.cnf -new -sha256 \
    -key intermediate/private/intermediate.key.pem \
    -out intermediate/csr/intermediate.csr.pem

# Sign Intermediate Certificate Authority (CA) CSR with Root CA
#
openssl ca -config openssl.cnf -extensions v3_intermediate_ca \
    -days 3650 -notext -md sha256 -in intermediate/csr/intermediate.csr.pem \
    -out intermediate/certs/intermediate.cert.pem
chmod 444 intermediate/certs/intermediate.cert.pem

# Verify Intermediate Certificate Authority (CA) certificate
#
openssl verify -CAfile certs/ca.cert.pem intermediate/certs/intermediate.cert.pem

# Create certificate chain file
#
cat intermediate/certs/intermediate.cert.pem certs/ca.cert.pem > \
    intermediate/certs/ca-chain.cert.pem
chmod 444 intermediate/certs/ca-chain.cert.pem

# Server key
#
openssl genrsa -out intermediate/private/server.key.pem 2048

# Server signing request
#
openssl req -config intermediate/openssl.cnf \
    -key intermediate/private/server.key.pem \
    -new -sha256 -out intermediate/csr/server.csr.pem \
    -addext "subjectAltName = DNS:server"

# Server cert
#
openssl ca -config intermediate/openssl.cnf -extensions server_cert \
    -days 375 -notext -md sha256 -in intermediate/csr/server.csr.pem \
    -out intermediate/certs/server.cert.pem
chmod 444 intermediate/certs/server.cert.pem

# Verify server certificate
#
openssl verify -CAfile intermediate/certs/ca-chain.cert.pem \
    intermediate/certs/server.cert.pem

# client key
#
openssl genrsa -out intermediate/private/client.key.pem 2048

# client signing request
#
openssl req -config intermediate/openssl.cnf \
    -key intermediate/private/client.key.pem \
    -new -sha256 -out intermediate/csr/client.csr.pem \
    -addext "subjectAltName = DNS:client"

# client cert
#
openssl ca -config intermediate/openssl.cnf -extensions usr_cert \
    -days 375 -notext -md sha256 -in intermediate/csr/client.csr.pem \
    -out intermediate/certs/client.cert.pem
chmod 444 intermediate/certs/client.cert.pem

# Verify client certificate
#
openssl verify -CAfile intermediate/certs/ca-chain.cert.pem \
    intermediate/certs/client.cert.pem
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
    environment:
      - CA_CERT_PATH=/tls/intermediate/certs/intermediate.cert.pem
      - TLS_CERT_PATH=/tls/intermediate/certs/server.cert.pem
      - TLS_KEY_PATH=/tls/intermediate/private/server.key.pem
    volumes:
      - ./tls:/tls

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
      - CA_CERT_PATH=/tls/intermediate/certs/intermediate.cert.pem
      - TLS_CERT_PATH=/tls/intermediate/certs/client.cert.pem
      - TLS_KEY_PATH=/tls/intermediate/private/client.key.pem
    volumes:
      - ./tls:/tls
```
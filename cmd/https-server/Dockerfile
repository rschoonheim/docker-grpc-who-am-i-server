FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

FROM alpine:3.9 AS final

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

ENTRYPOINT ["./server"]
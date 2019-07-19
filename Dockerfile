FROM golang:1.12.5 as build
WORKDIR /build
COPY go.mod /build
COPY go.sum /build
RUN go mod download
COPY . /build
ENV CGO_ENABLED=0
RUN  go build -o server

FROM ubuntu:latest
RUN apt update && apt install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY  --from=build /build/server /app/
COPY  --from=build /build/.env /app/.env
ENTRYPOINT /app/server 
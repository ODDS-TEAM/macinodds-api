FROM golang:1.12.5 as build
COPY . /build
WORKDIR /build
RUN go get
RUN go build -o server

FROM ubuntu:latest
RUN apt update && apt install iy ca-certificates && rm -rf /var/lib/apt/lists/*
COPY  --from=build /build/server /app/
COPY  --from=build /build/.env /app/.env
ENTRYPOINT /app/server 
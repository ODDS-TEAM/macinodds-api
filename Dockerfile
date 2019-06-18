FROM golang:1.12.5 as build
COPY . /build
WORKDIR /build
RUN go get
RUN go buid -o server

FROM ubuntu:latest
COPY  --from=build /build/server /app/
ENTRYPOINT /app/server
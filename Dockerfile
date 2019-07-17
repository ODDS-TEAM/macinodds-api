FROM golang:1.12.5 as build
COPY . /build
WORKDIR /build
RUN go get
RUN go build -o server

FROM ubuntu:latest
COPY  --from=build /build/server /app/
COPY  --from=build /build/.env /app/.env
ENTRYPOINT /app/server
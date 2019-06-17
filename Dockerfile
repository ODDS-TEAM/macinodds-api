FROM ubuntu:latest

ADD server /app/

ENTRYPOINT /app/server

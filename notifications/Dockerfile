FROM ubuntu:22.04

ADD ./bin/app /app
ADD ./config.yaml /config.yaml

RUN apt-get update && apt-get install ca-certificates -y && rm -rf /var/cache/apk/*

CMD ["/app"]
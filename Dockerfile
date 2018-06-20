FROM golang:latest

VOLUME /data
ADD md5x /

CMD ["/md5x", "-dir", "/data", "-out", "o.json"]
FROM ubuntu:16.04

COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 4500

COPY main /main
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
FROM scratch
COPY bin/firewall-exporter /firewall-exporter
USER 999
ENTRYPOINT ["/firewall-exporter"]

EXPOSE 9080

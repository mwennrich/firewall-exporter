
FROM golang:1.22-alpine as builder
RUN apk add make binutils
COPY / /work
WORKDIR /work
RUN make firewall-exporter

FROM scratch
COPY --from=builder /work/bin/firewall-exporter /firewall-exporter
USER 999
ENTRYPOINT ["/firewall-exporter"]

EXPOSE 9080

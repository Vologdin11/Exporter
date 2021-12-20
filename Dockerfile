FROM odise/busybox-curl
COPY /stp-exporter .
ENTRYPOINT ["/stp-exporter"]
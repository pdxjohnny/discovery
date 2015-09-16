FROM busybox
ADD ./discovery_linux-amd64 /app
CMD ["/app"]

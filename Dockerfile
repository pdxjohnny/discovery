FROM busybox
ADD ./frontend_linux-amd64 /app
CMD ["/app"]

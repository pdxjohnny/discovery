Frontend
---

A web service to receive files

It uses docker to compile the binaries and the main Dockerfile adds the linux
binary to the busybox image to create an extremely small final image

Building
---

```bash
go build -o frontend_linux-amd64 -tags netgo *.go
# Or
./script/build
# My Favorite
SKIP_BUILD=1 SKIP_IMAGE=1 ./script/build -osarch="linux/amd64"
```
> `-tags netgo` will help you achieve static binaries :)

Running
---

```bash
./frontend_linux-amd64
docker run --rm -ti pdxjohnny/frontend
```

Changing The Name
---

```bash
./script/change-name $GITHUB_USERNAME $PROJECT_NAME
```


- John Andersen

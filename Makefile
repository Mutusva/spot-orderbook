check_install:
    which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
   GO111MODULE=on swagger generate spec -o ./swagger.yaml --scan-models

redisDocker:
   docker run --name redis -d redis
   docker run -it --rm --link redis redis redis-cli -h redis

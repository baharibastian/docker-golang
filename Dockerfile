FROM golang:latest

WORKDIR /var/www

RUN go get -u github.com/golang/dep/cmd/dep \
    && go get -u github.com/go-redis/redis \
    && dep ensure -v \
    && go build \
    && go install \
    && cp shared/config/docker.yml shared/config/default.yml
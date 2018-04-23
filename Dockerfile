FROM alpine:latest

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

ADD ./bin/git-pull-linux-amd64 /git-pull

CMD ["/git-pull"]

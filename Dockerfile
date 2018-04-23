FROM alpine:latest

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

ADD ./script.sh /script.sh

CMD ["/script.sh"]

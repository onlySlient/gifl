FROM alpine:latest

ENV VERSION 0.1

WORKDIR /apps

COPY bin/app /apps/app

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ENV LANG C.UTF-8

EXPOSE 9999

ENTRYPOINT ["/apps/app"]
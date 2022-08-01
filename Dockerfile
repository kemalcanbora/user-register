FROM public.ecr.aws/bitnami/golang:latest as goBuilder
LABEL maintainer="kemalcanbora@gmail.com"
LABEL description="Rollic"

RUN mkdir app
COPY . /app/
WORKDIR /app

RUN apt update && apt install -y apt-transport-https ca-certificates sqlite3
RUN apt-get update && apt-get install -y supervisor openssh-server

RUN mkdir -p /var/log/supervisor
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
EXPOSE 8080

RUN go get -d -v ./...
RUN go build cmd/main.go


CMD ["/usr/bin/supervisord"]
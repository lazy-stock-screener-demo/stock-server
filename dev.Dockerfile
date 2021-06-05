FROM golang:latest as builder
LABEL maintainer_email="universetennis@gmail.com"
LABEL maintainer="Daniel Lin"
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN useradd --user-group --create-home --shell /bin/bash app
RUN chown -R app:app /home/app
WORKDIR /home/app
USER app
COPY . .
USER root
RUN chmod 777 go.sum
USER app
RUN go mod download
EXPOSE 80
CMD ["go","run", "main.go"]
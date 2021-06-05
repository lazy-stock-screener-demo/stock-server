FROM golang:alpine as builder

LABEL maintainer_email="universetennis@gmail.com"
LABEL maintainer="Daniel Lin"

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN addgroup -S app && adduser -S app -G app -h /home/app -s /bin/bash 

WORKDIR /home/app
USER app
COPY --chown=app:app go.mod go.sum main.go ./
COPY --chown=app:app ./pkg ./pkg
RUN go mod download
RUN go build -o server

FROM alpine as release
WORKDIR /home/app
COPY --from=builder /home/app/server ./
EXPOSE 80
CMD ["/home/app/server"]
FROM golang:1.18.0-alpine

WORKDIR /app
ADD . /app
RUN go build
EXPOSE 8080

ENTRYPOINT [ "/app/app" ]
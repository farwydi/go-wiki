FROM golang:1.11-alpine AS binary

ADD . /app
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o go-wiki

FROM alpine

COPY --from=binary /app/go-wiki /go/bin/go-wiki

EXPOSE 8080
CMD ["/go/bin/go-wiki"]
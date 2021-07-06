FROM golang:alpine AS api_dev

EXPOSE 8080
WORKDIR /go/src/app
COPY api/ .
RUN go clean --modcache
RUN go mod download
RUN export GO111MODULE=on && go build -o api main.go

ENTRYPOINT /go/src/app/api
FROM golang:alpine AS api_dev

EXPOSE 8080
WORKDIR /go/src/app
COPY api/go.mod api/go.sum ./
RUN go mod download
COPY api/ .
RUN export GO111MODULE=on && go build -o api main.go

ENTRYPOINT /go/src/app/api
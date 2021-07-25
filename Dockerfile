FROM golang:alpine AS api_dev

EXPOSE 8080
WORKDIR /go/src/elastic_books/api
COPY api/go.mod .
COPY api/go.sum .
RUN go mod download -x
COPY api/ .

CMD ["go", "run", "main.go"]
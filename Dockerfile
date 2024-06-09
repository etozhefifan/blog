FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build blog.go

FROM alpine:3
COPY --from=builder /app/blog /app/template.html /bin/
ENTRYPOINT ["/bin/blog"]

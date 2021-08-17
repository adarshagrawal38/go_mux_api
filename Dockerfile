From golang:alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .
FROM alpine:3.14
COPY --from=builder /app/bin/main .
CMD ["./main"]
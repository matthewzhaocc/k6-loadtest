FROM golang:latest AS builder

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN GOPATH= go build -o /server .

FROM alpine:latest

COPY --from=builder /server .

ENV GIN_MODE=release

CMD ["./server"]
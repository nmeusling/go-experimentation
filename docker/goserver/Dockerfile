FROM golang:1.25-alpine AS builder

ENV GOOS=linux
ENV GOARCH=amd64

COPY ["go.mod", "main.go", "/go"]
RUN go build -o goserver

# FROM debian:stable-slim
FROM scratch
ENV PORT=8991
COPY --from=builder /go/goserver /bin/goserver
ENTRYPOINT ["/bin/goserver"]
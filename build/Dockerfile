FROM golang:alpine as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN ls /app
RUN apk update && apk add ca-certificates openssl && update-ca-certificates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/api-binary ./cmd/api/main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /app/
COPY --from=builder /bin/api-binary /bin/api-binary
COPY --from=builder --chown=nonroot /app/config /app/config
ENTRYPOINT ["/bin/api-binary"]
FROM golang:1.17-alpine as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app-binary ./cmd/main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /app/
COPY --from=builder /bin/app-binary /app/app-binary
COPY --from=builder --chown=nonroot /app/config /app/config
ENTRYPOINT ["/app/app-binary"]
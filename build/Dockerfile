FROM golang:alpine3.12 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN apk update && apk add ca-certificates openssl && update-ca-certificates

#RUN wget -O install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh
#RUN chmod +x install.sh && sh install.sh 
#COPY bin/air /bin/air

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app-binary ./cmd/main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /app/
COPY --from=builder /bin/app-binary /app/app-binary
#COPY --from=builder /app/bin/air /app/air
COPY --from=builder --chown=nonroot /app/config /app/config
ENTRYPOINT ["/app/app-binary"]
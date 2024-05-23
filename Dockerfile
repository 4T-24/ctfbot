FROM golang:1.22-alpine AS builder

RUN apk --no-cache add ca-certificates && update-ca-certificates

USER 1000
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN --mount=type=ssh go mod download && go mod verify
COPY . .
RUN go build -o /app/ctfbot /app/cmd/bot/main.go

FROM scratch
USER 1000
COPY --from=builder /app/ctfbot .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY sql /sql
ENTRYPOINT ["/ctfbot"]
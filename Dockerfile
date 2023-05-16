FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o main .

FROM gcr.io/distroless/static-debian11 AS final
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]
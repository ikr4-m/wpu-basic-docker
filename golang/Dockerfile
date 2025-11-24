FROM docker.io/library/golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -C="." -ldflags="-w -s -linkmode external -extldflags -static" -o="app";

FROM scratch 
COPY --from=builder /app/app /app
CMD ["/app"]

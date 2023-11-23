FROM golang:alpine AS builder 
RUN apk --no-cache add tzdata
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o tiendanube cmd/consumer.go

FROM scratch AS final
COPY --from=builder /app/tiendanube ./
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV TZ=America/Sao_Paulo

ENTRYPOINT ["./tiendanube"]
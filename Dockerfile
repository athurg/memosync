FROM golang:1.20 as build
ADD . /app
RUN CGO_ENABLED=0 go build -C /app

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /app/memosync /memosync
ENTRYPOINT ["/memosync"]

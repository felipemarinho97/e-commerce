FROM golang:1.17.3-alpine3.14 AS builder

WORKDIR /app 

COPY ./ .
RUN GO111MODULE=on go mod download 

RUN CGO_ENABLED=0 GOOS=linux go build -o ecommerce-rest


FROM alpine:3.14

COPY --from=builder /app/ecommerce-rest /ecommerce-rest

CMD ["/ecommerce-rest"]
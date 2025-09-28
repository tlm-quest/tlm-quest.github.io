FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .

RUN <<EOF
go mod tidy 
go build -o hawaiian-test
EOF

FROM alpine:3.22.1
WORKDIR /app
COPY --from=builder /app/hawaiian-test .

CMD ["./hawaiian-test"]
FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o /clouflareupdatedns

FROM scratch

USER 1000:1000

COPY --from=builder /clouflareupdatedns .

ENTRYPOINT ["/clouflareupdatedns"]

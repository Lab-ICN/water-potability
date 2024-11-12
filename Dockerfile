FROM golang:1.22.9-alpine3.20 as builder

WORKDIR /tmp/build
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o server ./cmd/app

FROM gcr.io/distroless/static-debian12
COPY --from=builder /tmp/build/server /

CMD ["/server"]

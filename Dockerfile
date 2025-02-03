FROM golang as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /build

COPY go.mod go.sum /.

RUN go mod download

COPY . .

RUN go build -a -v -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /build/main .

EXPOSE 3000

ENTRYPOINT ["./main"]


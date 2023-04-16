FROM golang:latest as builder

RUN mkdir /build
WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD . .

RUN CGO_ENABLED=0 go build -o users .

FROM alpine:latest

COPY --from=builder /build/users .

ENTRYPOINT [ "./users" ]


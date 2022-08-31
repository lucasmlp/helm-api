# STAGE 0: Contruct builder image
FROM golang:1.18-stretch as builder

WORKDIR /helm_api

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

COPY . .

RUN go mod download

RUN go build -a -o cmd ./cmd

# STAGE 2: Build final image
FROM alpine as final
COPY --from=builder /helm_api/cmd/cmd /go/bin/cmd
RUN apk add -U --no-cache ca-certificates
ENTRYPOINT /go/bin/cmd

EXPOSE 8080

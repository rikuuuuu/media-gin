FROM golang:1.16.3-alpine as build

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -mod=mod -v -o server

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    automake \
    make \
    git \
    bash \
    curl \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /app/server /server
COPY ./config /config

CMD ["/server"]
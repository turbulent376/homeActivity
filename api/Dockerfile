FROM golang:1.16.8-stretch

ENV GO111MODULE=on

WORKDIR /src

COPY ./go.mod ./go.sum ./api/

WORKDIR /src/api

ARG SPACE_USER
ARG SPACE_TOKEN

RUN go env -w GOPRIVATE="git.jetbrains.space/*" && \
    echo -e "machine git.jetbrains.space\nlogin ${SPACE_USER}\npassword ${SPACE_TOKEN}" > ~/.netrc && \
    apt-get update && apt-get install make && \
    go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.0 && \
    go get -u github.com/swaggo/swag/cmd/swag

RUN go mod download

COPY . ./

WORKDIR /src/api

RUN make swagger

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN make build

FROM alpine:3.12.3

# install grpcurl - curl for gRPC
RUN mkdir /tmp/grpcurl && \
    wget https://github.com/fullstorydev/grpcurl/releases/download/v1.8.1/grpcurl_1.8.1_linux_x86_64.tar.gz -O /tmp/grpcurl/grpcurl.tar.gz  && \
    tar -xzf /tmp/grpcurl/grpcurl.tar.gz -C /tmp/grpcurl && \
    cp /tmp/grpcurl/grpcurl /usr/local/bin/grpcurl && \
    rm -f /tmp/grpcurl/*

WORKDIR /usr/local/bin/api

ENV FOCROOT="/usr/local/bin"

COPY --from=0 /src/api/bin/main ./bin/main
COPY ./config.yml ./config.yml

ENTRYPOINT ["./bin/main"]

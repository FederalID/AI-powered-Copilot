# build stage
FROM golang:alpine AS builder
ADD . /go/src/github.com/feiskyer/openai-copilot
RUN cd /go/src/github.com/feiskyer/openai-copilot && \
    apk update && apk add --no-cache gcc musl-dev openssl && \
    C
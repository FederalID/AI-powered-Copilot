# build stage
FROM golang:alpine AS builder
ADD . /go/src
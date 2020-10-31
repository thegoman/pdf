# This file is a template, and might need editing before it works on your project.
FROM golang:1.15 AS builder

WORKDIR /go/src/gitlab.com/thegoman/pdf

COPY . .
RUN go get -u golang.org/x/lint/golint
RUN go mod download

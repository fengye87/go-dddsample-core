FROM golang:1.17.2-alpine3.14 AS builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY cargo/ cargo/
RUN --mount=type=cache,target=/root/.cache/go-build go build cargo/main.go


FROM alpine:3.14.2

COPY --from=builder /workspace/main /cargo
ENTRYPOINT ["/cargo"]

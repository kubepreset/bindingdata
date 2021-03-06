# Build the manager binary
FROM docker.io/library/golang:alpine as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o bindingdata main.go

FROM docker.io/library/alpine:latest
WORKDIR /
COPY --from=builder /workspace/bindingdata .
USER 65532:65532

ENTRYPOINT ["/bindingdata"]

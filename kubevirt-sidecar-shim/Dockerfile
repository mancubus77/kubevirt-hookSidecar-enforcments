# Multi-stage build for kubevirt-sidecar-shim
FROM golang:1.23-alpine AS builder

# Install git and ca-certificates (needed for go modules)
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /workspace

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the onDefineDomain binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o onDefineDomain .

# Final stage - minimal runtime image
FROM registry.redhat.io/container-native-virtualization/sidecar-shim-rhel9@sha256:480114e10e4b16d963338caa81b1aab21be1a4064a954c11a8c2f9e021ecc9e1 

COPY --from=builder /workspace/onDefineDomain /usr/bin/onDefineDomain
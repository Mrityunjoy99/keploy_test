# ------------- Stage: Build -------------------
FROM public.ecr.aws/docker/library/golang:1.21.1-alpine3.18 as build

# Install git for downloading private Go dependencies
RUN apk add --no-cache git=2.40.1-r0
# Install tzdata for timezone
RUN apk add --no-cache tzdata

RUN apk --no-cache add curl

RUN apk --no-cache add bash

# Set workdir as build
WORKDIR /build

RUN curl -o ca.crt https://raw.githubusercontent.com/keploy/keploy/main/pkg/proxy/asset/ca.crt

RUN curl -o setup_ca.sh https://raw.githubusercontent.com/keploy/keploy/main/pkg/proxy/asset/setup_ca.sh

# # Give execute permission to the setup_ca.sh script
RUN chmod +x setup_ca.sh

# Copy go.mod and go.sum to get dependencies
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy the rest of the code, create .dockerignore to exclude files
COPY ./ ./

# Build the binary, './bin/<binary_name>'
RUN CGO_ENABLED=0 go build -tags timetzdata -o ./bin/sample_project cmd/main.go

# ------------- Stage: Release -----------------
FROM public.ecr.aws/docker/library/alpine:3.18 AS release

# Switch to non-root user
USER 1001

# Set workdir as app
WORKDIR /app

# Set environment variables
ENV GIN_MODE=release
ENV APP_ID=local

# Copy the binary from the build stage (use the same binary name)
COPY --chown=1001:1001 --from=build /build/bin/sample_project .
COPY --from=build /build/ca.crt .
COPY --from=build /build/setup_ca.sh .
USER root
RUN chmod +x ./setup_ca.sh
USER 1001
# Provide default command, this can be overridden by providing a command in the values-<env>.yaml file
CMD ["./sample_project", "start"]
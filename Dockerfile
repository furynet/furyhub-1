#
# Build image: docker build -t fanfury/fanfury .
#
FROM golang:1.18-alpine3.16 as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev bash linux-headers eudev-dev

WORKDIR /gridiron

# Add source files
COPY . .

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

RUN make build

# ----------------------------

FROM alpine:3.16

# p2p port
EXPOSE 26656
# rpc port
EXPOSE 26657
# metrics port
EXPOSE 26660

COPY --from=builder /gridiron/build/ /usr/local/bin/
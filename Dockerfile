# To build the Desmos image, just run:
# > docker build -t desmos .
#
# Simple usage with a mounted data directory:
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.desmos:/root/.desmos desmos desmos init
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.desmos:/root/.desmos desmos desmos start
#
# If you want to run this container as a daemon, you can do so by executing
# > docker run -td -p 26657:26657 -p 26656:26656 -v ~/.desmos:/root/.desmos --name desmos desmos
#
# Once you have done so, you can enter the container shell by executing
# > docker exec -it desmos bash
#
# To exit the bash, just execute
# > exit

###############################################################################
# Build go-cosmwasm
###############################################################################
FROM rustlang/rust:nightly as build-cosmwasm

# Install build dependencies
RUN apt-get update
RUN apt install -y clang gcc g++ zlib1g-dev libmpc-dev libmpfr-dev libgmp-dev
RUN apt install -y build-essential cmake git

# Install repository
RUN git clone https://github.com/confio/go-cosmwasm /cosmwasm

# Compile go-cosmwasm
WORKDIR /cosmwasm
RUN make build-rust-release

###############################################################################
# Build desmos
###############################################################################
FROM golang:1.15 AS build-desmos

# Set working directory for the build and copy the files over
WORKDIR /desmos
COPY . .

# Install Desmos, remove packages
RUN make build-linux

###############################################################################
# Create final image
###############################################################################
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Install bash
RUN apk add --no-cache bash

# Copy over binaries from the build-desmos
COPY --from=build-desmos /desmos/build/desmos /usr/bin/desmos
COPY --from=build-cosmwasm /cosmwasm/api/libwasmvm.so /usr/lib/libgo_cosmwasm.so

EXPOSE 26656 26657 1317 9090

# Run desmos by default, omit entrypoint to ease using container with desmos
CMD ["desmos"]

FROM golang:latest
EXPOSE 42000

# Set go bin which doesn't appear to be set already.
ENV GOPATH /go

ARG REPO_PATH=github.com/garritfra/blockchain-project

# build directories
RUN mkdir /app
RUN mkdir -p /go/src/${REPO_PATH}
ADD . /go/src/${REPO_PATH}
WORKDIR /go/src/${REPO_PATH}

# Build my app
RUN go install
RUN go build
CMD ["/go/bin/blockchain-project"]


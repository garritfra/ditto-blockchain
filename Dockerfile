FROM golang:alpine

# Set go bin which doesn't appear to be set already.
ENV GOPATH /go

ARG REPO_PATH=github.com/garritfra/blockchain-project

# build directories
RUN mkdir /app && mkdir -p /go/src/${REPO_PATH}
ADD . /go/src/${REPO_PATH}
WORKDIR /go/src/${REPO_PATH}

# Build my app
RUN go install && go build
EXPOSE 42000
CMD ["/go/bin/blockchain-project"]


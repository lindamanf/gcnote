FROM golang:1.15.6

ENV SRC_DIR=/go/src/github.com/lindamanf/gcnote
ENV GOBIN=/go/bin

RUN mkdir -p $SRC_DIR
WORKDIR $SRC_DIR

RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs;
RUN go get github.com/ramya-rao-a/go-outline;
RUn go get golang.org/x/tools/gopls;

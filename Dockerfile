FROM golang:latest
RUN mkdir -p $GOPATH/src/github.com/url-shortener
ADD . $GOPATH/src/github.com/url-shortener
WORKDIR $GOPATH/src/github.com/url-shortener
RUN go install .
CMD ["url-shortener"]
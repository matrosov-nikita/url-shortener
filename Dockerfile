FROM golang:latest
RUN mkdir -p $GOPATH/src/github.com/matrosov-nikita/url-shortener
ADD . $GOPATH/src/github.com/matrosov-nikita/url-shortener
WORKDIR $GOPATH/src/github.com/matrosov-nikita/url-shortener
RUN go install .
CMD ["url-shortener"]
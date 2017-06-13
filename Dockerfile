FROM golang

ADD . /go/src/lowes.com/smweathe/recommendations-services
RUN cd /go/src/lowes.com/smweathe/recommendations-services; go get
RUN go install lowes.com/smweathe/recommendations-services

EXPOSE 8080

ENTRYPOINT /go/bin/recommendations-services

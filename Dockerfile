FROM golang:1.6.3-onbuild

ENV GOBIN /go/bin
RUN true && \
	mkdir -p /go/src/github.com/callowaylc/datadog && \
	cp -rf . /go/src/github.com/callowaylc/datadog && \
	go install ./cmd/datadog

ENTRYPOINT ["/go/bin/datadog"]
CMD ["-h"]

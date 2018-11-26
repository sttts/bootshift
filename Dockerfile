FROM golang:1.11.1
COPY . /go/src/github.com/openshift/bootshift
RUN cd /go/src/github.com/openshift/bootshift && GOPATH=/go make build

FROM scratch
COPY --from=0 /go/src/github.com/openshift/bootshift/bootshift /bootshift
ENTRYPOINT /bootshift
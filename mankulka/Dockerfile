FROM golang:latest AS builder
RUN mkdir containers
COPY ../hello_openshift.go containers/
RUN go build -o containers/hello_openshift containers/hello_openshift.go
FROM registry.access.redhat.com/ubi8/ubi:latest
RUN mkdir containers
COPY --from=builder go/containers/hello_openshift containers/
ENTRYPOINT ["./containers/hello_openshift"]

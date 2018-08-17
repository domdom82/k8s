FROM ubuntu:18.04

ENV GOROOT "/usr/lib/go-1.10"
ENV PATH "${PATH}:$GOROOT/bin"

RUN mkdir -p /root/go/src/github.com/domdom82/k8s/echo && \
    apt-get update && \
    apt-get install -y golang-1.10 \
                       golang-glide

COPY src /root/go/src/github.com/domdom82/k8s/echo/src

RUN cd /root/go/src/github.com/domdom82/k8s/echo/src && \
    glide update && \
    go build main.go && \
    mv main /

EXPOSE 8080

CMD /main
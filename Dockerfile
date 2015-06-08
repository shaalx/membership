FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get github.com/shaalx/merbership
RUN go install github.com/shaalx/merbership
RUN ls
EXPOSE 80
CMD ["/gopath/app/bin/merbership"]

FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

#RUN go get github.com/shaalx/membership
#RUN go install github.com/shaalx/membership
RUN go build
EXPOSE 80
CMD ["/gopath/app/bin/membership"]

FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get -u github.com/Unknwon/macaron
RUN go get -u labix.org/v2/mgo/bson
RUN mkdir -p $GOPATH/src/github.com/shaalx/membership;cd $GOPATH/src/github.com/shaalx/membership;git init;git remote add origin https://github.com/shaalx/membership;git fetch origin devm:devm;go install
EXPOSE 80
CMD ["/gopath/app/bin/membership"]

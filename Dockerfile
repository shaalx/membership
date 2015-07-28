FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN mkdir -p $GOPATH/bin $GOPATH/pkg $GOPATH/src
RUN go get -u github.com/Unknwon/macaron
RUN go get -u labix.org/v2/mgo/bson
# RUN mkdir -p $GOPATH/src/github.com/shaalx/membership;cd $GOPATH/src/github.com/shaalx/membership;git init;git remote add origin https://github.com/shaalx/membership;git fetch origin devm:devm;go install
RUN mkdir -p $GOPATH/src/github.com/shaalx/membership;cd $GOPATH/src/github.com/shaalx/membership;git init;git remote add member https://github.com/shaalx/membership;git fetch member devm:devm;git checkout devm;ls;go build -o membership;cp membership $GOPATH/bin/
EXPOSE 80
CMD ["/gopath/app/bin/membership"]

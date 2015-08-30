FROM golang

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/
RUN mkdir -p src bin pkg
RUN go get github.com/shaalx/membership
RUN go build -o bookmark
EXPOSE 80
CMD ["/gopath/app/bookmark"]
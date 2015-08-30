FROM golang

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/
RUN go build -o bookmark
EXPOSE 80
CMD ["/gopath/app/bookmark"]
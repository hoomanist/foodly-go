FROM golang

RUN mkdir /opt/app
ADD . /opt/app
WORKDIR /opt/app
RUN go build
CMD ["./foodly"]

FROM golang:1.14

WORKDIR /opt/foodly
COPY . .

RUN go build 

CMD ["./foodly"]

FROM golang:latest
ENV GO111MODULE=on

ADD . /sw-sys/api-service
WORKDIR /sw-sys/api-service/

COPY go.mod .
COPY go.sum .

EXPOSE 8890
RUN go build

CMD ["./api-service"]

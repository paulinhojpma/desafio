FROM golang:latest
ENV GO111MODULE=on

ADD . /backend/desafio
WORKDIR /backend/desafio/

COPY go.mod .
COPY go.sum .

EXPOSE 8890
RUN go build

CMD ["./desafio"]

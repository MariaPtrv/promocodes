FROM golang:1.23

RUN go version
ENV GOPATH=/

COPY ./ ./ 

COPY go.mod go.sum ./

RUN go mod download

RUN apt-get update
RUN apt-get -y install postgresql-client


RUN go build -o admin ./cmd/main.go

RUN git clone https://github.com/vishnubob/wait-for-it.git

CMD ["./admin"]
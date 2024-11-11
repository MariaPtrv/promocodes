FROM golang:1.23

RUN go version
ENV GOPATH=/

COPY ./ ./ 

COPY go.mod go.sum ./

RUN go mod download

RUN apt-get update
RUN apt-get -y install postgresql-client


RUN go build -o admin ./cmd/main.go



RUN apt-get update && \
    apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent

RUN apt-get update
RUN apt-get install wget
RUN dpkg -l | grep wget
RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-arm64.deb
RUN dpkg -i migrate.linux-arm64.deb



RUN if [ -e "wait-for-it" ];then rm -rf "wait-for-it" ; fi 
RUN git clone https://github.com/vishnubob/wait-for-it.git



COPY migrations migrations

CMD ["./admin"]
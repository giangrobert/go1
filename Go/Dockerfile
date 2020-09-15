FROM golang:latest

RUN apt-get update
RUN apt-get install vim -y
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux
RUN go get github.com/cespare/reflex

EXPOSE 8080

CMD make watch
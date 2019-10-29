FROM golang:alpine
RUN mkdir /cinsear
RUN mkdir -p /media/images/users
RUN mkdir -p /media/images/films
RUN mkdir -p /log
WORKDIR /cinsear
COPY . .
WORKDIR /cinsear/cmd
RUN go build -o server/main ./server
RUN go build -o database/initDB ./database
WORKDIR /cinsear/cmd/server 
CMD ["/cinsear/cmd/server/main"]
WORKDIR /cinsear/cmd/database
CMD  ["/cinsear/cmd/database/initDB"]
EXPOSE 8000

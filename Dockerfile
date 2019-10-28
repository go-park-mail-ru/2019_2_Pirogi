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
CMD ["/cinsear/cmd/server/main", "/cinsear/cmd/database/initDB"]
EXPOSE 8000
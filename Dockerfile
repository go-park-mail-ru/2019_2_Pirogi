FROM golang:alpine
RUN mkdir /cinsear
RUN mkdir -p /media/images/users
RUN mkdir -p /media/images/films
RUN mkdir -p /log
WORKDIR /cinsear
COPY . .
WORKDIR /cinsear/cmd/server
RUN go build -o main .
CMD ["/cinsear/cmd/server/main"]
EXPOSE 8000
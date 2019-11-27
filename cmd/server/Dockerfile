FROM golang:alpine
RUN mkdir /cinsear
RUN mkdir -p /media/images/users
RUN mkdir -p /media/images/films
RUN mkdir -p /log
WORKDIR /cinsear
COPY . .
WORKDIR /cinsear/cmd/server
RUN go build -mod=vendor -o main .
WORKDIR /cinsear
CMD ["/cinsear/cmd/server/main"]
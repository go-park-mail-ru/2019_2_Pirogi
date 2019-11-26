FROM golang:alpine
RUN mkdir /sessions
WORKDIR /sessions
COPY . .
WORKDIR /sessions/cmd/sessions
RUN go build -mod=vendor -o main .
WORKDIR /sessions
CMD ["/sessions/cmd/sessions/main"]
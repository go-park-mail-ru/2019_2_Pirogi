FROM golang:alpine
RUN mkdir /users
WORKDIR /users
COPY . .
WORKDIR /users/cmd/users
RUN go build -mod=vendor -o main .
WORKDIR /users
CMD ["/users/cmd/users/main"]
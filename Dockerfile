FROM golang:latest as builder
RUN mkdir /app
ADD . /app/
RUN go get github.com/gorilla/websocket
WORKDIR /app
RUN go build -o main .

FROM golang:latest
RUN mkdir /app
COPY --from=builder app /app
CMD ["/app/main"]
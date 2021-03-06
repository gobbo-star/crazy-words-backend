FROM golang:latest as builder
RUN go get github.com/gorilla/websocket
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .

FROM golang:latest
RUN mkdir /app
COPY --from=builder app/main /app/
COPY --from=builder app/list-of-colors /app/
COPY --from=builder app/animals-list /app/
CMD ["/app/main", "-words", "/dict/words", "-colors", "list-of-colors", "-animals", "animals-list"]
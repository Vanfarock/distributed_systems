FROM golang:1.18-alpine
WORKDIR /app

COPY ./pkg/tweet_writer/go.mod .
COPY ./pkg/tweet_writer/go.sum .
RUN go mod download

COPY ./pkg/tweet_writer .
RUN go build -o /tweet_writer

EXPOSE 8080

CMD [ "/tweet_writer" ]
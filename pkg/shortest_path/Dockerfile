FROM golang:1.18-alpine
WORKDIR /app

COPY ./pkg/shortest_path/go.mod .
COPY ./pkg/shortest_path/go.sum .
RUN go mod download

COPY ./pkg/shortest_path .
RUN go build -o /shortest_path

EXPOSE 3333

CMD [ "/shortest_path" ]

FROM golang:alpine

RUN apk add libc-dev gcc

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /trouvee

EXPOSE 8080

ENV GIN_MODE=release

CMD [ "/trouvee" ]
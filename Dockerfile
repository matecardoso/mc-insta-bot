FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /mc-insta-bot ./cmd/mc-insta-bot

EXPOSE 3000

CMD [ "/mc-insta-bot" ]

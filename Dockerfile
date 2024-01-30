FROM golang:latest

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download && go mod verify

COPY . .

RUN make build

# Commented since railway bans VOLUMES, if deploying on another service uncomment this
# VOLUME /app/data

EXPOSE $PORT

CMD ["./build/server"]

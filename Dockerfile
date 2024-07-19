FROM golang:1.22.4

ARG APP_PORT

RUN mkdir /app

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

EXPOSE ${APP_PORT}

CMD ["go", "run", "cmd/main.go"]
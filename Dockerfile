FROM golang:1.22.4

RUN mkdir /app

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]
FROM golang

WORKDIR ./app

COPY . .

RUN go mod tidy

WORKDIR ./init

EXPOSE 8080

CMD ["go", "run", "main.go"]


FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o miniproject2

CMD ["./miniproject"]
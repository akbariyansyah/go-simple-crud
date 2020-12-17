FROM golang:alipine

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 9090

CMD /app/main
FROM golang:1.20

LABEL org.opencontainers.image.authors="soyacen@outlook.com"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o mix-server ./main.go

EXPOSE 60062

ENTRYPOINT [ "/app/mix-server" ]

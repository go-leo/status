FROM golang:1.20

LABEL org.opencontainers.image.authors="soyacen@outlook.com"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o http-server ./main.go

EXPOSE 60061

ENTRYPOINT [ "/app/http-server" ]

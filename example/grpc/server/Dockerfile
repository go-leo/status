FROM golang:1.20

LABEL org.opencontainers.image.authors="soyacen@outlook.com"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o grpc-server ./main.go

EXPOSE 50051

ENTRYPOINT [ "/app/grpc-server" ]

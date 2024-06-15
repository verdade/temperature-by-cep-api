FROM golang:latest

WORKDIR /app
EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]


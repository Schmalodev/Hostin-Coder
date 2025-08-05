FROM golang:1.24.4

WORKDIR /HostingCoder

COPY . .

EXPOSE 8080

CMD ["go", "run", "run.go"]

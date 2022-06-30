FROM golang:1.16
WORKDIR /url-shortener 
COPY ./ /url-shortener
CMD go run cmd/main.go -docker
FROM golang:1.17
WORKDIR /url-shortener 
COPY ./ /url-shortener
CMD go run cmd/main.go -docker
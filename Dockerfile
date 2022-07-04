FROM golang:1.17
WORKDIR /url-shortener 
COPY ./ /url-shortener
CMD ./cmd/url-shortener -docker
FROM golang:1.23.2-alpine

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

# RUN go install github.com/golang/mock/mockgen@v1.5.0

RUN go install github.com/air-verse/air@latest

# RUN apt-get update && apt-get install sqlite3 -y

# RUN usermod -u 1000 www-data
# RUN mkdir -p /var/www/.cache
# RUN chown -R www-data:www-data /go
# RUN chown -R www-data:www-data /var/www/.cache
# USER www-data

# CMD ["tail", "-f", "/dev/null"]
# CMD ["go", "run", "main.go"]
CMD ["air"]
# Image repository
FROM golang:1.12.0-alpine3.9

# create directory
RUN mkdir /app

# copy every in current to app directory
ADD . /app

# all execution should be app directory
WORKDIR /app

# compile binary
RUN go build -o main .

# port running on
EXPOSE 8080

# start program
CMD ["/app/main"]

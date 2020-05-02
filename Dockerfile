# Image repository
FROM golang:1.12.0-alpine3.9

# create directory
RUN mkdir /food-hub-api

# copy every in current to food-hub-api directory
ADD . /food-hub-api

# all execution should be food-hub-api directory
WORKDIR /food-hub-api

# compile binary
RUN go build -o main .

# port running on
EXPOSE 8080

# start program
CMD ["/food-hub-api/main"]

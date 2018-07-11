# Use an official Golang runtime as a parent image
FROM golang:alpine
ENV APP_NAME=stock-analytics
ENV WORKING_DIR=$GOPATH/src/github.com/marvin5064/$APP_NAME
# Set the working directory to golang working space
WORKDIR $WORKING_DIR

# Copy the current directory contents into the container at /app
ADD . $WORKING_DIR

RUN go build

# Run by makefile when the container launches
CMD ["sh", "-c", "./$APP_NAME"]
# CMD ["which", "go"]

# download base/parent image.
FROM golang:1.21

# specify working directory for running in container
WORKDIR /app

# copy go module dependency file first
COPY go.mod .
COPY go.sum .

# download all the dependencies
RUN go mod tidy

# copy all the files 
COPY . .

# build the application
RUN go build -o coffeeShop .

# expose port
EXPOSE 8080

# execute the command which gets the application running
CMD ["./coffeeShop"]

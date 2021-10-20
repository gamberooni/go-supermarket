FROM golang:1.17-alpine

WORKDIR /app

# download go modules
COPY go.mod .
COPY go.sum /
RUN go mod download

# copy everything in the current Docker context into the /app directory in the container
COPY . .

# build go bin for the application
RUN go build -o ./bin/go-supermarket

EXPOSE 1323

CMD [ "/app/bin/go-supermarket" ]

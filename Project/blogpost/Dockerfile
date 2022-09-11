#docker file is used to create an image so that we can make a container

#take the base image golang:1.18.6-alpine3.16
FROM golang:1.18.6-alpine3.16

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /app
WORKDIR /app

#copy all files from local machine to the container
COPY . .

# Copy and download dependency using go mod
RUN go mod download

# You may want to add “go mod verify” after download, 
# so you can be sure your go.sum file matches what you 
# downloaded and someone didn’t “git push — force” 
# over a tag you’re using into something hacked.
RUN go mod verify

# Build the application.This will generate a blogpost.exe file
RUN go build

# Command to run when starting the container of the newly formed image.
CMD ["./blogpost"]
#This will start the blogpost application by executing the exe file

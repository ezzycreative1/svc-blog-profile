# BUILD STAGE ############################################################
FROM golang:1.18.3-alpine3.16 as builder

# maintainer Info
LABEL maintainer="Ezzycreative <adi.setaydharma83@gmail.com>"

# Update alpine repository indexs and next install git and openssh-client
RUN apk update && apk add --no-cache git openssh-client

# go get uses git internally. The following one liners will make git and consequently go get clone your package via SSH.
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

# set go private
RUN export GOPRIVATE="github.com/ezzycreative1/*"

# add credentials on build
ARG SSH_PRIVATE_KEY
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa
RUN echo "StrictHostKeyChecking no " > /root/.ssh/config
RUN chmod 400 /root/.ssh/id_rsa

# make sure your domain is accepted
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan bitbucket.org >> /root/.ssh/known_hosts

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy dependencies list and download all.
COPY go.mod go.sum ./
RUN go mod download -x

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./app/cogs

# RUN STAGE ############################################################
FROM alpine:3.16

# Update alpine repository indexs
RUN apk update

# Setting certificate for Space S3
RUN apk add --no-cache ca-certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Install tzdata for set timezone
RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta

# Set Working Directory
WORKDIR /app

# Copy main app from builder stage
COPY --from=builder /app/main .

EXPOSE 8081 4000

# Command to run the executable
CMD ["./main"]
FROM golang:latest 

LABEL maintainer="Abhijit <gta4roy@gmail.com>"

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download 
COPY . .

ENV MYSQL_HOST=192.168.0.110
ENV MYSQL_USER="gta4roy"
ENV MYSQL_PASSWORD="71201"
ENV DB_NAME="address"

RUN go build -o address_client
RUN cd service 
RUN go build -o address_service 
RUN cd ../


CMD [ "chmod 755 ./startService.sh","./startService.sh" ]






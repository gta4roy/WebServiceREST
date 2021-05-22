FROM golang:latest 

LABEL maintainer="Abhijit <gta4roy@gmail.com>"

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download 
COPY . .

ENV MYSQL_USER="root"
ENV MYSQL_PASSWORD="71201"
ENV DB_NAME="address"
ENV HOST="192.168.0.110"
ENV MYSQL_PORT="3306"

EXPOSE 9098

RUN go build -o address_client .
RUN cd /app/service 
RUN go build -o address_service .
RUN cd /app


CMD [ "chmod 755 ./startService.sh","./startService.sh" ]






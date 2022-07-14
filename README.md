# fastpv-auth

### set package manager mod
    
    export GO111MODULE=on
 

### install packages dependen
    go get -u github.com/gorilla/mux
    go get github.com/go-sql-driver/mysql
    go get github.com/gorilla/handlers

## run redis and mysql
    docker volume create mysqlfastpv
    cd docker
    docker-compose up -d

## files to create database and load data
    cd database

## run local
    
    go run server.go
    


## Run  microservice on docker

    cd docker/app
    docker build -t edumar111/fastpv-auth:1.0.0 -f Dockerfile .
    docker run  -it  --name fastpv-auth -p 8080:8080 --network docker_fastpv  -e HOST_MYSQL=mysql-service-svc -e  HOST_REDIS=redis-service-svc -e GO_ENV=production edumar111/fastpv-auth:1.0.0


## API login 

    curl   POST 'localhost:8080/token-auth' \
    --header 'Content-Type: application/json' \
    --data-raw '{
	"username":"admin",
	"password":"123456"
    }'

# fastpv-auth

### set package manager mod
    
    export GO111MODULE=on
    go mod init
## run redis and mysql
    docker volume create mysqlfastpv
    cd docker
    docker-compose up -d

    
## run
    
    go run server.go
    
### install packages dependen
    
    go get -u github.com/gorilla/mux
    go get github.com/go-sql-driver/mysql
    go get github.com/gorilla/handlers
    
    
## login 

    curl   POST 'localhost:8080/token-auth' \
    --header 'Content-Type: application/json' \
    --data-raw '{
	"username":"admin",
	"password":"123456"
    }'
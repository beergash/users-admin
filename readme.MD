## Prerequirements:

## Database:

1) Create a postgres Database instance
2) Set database properties (user, password, host, db name) in the file $GOPATH/src/users-admin/config/config.go
3) Set database properties for the db test  in the file (view const variables) $GOPATH/src/users-admin/test/dao_test.go
4) Execute the sql script : $GOPATH/src/users-admin/sql/create_schema_ddl.sql to initialize schema

## Test:
To run unit test go to the folder $GOPATH/src/users-admin/test and execute command :
go test -v
To get test coverage execute command:
go test -cover -coverpkg=../...

## Application Startup:
To start application execute command:
go run $GOPATH/src/users-admin/main.go

## Build application and give an executable file name
GOOS=linux GOARCH=386 go build -o users-admin.sh cmd/users-admin/main.go

## Rest Calls curl examples:

### Get all users: 
curl localhost:8080/users

### Get user By Id: 
curl localhost:8080/users/1

### Save user: 
curl -X POST -H "Content-Type: application/json"  --data '{"username":"","password":"","name":"","surname":"","birth_date":"1990-03-29T00:00:00Z"}' localhost:8080/create-user

### Modify user: 
curl -X PUT  -H "Content-Type: application/json"  --data '{"username":"","password":"","name":"","surname":"","birth_date":"1990-03-29T00:00:00Z"}' localhost:8080/users/1

### Delete user:
curl -X DELETE localhost:8080/users/1

### Search users:
curl -X POST -H "Content-Type: application/json"  --data '{"name": "", "surname": "", "page": 1, "size": 10 }' localhost:8080/users/search


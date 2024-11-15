# Communication between two or more docker-compose projects

## main.go
Simple API to test database connect.

## postgres-compose.yml
Creates postgres container and commonNetwork (bridge network)

## golang-compose.yml
Creates golang app and join to commonNetwork (as external network)


## Steps

Init go project

```go mod init helloworld```

Run postgres compose

```docker-compose -f postgres-compose.yml up -d```

Last, run golang compose to join to network

```docker-compose -f golang-compose.yml up --build```
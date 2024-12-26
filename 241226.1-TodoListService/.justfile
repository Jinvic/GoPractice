set shell := ["powershell.exe", "-c"]

run:
    go run main.go

swagger:
    swag fmt
    swag init

sr:
    swag fmt
    swag init
    go run main.go

alias sw:=swagger
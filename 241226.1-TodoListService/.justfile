set shell := ["powershell.exe", "-c"]

run:
    go run main.go

swagger:
    swag init

sr:
    swag init
    go run main.go

alias sw:=swagger
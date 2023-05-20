proto project:
	protoc --go_out=. -I=proto \
				   --go_opt=paths=import \
				   --go-grpc_out=. -I=proto \
				   --go-grpc_opt=paths=import \
				   proto/project/*.proto
go linux:
	GOARCH=amd64 GOOS=linux go build cmd/main.go

go build:
	go build cmd/main.go

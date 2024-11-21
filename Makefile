migrateup:
	migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/simple_bank" -path db/migrations -verbose up

migrateup1:
	migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/simple_bank" -path db/migrations -verbose up 1

migratedown:
	migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/simple_bank" -path db/migrations -verbose down

migratedown1:
	migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/simple_bank" -path db/migrations -verbose down 1

server:
	go run main.go

proto:
	rm -f pn/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
  proto/*.proto



.PHONY: migrateup migratedown migrateup1 migratedown1 server proto
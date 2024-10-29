migrateup:
	migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/my_bank" -path db/migrations -verbose up

migrateup1:
	migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/my_bank" -path db/migrations -verbose up 1

migratedown:
	migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/my_bank" -path db/migrations -verbose down

migratedown1:
	migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/my_bank" -path db/migrations -verbose down 1

server:
	go run main.go

.PHONY: migrateup migratedown migrateup1 migratedown1 server 
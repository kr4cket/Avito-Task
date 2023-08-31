.SILENT:

build-containers:
	docker-compose up -d --build avito-app

add-tables: 
	migrate -database postgres://postgres:task@localhost:5436/?sslmode=disable -path ./schema up


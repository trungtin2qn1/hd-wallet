# Make file

build:
	go build -o hd-wallet main.go

run:
	go run main.go

# stop-services:
# 	docker-compose -f ./setup/docker-compose/postgreSQL.yml down
# 	docker-compose -f ./setup/docker-compose/app.yml down
# 	docker volume rm docker-compose_database-data

# local-db:
# 	docker-compose -f ./setup/postgreSQL.yml up -d
# 	@while ! docker exec hd-wallet_postgres pg_isready -h localhost -p 5432 > /dev/null; do \
# 		sleep 1; \
# 	done
# 	docker cp ./setup/data.sql hd-wallet_postgres:/
# 	docker exec -u postgres hd-wallet_postgres psql hd-wallet user -f /data.sql

# app:
# 	docker-compose -f ./setup/docker-compose/app.yml up -d

setup-package:
#	mkdir logs
	go get github.com/Masterminds/glide
	glide install

test:
	go test -v ./...

docker-image:
	make build 
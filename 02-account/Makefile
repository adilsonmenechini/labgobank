.PHONY: help start tidy db  clean 

LDEPLOY="deploy"
migrationDir=migration
include deploy/.env
export DB_HOST=$(shell docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' db_psql)

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
##
## ----------------
## Codigo
## ----------------
## make start - run app
start:
	@go run main.go

## make tidy - clean cache and update mod
tidy:
	@go clean --modcache
	@go mod tidy 
	@go mod vendor
##
## ----------------
## Docker Compose
## ----------------
## make db - Docker-compose PostgreSQL
	
db:  
	@docker compose -f ${LDEPLOY}/docker-compose.yaml up -d


## make clean - Docker-compose Cleaning
clean: 
	@docker compose  -f ${LDEPLOY}/docker-compose.yaml down -v
	@docker system prune -a --volumes

##
## ----------------
## Migrate
## ----------------
## make migrate - Migrate Create
migrate: 
	docker run --rm -v .\${migrationDir}:/migration migrate/migrate create -ext sql -dir migration/ -seq transaction

## make migrateUp - Migrate Up
migrateUp: 
	 docker run --rm  -v .\${migrationDir}:/migrations --network bridge migrate/migrate -path=/migrations/ -database "${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"  -verbose up

## make migrateDown - Migrate Down
migrateDown: 
	 docker run --rm  -v .\${migrationDir}:/migrations migrate/migrate -path=/migrations/ -database "${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"  -verbose down

## make migrateDrop - Migrate Drop
migrateDrop: 
	 docker run --rm  -v .\${migrationDir}:/migrations migrate/migrate -path=/migrations/ -database "${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose drop -f

##
## ----------------
## SQLC
## ----------------
## make sqlc - sqlc generate
sqlcg:
	sqlc generate
server:
	go run ./server start

client:
	go run client/*.go

pull_pg: 
	docker pull postgres

create_pg_volume:
	docker volume create postgres

# https://geshan.com.np/blog/2021/12/docker-postgres/
run_pg:
	docker run --name pg01 --rm -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 \
	 -e PGDATA=/var/lib/postgresql/data/pgdata \
	 -v postgres:/var/lib/postgresql/data -p 5432:5432 -it postgres -c log_statement=all

exec_pg:
	docker exec -it pg01 /bin/sh

stop_pg:
	docker container stop pg01

rm_pg:
	docker container rm -f pg01

add_migration_file:
	migrate create -dir=server/migrations -ext=sql $(name)

run_migration:
	go run ./server migration

seed:
	go run ./server --seednum=$(n) seed

.PHONY: server client pull_pg create_pg_volume run_pg start_pg stop_pg rm_pg add_migration_file run_migration seed

# 
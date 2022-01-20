dep:
	go mod download

gen:
	go get -u github.com/99designs/gqlgen/cmd
	go run github.com/99designs/gqlgen generate

run:
	docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=linksdb -d mysql:latest

rm:
	docker stop $$(docker container ls --filter=name=mysql -q)
	docker rm $$(docker container ls --filter=status=exited -q)

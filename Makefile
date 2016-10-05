TEAMID ?= 7357
all:;: '$(T)'
repo=tompscanlan/labreserved
bin=labreserved-server

all: $(bin)-local

$(bin): deps
	CGO_ENABLED=0 go build -a -v --installsuffix cgo  ./cmd/$(bin)
	file $(bin)
	ldd $(bin)

$(bin)-local: deps
	go build -v -o $(bin)-local  ./cmd/$(bin)
deps:
	go get -v ./...

docker: $(bin) temp.crt temp.key
	docker build -t $(repo) --rm=true .

temp.crt temp.key:
	./scripts/make-cert.sh

dockerclean:
	echo "Cleaning up Docker Engine before building."
	docker rm $$(docker ps -a | awk '/$(bin)/ { print $$1}') || echo -
	docker rmi $(bin)

clean: stop dockerclean
	go clean
	rm -f $(bin)

run:
	docker run -d -p2080:80 -p20443:443 -e TEAM_ID=$(TEAMID)  --name labd $(repo)

stop:
	docker kill $$(docker ps -a | awk '/$(bin)/ { print $$1}') || echo -

valid:
	go tool vet .
	go test -v -race ./...

.PHONY: imports docker clean run stop deps


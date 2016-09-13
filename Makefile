TEAMID ?= 7357
all:;: '$(T)'
repo=tompscanlan/labreserved
bin=labreserved-server

all: docker
local: $(bin)-local

$(bin):
	docker run -it -v "$$(PWD)":/go/src/github.com/$(repo) -w /go/src/github.com/$(repo) golang:1.6 bash -c "CGO_ENABLED=0 go get -v ./... && go build -a -v --installsuffix cgo  ./cmd/$(bin)"
#	go get -v ./...
#	go build -a -v ./cmd/$(bin)

$(bin)-local: deps
	go build -v -o $(bin)-local  ./cmd/$(bin)
deps:
	go get -v ./...

docker: $(bin) temp.crt temp.key
	./scripts/make-cert.sh
	docker build -t $(repo) --rm=true .

dockerclean:
	echo "Cleaning up Docker Engine before building."
	docker rm $$(docker ps -a | awk '/$(bin)/ { print $$1}') || echo -
	docker rmi $(bin)

clean:
	go clean
	rm -f $(bin)
	echo "Cleaning up Docker Engine before building."
	docker kill $$(docker ps -a | awk '/$(bin)/ { print $$1}') || echo -
	docker rm $$(docker ps -a | awk '/$(bin)/ { print $$1}') || echo -
	docker rmi $(bin)

run:
	docker run -d -p2080:80 -p20443:443 -e BLOB_ID=$(TEAMID)  $(repo)

stop:
	docker kill $$(docker ps -a | awk '/$(bin)/ { print $$1}') || echo -

.PHONY: imports docker clean run stop


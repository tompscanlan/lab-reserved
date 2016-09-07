
bin=labreserved-server

all: docker

$(bin):
	docker run -it -v "$$(PWD)":/go/src/github.com/tompscanlan/$(bin) -w /go/src/github.com/tompscanlan/$(bin) golang:1.6 bash -c "CGO_ENABLED=0 go get -v ./... && go build -a -v --installsuffix cgo  ./cmd/$(bin)"
#	go get -v ./...
#	go build -a -v ./cmd/$(bin)

docker: $(bin) temp.crt temp.key
	./scripts/make-cert.sh
	docker build -t $(bin) --rm=true .

clean:
	go clean
	rm -f $(bin)
	echo "Cleaning up Docker Engine before building."
	docker kill $$(docker ps -a | awk '/$(bin)/ { print $$1}') || echo -
	docker rm $$(docker ps -a | awk '/$(bin)/ { print $$1}') || echo -
	docker rmi $(bin)

run:
	docker run -d -p2080:80 -p20443:443  $(bin)

stop:
	docker kill $$(docker ps -a | awk '/$(bin)/ { print $$1}') || echo -

.PHONY: imports docker clean run stop


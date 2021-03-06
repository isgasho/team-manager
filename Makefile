all: local

docker-image:
	docker build -t cilium/team-manager:${VERSION} .

tests:
	go test -mod=vendor ./...

team-manager: tests
	CGO_ENABLED=0 go build -mod=vendor -a -installsuffix cgo -o $@ ./cmd/main.go

local: team-manager
	strip team-manager

clean:
	rm -fr team-manager

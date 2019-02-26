ci_deps:
	go get github.com/stretchr/testify/assert
	go get -u github.com/mitchellh/gox
	go get github.com/jteeuwen/go-bindata/...
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	go get github.com/mattn/goveralls

test:
	go test -v ./... -covermode=count -coverprofile=profile.cov -bench=.

goveralls:
	goveralls -coverprofile=profile.cov -service=travis-ci

fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

lint:
	gometalinter --vendor --disable-all \
		--enable=deadcode \
		--enable=ineffassign \
		--enable=gosimple \
		--enable=gofmt \
		--enable=goimports \
		--enable=misspell \
		--enable=errcheck \
		--enable=vet \
		--enable=vetshadow \
		--deadline=10m \
		./...

build_rudy:
	docker-compose -f build/rudy-docker-compose.yml build

build_salz:
	docker-compose -f build/salz-docker-compose.yml build

down:
	docker-compose -f build/rudy-docker-compose.yml down
	docker-compose -f build/salz-docker-compose.yml down

rudy: down build_salz
	docker-compose -f build/rudy-docker-compose.yml up --scale rudyc=3

salz: down build_salz
	docker-compose -f build/salz-docker-compose.yml up --scale salzc=3

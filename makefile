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
		--enable=staticcheck \
		--enable=gofmt \
		--enable=goimports \
		--enable=misspell \
		--enable=errcheck \
		--enable=vet \
		--enable=vetshadow \
		--deadline=10m \
		./...

build_rudy:
	sudo docker-compose -f build/rudy-docker-compose.yml build

rudy: build_rudy
	sudo docker-compose -f build/rudy-docker-compose.yml up --scale rudyc=5

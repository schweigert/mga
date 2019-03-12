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

tag_images:
	docker tag build_rudyc schweigert/tcc:rudyc_v001
	docker tag build_rudyweb schweigert/tcc:rudyweb_v001
	docker tag build_rudydb schweigert/tcc:rudydb_v001
	docker tag build_rudygh schweigert/tcc:rudygh_v001
	docker tag build_rudya schweigert/tcc:rudya_v001

docker_push: tag_images
	docker push schweigert/tcc:rudyc_v001
	docker push schweigert/tcc:rudyc_v001
	docker push schweigert/tcc:rudyweb_v001
	docker push schweigert/tcc:rudydb_v001
	docker push schweigert/tcc:rudygh_v001
	docker push schweigert/tcc:rudya_v001
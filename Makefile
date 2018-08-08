TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: test

fmt:
	gofmt -w $(GOFMT_FILES)

test: fmt
	docker-compose down
	docker-compose up -d --build --force-recreate
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test -v
	docker-compose down
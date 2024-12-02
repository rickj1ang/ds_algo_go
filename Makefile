.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test -v -race

.PHONY: push
push:
	git push -u origin main

.PHONY: app
app: bin/epitome
	mkdir -p bin/.epitome
	cp -r migrations bin/

bin/epitome: $(shell find . -path '**/*.go')
	go build -o bin/epitome github.com/arturovm/epitome/cmd/epitome

.PHONY: test
test:
	go test ./...

.PHONY: run
run: bin/epitome
	./bin/epitome --debug

.PHONY: clean
clean:
	rm -rf bin

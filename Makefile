
all: restore-deps test vet

test:
	go test -v ./pkg/*
	
vet: 
	go vet ./pkg/*

clean:
	rm *.test

restore-deps:
	go mod tidy

build: # explicit defaults bc they might change
	CGO_ENABLED=1 GOOS=linux go build -o bin/shakesearch cmd/*.go
build:
	@/opt/go/bin/go build -o bin/apistore
run:  build
	./bin/apistore
test: 
	@/opt/go/bin/go test -v ./...
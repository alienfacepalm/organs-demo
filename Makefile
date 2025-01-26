# Build targets
.PHONY: build clean test run

# Ensure bin directory exists
$(shell mkdir -p bin)

build:
	cd go && go build -o ../bin/organs-demo

clean:
	rm -rf bin/

test:
	cd go && go test ./...

run: build
	./bin/organs-demo 
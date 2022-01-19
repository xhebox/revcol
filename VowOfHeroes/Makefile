BIN_LIST = $(patsubst bin/%,dist/%,$(wildcard bin/*))
PROTO_LIST = $(patsubst %.proto,%.pb.go,$(wildcard proto/*.proto))
PACKAGE = github.com/xhebox/revcol/VowOfHeroes
OS = linux

.PHONY: clean

all: $(BIN_LIST)

clean:
	rm -rf proto/*.pb.go dist/

%.pb.go: %.proto
	protoc -I=proto/ --go_out=$(shell pwd) $(patsubst %.pb.go,%.proto,$@)

dist/%: $(PROTO_LIST) bin/%
	env GOOS=$(OS) go build -o $(patsubst dist/%,dist/$(OS)/%,$@) $(PACKAGE)/$(patsubst dist/%,bin/%,$@)

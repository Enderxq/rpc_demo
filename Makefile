GOCMD=go
GOBUILD=$(GOCMD) build

build: proto
	$(GOBUILD) -o bin/cli -v cli/main.go
	$(GOBUILD) -o bin/srv -v srv/main.go

proto:
	export PATH=$(PATH):/Users/enderx/go/bin && \
	protoc -I ./myProto --go_out=plugins=grpc:./myProto ./myProto/message.proto


clean:
	rm -f bin/cli
	rm -f bin/srv

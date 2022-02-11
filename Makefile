PROTOC_ZIP := protoc-3.15.8-linux-x86_64.zip

install-grpc:
	apt-get install -y sudo
	sudo apt-get install zip unzip
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/$(PROTOC_ZIP)
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_ZIP)

generate:
	protoc --go_out=. --go_opt=paths=source_relative ./1.18/11-workspace/def/*.proto

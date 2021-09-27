
run:
	@go build -o a.out && ./a.out
	-@rm ./a.out

gen:
	protoc \
		-I=. -I=$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		service/helloworld.proto
	protoc \
		-I=. -I=$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
		service/helloworld.proto
	protoc \
		-I=. -I=$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--swagger_out=. \
		service/helloworld.proto

clean:
	-rm service/*.pb.go service/*.pb.*.go
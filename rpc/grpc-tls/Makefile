
run:
	@go build -o a.out && ./a.out
	-@rm ./a.out

gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative service/helloworld.proto

clean:
	-rm *.pb.go
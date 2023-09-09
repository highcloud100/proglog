compile:
	protoc 	api/v1/*.proto \
		--proto_path=. \
		--go_out=. \
		--go_opt=paths=source_relative 
		
test:
	go test -race ./...
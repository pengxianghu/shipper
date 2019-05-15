# vessel-service/Makefile
build:  
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/pengxianghu/shipper/vessel-service proto/vessel/vessel.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go
	# docker build -t vessel-service .

run:  
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service

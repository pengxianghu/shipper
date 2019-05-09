build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

run:  
	docker run -e MICRO_REGISTRY=mdns user-cli

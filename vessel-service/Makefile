build:
	protoc --micro_out=. --go_out=. ./proto/vessel/vessel.proto
	docker build -t vessel-service .

run:
	docker run -d --net="host" \
		-p 50053 \
		-e MICRO_SERVER_ADDRESS=:50053 \
		-e MICRO_REGISTRY=mdns \
		vessel-service
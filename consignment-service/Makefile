build:
	protoc --micro_out=. --go_out=. ./proto/consignment/consignment.proto
	docker build -t tsicg/consignment-service .
	docker push tsicg/consignment-service:latest

run:
	docker run -d --net="host" \
		-p 50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns \
		-e DISABLE_AUTH=true \
		tsicg/consignment-service:latest

deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml
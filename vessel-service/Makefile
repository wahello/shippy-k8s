build:
	protoc --micro_out=. --go_out=. ./proto/vessel/vessel.proto
	docker build -t tsicg/vessel-service .
	docker push tsicg/vessel-service:latest

run:
	docker run -d --net="host" \
		-p 50053 \
		-e MICRO_SERVER_ADDRESS=:50053 \
		-e MICRO_REGISTRY=mdns \
		tsicg/vessel-service:latest

deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml
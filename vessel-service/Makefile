namespace='default'

build:
	protoc --micro_out=. --go_out=. ./proto/vessel/vessel.proto
	docker build -t gcr.io/cgault-sandbox/vessel-service:latest .
	docker push gcr.io/cgault-sandbox/vessel-service:latest

# run:
# 	docker run -d --net="host" \
# 		-p 50053 \
# 		-e MICRO_SERVER_ADDRESS=:50053 \
# 		-e MICRO_REGISTRY=mdns \
# 		tsicg/vessel-service:latest

deploy:
	sed "s/NAMESPACE_ID/${namespace}/g; s/{{ UPDATED_AT }}/$(shell date)/g;" ./deployments/deployment.yaml > ./deployments/tmp/deployment.yaml
	kubectl apply -f ./deployments/tmp/deployment.yaml
	# sed "s/{{ NAMESPACE_ID }}/$(NAMESPACE_ID)/g;" ./deployments/service.yaml > ./deployments/tmp/service.yaml
	kubectl apply -f ./deployments/service.yaml
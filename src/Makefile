build_image:
	docker build . -t user-service:v1.0.0

run_with_dapr:
	dapr run --app-id userservice --app-port 8085 --dapr-http-port 3601 --dapr-grpc-port 60001 --resources-path ../.dapr/components go run ./cmd/main.go

yml_create:
	az containerapp create \
		-g microservice-reference-framework \
		-n test-service-1 \
		--yaml ../.azure/app_config.yml

enable_dapr:
	az containerapp dapr enable \
		-n test-service-1 \
		-g microservice-reference-framework \
		--dapr-app-id userservice \
		--dapr-app-protocol http \
		--dapr-app-port 8085

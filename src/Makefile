build_image:
	docker build . -t user-service:v1.0.0

run_with_dapr:
	dapr run --app-id userservice --app-port 8085 --dapr-http-port 3601 --dapr-grpc-port 60001 --resources-path ../.dapr/components go run ./cmd/main.go

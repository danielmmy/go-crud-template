up:
	@echo "go_crud_template stopping docker image if running..."
	docker-compose down
	@echo "bulding and starting go_crud_template docker img..."
	DOCKER_SCAN_SUGGEST=false docker-compose up --build -d
	@echo "go_crud_template docker img built and started."

down:
	@echo "stopping go_crud_template docker img..."
	docker-compose down
	@echo "go_crud_template docker img down"

start:
	@echo "compiling and running go_crud_template locally. CTRL+c to end"
	go run ./cmd/api
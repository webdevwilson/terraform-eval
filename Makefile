
.PHONY: clean
clean:
	rm -rf bin

bin/terraform-eval:
	go build -o bin/terraform-eval src/main.go

.PHONY: docker-build
docker-build:
	docker build -t terraform-eval .


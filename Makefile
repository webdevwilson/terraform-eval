IMG := webdevwilson/terraform-eval:latest

.PHONY: clean
clean:
	rm -rf bin

bin/terraform-eval:
	go build -o bin/terraform-eval src/main.go

.PHONY: docker-build
docker-build:
	docker build -t $(IMG) .

.PHONY: docker-push
docker-push: docker-build
	docker push $(IMG)
prune:
	@up prune
deploy:
	@up
run:
	@go run main.go

test-local:
	@./test_local.sh

test-remote:
	@./test.sh

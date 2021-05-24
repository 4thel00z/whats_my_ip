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

test-remote-prod:
	@./test-prod.sh

prod-url:
	@up url -s production

staging-url:
	@up url

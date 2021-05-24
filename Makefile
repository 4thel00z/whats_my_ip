prune:
	@up prune
deploy:
	@up
run:
	@./run.sh

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

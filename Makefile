test-watch:
	fd | entr -r gotest ./...

run-test:
	go test -race ./test/... -coverpkg=./... -coverprofile ./coverage.out
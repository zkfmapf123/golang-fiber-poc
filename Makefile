dev-run:
	export ENV=dev && go run app.go

debug-run:
	export ENV=debug && go run app.go

test:
	export ENV=test && go test ./src/*
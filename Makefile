dev-run:
	export ENV=dev && go run app.go

debug-run:
	export ENV=debug && go run app.go
dev-run:
	export ENV=dev && go run app.go

debug-run:
	export ENV=debug && go run app.go

docker-dev-run:
	@docker build -t apiserver -f Dockerfile . && docker run --rm -it \
	-e ENV=dev \
	-e IS_LOCAL=false \
	-e IS_PORK=true \
	-e VERSION=1.0.0 \
	-e PORT=3000 \
	-p 3000:3000 apiserver

test:
	export ENV=test && go test ./src/*
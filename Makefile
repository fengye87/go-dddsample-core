.DEFAULT_GOAL := test

generate:
	docker build -f hack/Dockerfile . | tee /dev/tty | tail -n1 | cut -d' ' -f3 | xargs -I{} \
		docker run --rm -v $$PWD:/workspace -w /workspace {} go generate ./...

fmt:
	go fmt ./...

test:
	go test ./... -coverprofile cover.out

run:
	skaffold run --tail

manifest:
	skaffold render --default-repo=fengye87 --offline=true --digest-source=tag > manifest.yaml

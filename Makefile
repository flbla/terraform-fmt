IMAGE := terraform-fmt/terraform-fmt
SHELL := /bin/bash
export PATH := /home/ubuntu/.local/bin:/home/ubuntu/go/bin:$(PATH)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: init
init:
	git config core.hooksPath .githooks

.PHONY: image
image:
	docker build  -t $(IMAGE) .

.PHONY: test
test:
	which gotestsum || (pushd /tmp && go install gotest.tools/gotestsum@latest && popd)
	gotestsum -- --mod=readonly -bench=^$$ -race ./...

.PHONY: tagger
tagger:
	@git checkout master
	@git fetch --tags
	@echo "the most recent tag was `git describe --tags --abbrev=0`"
	@echo ""
	read -p "Tag number: " TAG; \
	 git tag -a "$${TAG}" -m "$${TAG}"; \
	 git push origin "$${TAG}"

.PHONY: cyclo
cyclo:
	which gocyclo || go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	gocyclo -over 15 -ignore 'vendor/|funcs|cmd/tfsec-skeleton' .

.PHONY: vet
vet:
	go vet ./...

.PHONY: typos
typos:
	which codespell || pip install codespell
	codespell -S vendor,funcs,.terraform,.git --ignore-words .codespellignore -f

.PHONY: quality
quality: cyclo vet

.PHONY: fix-typos
fix-typos:
	which codespell || pip install codespell
	codespell -S vendor,funcs,.terraform --ignore-words .codespellignore -f -w -i1

.PHONY: pre-pr
pre-pr: quality typos test

.PHONY: bench
bench:
	go test -run ^$$ -bench . ./...
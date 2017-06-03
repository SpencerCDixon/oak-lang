NAME := oak-lang
VERSION := v0.0.1
GIT_REPO := github.com/spencercdixon/oak

# Builds language injecting proper version num in.
build:
	go build -ldflags "-X $(GIT_REPO)/cmd.version=$(VERSION)"

# Compiles go using Mitch's cross compile tool.
compile: 
	@rm -rf build/
	@gox -ldflags "-X $(GIT_REPO)/cmd.version=$(VERSION)" \
		-osarch="darwin/amd64" \
		-os="linux" \
		-output "build/{{.Dir}}_$(VERSION)_{{.OS}}_{{.Arch}}/$(NAME)" \
		./...

# Installs required dependencies.
deps:
	# Dev Deps
	go get github.com/c4milo/github-release
	go get github.com/mitchellh/gox
	go get github.com/tj/mmake       
	# Prod Deps (none atm)

# Compile and deploy latest doc site
docs:
	@echo "TODO!  Once basic lang API is complete"

# Compiles source for supported OS's and then tar/sha's them.
dist: compile
	$(eval FILES := $(shell ls build))
	@rm -rf dist && mkdir dist
	@for f in $(FILES); do \
		(cd $(shell pwd)/build/$$f && tar -cvzf ../../dist/$$f.tar.gz *); \
		(cd $(shell pwd)/dist && shasum -a 512 $$f.tar.gz > $$f.sha512); \
		echo $$f; \
		done

# Releases Oak to github.  
#
# Note: make sure to add required GITHUB_TOKEN env variable.
release: dist
	@latest_tag=$$(git describe --tags `git rev-list --tags --max-count=1`); \
		comparison="$$latest_tag..HEAD"; \
		if [ -z "$$latest_tag" ]; then comparison=""; fi; \
		changelog=$$(git log $$comparison --oneline --no-merges); \
		github-release spencercdixon/$(NAME) $(VERSION) "$$(git rev-parse --abbrev-ref HEAD)" "**Changelog**<br/>$$changelog" 'dist/*'; \
		git pull

# Removes the dist/build folders that get generated every release.
clean:
	@rm -rf dist build

# Display list of TODO's inside source.
todo:
	@grep \
		--text \
		--color \
		--exclude Makefile \
		-nRo ' TODO:.*' .

test:
	@go test ./...

.PHONY: build compile deps dist release clean todo test docs

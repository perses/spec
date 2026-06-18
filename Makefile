GO                    ?= go
CUE                   ?= cue
GOCI                  ?= golangci-lint
GOFMT                 ?= $(GO)fmt

.PHONY: checkformat
checkformat:
	@echo ">> Check Go code format"
	! $(GOFMT) -d $$(find . -name '*.go' -not -path "./ui/*" -print) | grep '^'

.PHONY: checkunused
checkunused:
	@echo ">> Check for unused/missing packages in go.mod"
	$(GO) mod tidy
	@git diff --exit-code -- go.sum go.mod

.PHONY: checkstyle
checkstyle:
	@echo ">> Check Go code style"
	$(GOCI) run --timeout 5m

.PHONY: go-test
go-test:
	@echo ">> Run all go tests"
	$(GO) test -count=1 -v ./...

.PHONY: cue-eval
cue-eval:
	@echo ">> Validate CUE schemas"
	cd cue && $(CUE) eval ./...

.PHONY: cue-gen
cue-gen:
	@echo ">> Generate CUE definitions from golang datamodel"
	@for pkg in $$($(GO) list github.com/perses/spec/go/...); do \
		$(CUE) get go $$pkg; \
	done
	cp -r cue.mod/gen/github.com/perses/spec/go/* cue/ && rm -r cue.mod/gen
	find cue/ -name "*.cue" -exec sed -i 's/\"github.com\/perses\/spec\/go/\"github.com\/perses\/spec\/cue/g' {} \;

.PHONY: cue-test
cue-test:
	@echo ">> Run the unit tests for CUE schemas"
	$(GO) run ./scripts/test-cue/test-cue.go

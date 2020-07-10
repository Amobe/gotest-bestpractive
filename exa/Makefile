
TOOL := run_gotest

define run_gotest
	go test -tags=$(1) -v -cover -count=1 -parallel=1 ./...
endef

define run_gotestsum
	gotestsum -f=pkgname-and-test-fails -- -tags=$(1) -v -cover -count=1 -parallel=1 ./...
endef

test:
	$(call $(TOOL),"unit integration")

test.unit:
	$(call $(TOOL),"unit")

test.integration:
	$(call $(TOOL),"integration")

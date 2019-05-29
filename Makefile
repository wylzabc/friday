.PHONY: friday test

GOBIN = $(shell pwd)/build/bin

friday:
	$(shell pwd)/build/env.sh go build -o $(GOBIN)/friday
	@echo "Done building."
	@echo "Run \"$(GOBIN)/friday\" to launch."

test:
	@echo "Done test."


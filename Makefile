.PHONY: friday test

GOBIN = $(shell pwd)/build/bin

friday:
	$(shell pwd)/build/env.sh go build -o $(GOBIN)/friday
	@echo "Done building."
	@echo "Run \"$(GOBIN)/friday\" to launch."

test: addtest
	@echo "Done test."

addtest:
	$(shell pwd)/build/env.sh go test github.com/wylzabc/friday/add -v -cover
	@echo "Done test add."


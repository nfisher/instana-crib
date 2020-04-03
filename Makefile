
.PHONY: all
all: infraq

infraq: cmd/infraq/*.go
	go build -v ./cmd/infraq

.PHONY: clean
clean:
	go clean
	rm infraq
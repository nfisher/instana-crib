
.PHONY: all
all: infraq

infraq:
	go build -v ./cmd/infraq

.PHONY: clean
clean:
	go clean
	rm infraq
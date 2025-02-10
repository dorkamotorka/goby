.PHONY: init generate build run exec

# Only build and run
run: generate build exec

# Initialize go eBPF project
init:
	@go mod init main
	@go mod tidy
	@go install github.com/cilium/ebpf/cmd/bpf2go@latest

# Generate eBPF skeleton files
generate:
	@go generate

# Build the eBPF object file
build:
	@go build

# Run the program
exec:
	@sudo ./main


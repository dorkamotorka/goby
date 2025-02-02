.PHONY: generate build run

# Only build and run
build-run: generate build run

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
run:
	@sudo ./main

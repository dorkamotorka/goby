package generator

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func GenerateGoMain(path string) {
	outputFile := filepath.Join(path, "main.go")

	// Define the file content
	content := `package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target bpf program program.bpf.c

import (
	"os"
	"os/signal"
	"syscall"
	"log"

	"github.com/cilium/ebpf/rlimit"
)

func main() {
	// Remove resource limits for kernels <5.11.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock:", err)
	}

	// Load the compiled eBPF ELF and load it into the kernel.
	//var objs programObjects
	//if err := loadProgramObjects(&objs, nil); err != nil {
	//	log.Fatal("Loading eBPF objects:", err)
	//}
	//defer objs.Close()

	// Here you can now attach eBPF Programs
	// Check Cilium docs for more functionalities: https://pkg.go.dev/github.com/cilium/ebpf
	
	log.Println("eBPF program attached. Press Ctrl+C to exit.")

	// Set up signal handling to cleanly exit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Keep the program running until a signal is received
	<-stop

	log.Println("Received signal, exiting...")
}`

	// Create a new file or overwrite an existing one
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File 'main.go' created successfully!")
}

func GenerateeBPFProgram(path string) {
	outputFile := filepath.Join(path, "program.bpf.c")

	// Define the C code content
	content := `//go:build ignore

// Here you write an eBPF program

char _license[] SEC("license") = "GPL";`

	// Create a new eBPF file
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File 'program.bpf.c' created successfully!")
}

// DumpBTF runs the bpftool command to dump BTF information into vmlinux.h
func DumpBTF(path string) error {
	outputFile := filepath.Join(path, "vmlinux.h")

	// Create vmlinux.h for writing
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	cmd := exec.Command("bpftool", "btf", "dump", "file", "/sys/kernel/btf/vmlinux", "format", "c")
	cmd.Stdout = file 
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err == nil {
		log.Println("File 'vmlinux.h' created successfully!")
	}

	// Run the command and return any error
	return err 
}

// DumpMake creates a Makefile for convenience
func DumpMake(path string) error {
	outputFile := filepath.Join(path, "Makefile")

	// Create vmlinux.h for writing
	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	content := `.PHONY: init generate build run exec

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
`

	// Write the content to the file
        _, err = file.WriteString(content)
        if err != nil {
                log.Fatal(err)
        }

	log.Println("File 'Makefile' created successfully!")

	// Run the command and return any error
	return err 
}

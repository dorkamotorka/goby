package generator

import (
	"log"
	"os"
	"os/exec"
)

func GenerateGoMain() {
	// Define the file content
	content := `package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target bpf program program.c

import (
	"os"
	"os/signal"
	"syscall"
	"log"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

func main() {
	// Remove resource limits for kernels <5.11.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock:", err)
	}

	// Load the compiled eBPF ELF and load it into the kernel.
	var objs helloObjects
	if err := loadHelloObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	// Attach Tracepoint
	tp, err := link.Tracepoint("syscalls", "sys_enter_execve", objs.HandleExecveTp, nil)
	if err != nil {
		log.Fatalf("Attaching Tracepoint: %s", err)
	}
	defer tp.Close()
	
	log.Println("eBPF program attached to tracepoint. Press Ctrl+C to exit.")

	// Set up signal handling to cleanly exit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Keep the program running until a signal is received
	<-stop

	log.Println("Received signal, exiting...")
}`

	// Create a new file or overwrite an existing one
	file, err := os.Create("main.go")
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

func GenerateeBPFProgram() {
	// Define the C code content
	content := `//go:build ignore
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>

char _license[] SEC("license") = "GPL";

SEC("tracepoint/syscalls/sys_enter_execve")
int handle_execve_tp(struct trace_event_raw_sys_enter *ctx) {
    bpf_printk("Hello world");
    return 0;
}`

	// Create a new file or overwrite an existing one
	file, err := os.Create("program.c")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File 'program.c' created successfully!")
}

// DumpBTF runs the bpftool command to dump BTF information into vmlinux.h
func DumpBTF() error {
	// Create or open vmlinux.h for writing
	file, err := os.Create("vmlinux.h")
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

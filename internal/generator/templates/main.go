package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target bpf program program.bpf.c

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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
}

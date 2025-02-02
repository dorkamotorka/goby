# goby

goby (inspired by combining Golang and eBPF in a cartoon fashion) is a CLI to help you generate an initial project setup for developing eBPF program using [Cilium eBPF library](https://github.com/cilium/ebpf).

## Prerequisites

Please check the official Cilium Guide to install the required prerequisites: https://ebpf-go.dev/guides/getting-started/#ebpf-c-program.

Additionally, you need install [bpftool](https://github.com/libbpf/bpftool).

## How to use it

It is a simple as running from the root of this project:
```
go run cmd/goby/main.go init <path> 
```
where `<path> is the directory where you want to create your eBPF project`

## What does it do?

It generates a set of files to speed-up the initial setup of an eBPF project. At it's current state, it takes focuses on simplicity to get you quickly up and running with your idea. 

Namely, it generates:

- `program.bpf.c`: Simple eBPF kernel program with an eBPF Tracepoint example.
- `main.go`: Golang main program, to load and attach eBPF program.
- `vmlinux.h`: Using `bpftool` it generated a header file with kernel struct utilized inside the eBPF program.
- `Makefile`: Wrapper around go command to initialize, generate, build and run eBPF program

## Inside your eBPF Golang project

Check the corresponding files for comments, but initially an example of eBPF Tracepoint is set up.

To run your project, first initialize your Golang project using:
```
make init
```

And then you can generate eBPF skeleton, build the binary and run it using:
```
make build-run
```

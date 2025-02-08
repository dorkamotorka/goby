# goby

goby (inspired by combining Golang and eBPF in a cartoon fashion) is a CLI to help you generate an initial project setup for developing eBPF program using [Cilium eBPF library](https://github.com/cilium/ebpf).

## Prerequisites

Please check the official Cilium Guide to install the required prerequisites: https://ebpf-go.dev/guides/getting-started/#ebpf-c-program.

Additionally, you need install [bpftool](https://github.com/libbpf/bpftool).

## How to use it

First install it from the release:
```
curl -L -o goby https://github.com/dorkamotorka/goby/releases/download/main/goby
chmod +x goby
sudo mv goby /usr/local/bin/
```

Then it is a simple as running from the root of this project:
```
goby init <path> 
```
where `<path> is the directory where you want to create your eBPF project`.

**NOTE**: If the path directory doesn't exists it will create it for you.

## What does it do?

It generates a set of files to speed-up the initial setup of an eBPF project. At it's current state, it takes focuses on simplicity to get you quickly up and running with your idea. 

Namely, it generates:

- `program.bpf.c`: File where you insert your eBPF kernel code.
- `main.go`: Golang main program, to load and attach eBPF programs.
- `vmlinux.h`: Using `bpftool` it generates a header file utilized inside the eBPF Kernel program for interacting with various kernel structs.
- `Makefile`: Wrapper around go command to initialize, generate, build and run eBPF program

## Inside your eBPF Golang project

Check the corresponding files for comments.

To run your project, first initialize your Golang project using:
```
make init
```

And then you can generate eBPF skeleton, build the binary and run it using:
```
make run
```

## Development Tips

Checkout `example_project` which was initialized using `goby` and added an eBPF tracepoint example.

### eBPF Kernel Program (`program.bpf.c`)

In eBPF kernel program it's mainly about writing eBPF subprogram, each attaching to a certain hook, be it an LSM Hook, or XDP, or for tracing.

When you are designing your eBPF program, make sure to check also the eBPF documentation for [eBPF header files](https://docs.ebpf.io/ebpf-library/libbpf/ebpf/) which include multiple useful helper functions.

### eBPF User Space Program (`main.go`)

User space program is mainly responsible for loading and attaching your eBPF Kernel program.

Cilium has a nice [ebpf library documentation](https://pkg.go.dev/github.com/cilium/ebpf) that can help you achieve that. 

**NOTE**: Checkout some of my other repositories for different use cases.

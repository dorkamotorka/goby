
# Goby

**Goby** (inspired by combining Golang and eBPF in a cartoon fashion) is a CLI tool that helps you generate the initial project setup for developing eBPF programs using the [Cilium eBPF library](https://github.com/cilium/ebpf).

![ebpf-gopher](https://github.com/user-attachments/assets/0f1187d2-361a-414f-b933-28cf2416f657)


## Prerequisites

Before getting started, follow the official Cilium guide to install the required prerequisites: [Getting Started Guide](https://ebpf-go.dev/guides/getting-started/#ebpf-c-program).

Additionally, you'll need to install [bpftool](https://github.com/libbpf/bpftool).

## How to Use Goby

1. **Install Goby** from the latest release:

   ```bash
   curl -L -o goby https://github.com/dorkamotorka/goby/releases/download/main/goby
   chmod +x goby
   sudo mv goby /usr/local/bin/
   ```

2. **Initialize your eBPF project** by running the following command from the root of your project:

   ```bash
   goby init <path>
   ```

   Replace `<path>` with the directory where you want to create your eBPF project.

   **Note**: If the directory doesn't exist, Goby will create it for you.

## What Does Goby Do?

Goby generates a set of files to speed up the initial setup of your eBPF project. The current version focuses on simplicity to get you quickly up and running with your idea. Specifically, it generates the following files:

- `program.bpf.c`: The file where you write your eBPF kernel code.
- `main.go`: The Golang program that loads and attaches eBPF programs.
- `vmlinux.h`: A header file generated using `bpftool`, which is used inside the eBPF kernel program to interact with various kernel structs.
- `Makefile`: A wrapper around Go commands to initialize, generate, build, and run the eBPF program.

## Inside Your eBPF Golang Project

Check the corresponding files for helpful comments.

1. **Initialize your Golang project** using:

   ```bash
   make init
   ```

2. **Generate the eBPF skeleton, build the binary, and run it** using:

   ```bash
   make run
   ```

## Development Tips

Check out the `example_project`, which was initialized using Goby and includes an eBPF tracepoint example.

### eBPF Kernel Program (`program.bpf.c`)

In the eBPF kernel program, you'll mostly write eBPF subprograms that attach to various hooks, such as LSM hooks, XDP, or tracing points.

When designing your eBPF program, refer to the [eBPF documentation](https://docs.ebpf.io/ebpf-library/libbpf/), which includes useful helper functions and header files for various eBPF hooks.

### eBPF User Space Program (`main.go`)

The user-space program is responsible for loading and attaching your eBPF kernel program.

Cilium’s [eBPF library documentation](https://pkg.go.dev/github.com/cilium/ebpf) provides useful guidance on how to achieve this.

**Note**: Be sure to explore my other repositories for different use cases.

---

With Goby, getting started with eBPF is simple and fast. Happy coding!

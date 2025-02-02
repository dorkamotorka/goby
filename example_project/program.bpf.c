//go:build ignore

// Here you write an eBPF program
// An example for eBPF Tracepoint
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>

SEC("tracepoint/syscalls/sys_enter_execve")
int handle_execve_tp(struct trace_event_raw_sys_enter *ctx) {
    bpf_printk("Hello world - I'm goby :)");
    return 0;
}

char _license[] SEC("license") = "GPL";
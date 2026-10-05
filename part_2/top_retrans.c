// tcp_retrans.c
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY);
    __uint(max_entries, 1);
    __type(key, __u32);
    __type(value, __u64);
} retransmissions SEC(".maps");

SEC("kprobe/tcp_retransmit_skb")
int BPF_KPROBE(tcp_retransmit_skb, struct sock *sk)
{
    __u32 key = 0;
    __u64 *val = bpf_map_lookup_elem(&retransmissions, &key);
    if (val) {
        (*val)++;
    }
    return 0;
}

char _license[] SEC("license") = "GPL";


sudo bpftrace -e 'tracepoint:block:block_rq_issue { printf("PID %d: %s %d bytes\n", pid, args->rwbs, args->bytes); }'

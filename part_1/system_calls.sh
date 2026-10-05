sudo bpftrace -e 'tracepoint:syscalls:sys_enter_open { printf("PID %d opened %s\n", pid, str(args->filename)); }'

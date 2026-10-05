sudo bpftrace -e 'tracepoint:cuda:kernel_start { printf("Kernel launched: %s\n", str(args->name)); }'

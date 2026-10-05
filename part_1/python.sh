sudo bpftrace -e 'uprobe:python3:PyEval_CallObject { printf("Function called: %s\n", str(arg1)); }'

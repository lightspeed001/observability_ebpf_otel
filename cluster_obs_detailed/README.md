

[Kubernetes Cluster]
       │
       ├── [eBPF Agents] (on each node)
       │       ├── Custom eBPF programs (CPU, memory, network, GPU)
       │       └── BCC tools (quick experiments)
       │
       ├── [OpenTelemetry Collector]
       │       ├── Receives eBPF metrics (via OTLP)
       │       ├── Receives application traces (via OTLP)
       │       └── Exports to Prometheus + Grafana
       │
       ├── [Prometheus]
       │       └── Stores eBPF metrics
       │
       └── [Grafana]
               └── Visualizes metrics + traces


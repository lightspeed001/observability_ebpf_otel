

Example architecture:

ML Application (PyTorch/TensorFlow)
       │
       ▼
[OpenTelemetry Tracing] ────┐
       │                   │
       ▼                   ▼
[Prometheus Metrics]  [Falco Security Events]
       │                   │
       └─────────► [Grafana Dashboard]


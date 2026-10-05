import time
from opentelemetry import metrics
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.exporter.otlp.proto.grpc.metric_exporter import OTLPMetricExporter

# Set up OpenTelemetry metrics
metric_exporter = OTLPMetricExporter(endpoint="http://localhost:4317", insecure=True)
provider = MeterProvider()
metrics.set_meter_provider(provider)
meter = provider.get_meter("ebpf_metrics")

# Create a counter metric
cpu_counter = meter.create_counter("cpu_usage", unit="1")

while True:
    # Read eBPF map (simplified)
    cpu_counter.add(1)  # Replace with actual value
    time.sleep(5)


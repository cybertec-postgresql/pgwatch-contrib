# Prometheus-specific Grafana Dashboards

This directory contains Grafana dashboard configurations specifically designed for Prometheus integration within the pgwatch ecosystem.

## Purpose

Prometheus-specific Grafana dashboards provide:

- Prometheus-formatted metric visualization for PostgreSQL monitoring
- Time-series specific dashboard configurations
- PromQL-based query templates for PostgreSQL metrics
- Integration dashboards for Prometheus alerting

## Content

This directory may contain:

- JSON dashboard definitions optimized for Prometheus data sources
- PromQL query templates for PostgreSQL metrics
- Alerting rule configurations for Prometheus
- Multi-target dashboard templates for Prometheus federation

## Usage

These dashboards are designed to work when pgwatch is configured to use Prometheus as a storage backend, providing:

- PostgreSQL metrics visualization through Prometheus
- Integration with Prometheus alerting and recording rules
- Multi-instance PostgreSQL monitoring dashboards
- Time-series analysis optimized for Prometheus storage

## Integration

These dashboards integrate with pgwatch's Prometheus output capability and are designed to leverage Prometheus's time-series database features for PostgreSQL monitoring.

For more information about pgwatch Prometheus integration, visit the [official documentation](https://pgwat.ch/latest/).

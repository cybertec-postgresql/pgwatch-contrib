# PostgreSQL-specific Grafana Dashboards

This directory contains Grafana dashboard configurations specifically designed for PostgreSQL monitoring within the pgwatch ecosystem.

## Purpose

PostgreSQL-specific Grafana dashboards provide:

- Detailed PostgreSQL performance metrics visualization
- Database-specific health monitoring dashboards
- Query performance analysis dashboards
- Connection and transaction monitoring views

## Content

This directory may contain:

- JSON dashboard definitions for PostgreSQL monitoring
- Custom panel configurations for PostgreSQL metrics
- Templated dashboards for different PostgreSQL versions
- Specialized views for PostgreSQL cluster monitoring

## Usage

These dashboards are designed to work with pgwatch's PostgreSQL metric collection and provide deep insights into:

- Database performance metrics
- Query execution statistics
- Index usage and performance
- Lock and wait event analysis
- Memory and disk utilization

## Integration

These dashboards integrate with pgwatch's metric collection system and are designed to work with PostgreSQL as both the monitoring target and potentially as a storage backend.

For more information about pgwatch PostgreSQL monitoring, visit the [official documentation](https://pgwat.ch/latest/).

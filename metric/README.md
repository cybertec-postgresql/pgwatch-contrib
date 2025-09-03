# Custom Metric Definitions for pgwatch

This directory contains custom metric definitions, SQL queries, and metric configuration files for extending pgwatch PostgreSQL monitoring capabilities.

## Purpose

Custom metric definitions in this directory provide:

- Additional PostgreSQL metric collection queries
- Application-specific monitoring metrics
- Custom SQL-based monitoring definitions
- Extended metric sets for specialized PostgreSQL configurations

## Content

This directory may contain:

- SQL files with custom metric collection queries
- JSON/YAML metric definition configurations
- Custom metric sets for specific PostgreSQL extensions
- Application-specific monitoring queries
- Performance monitoring templates for specialized workloads

## Usage

These custom metrics extend pgwatch's built-in monitoring capabilities:

- Monitor application-specific database metrics
- Collect custom performance indicators
- Track business-specific database KPIs
- Monitor PostgreSQL extension-specific metrics
- Implement custom alerting thresholds

## Integration

Custom metrics integrate with pgwatch's metric collection system and can:

- Be loaded dynamically by the pgwatch monitoring agent
- Extend the default metric collection intervals
- Provide data for custom Grafana dashboards
- Support custom alerting rules and thresholds

For more information about pgwatch custom metrics, visit the [official documentation](https://pgwat.ch/latest/reference/metric_definitions.html).

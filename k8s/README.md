# Kubernetes Deployments for pgwatch

This directory contains Kubernetes manifests, Helm charts, and deployment configurations for running pgwatch PostgreSQL monitoring solution in Kubernetes environments.

## Purpose

Kubernetes configurations in this directory provide:

- Production-ready Kubernetes deployment manifests
- Helm charts for pgwatch components
- ConfigMaps and Secrets management for monitoring configuration
- Service definitions and ingress configurations

## Content

This directory may contain:

- Deployment manifests for pgwatch monitoring agents
- Service definitions for Grafana dashboards
- ConfigMap templates for PostgreSQL monitoring configuration
- Helm charts for complete pgwatch stack deployment
- Ingress configurations for external access
- Persistent volume configurations for metric storage

## Usage

These Kubernetes resources enable deployment of pgwatch in containerized environments:

- Scalable pgwatch monitoring agent deployments
- High-availability Grafana dashboard services
- Persistent storage for monitoring metrics and configuration
- Integration with Kubernetes service discovery
- Resource limits and monitoring for pgwatch components

## Integration

These configurations integrate with Kubernetes native features and can leverage:

- Kubernetes service discovery for PostgreSQL targets
- ConfigMap-based configuration management
- Secret management for database credentials
- Horizontal pod autoscaling for monitoring components

For more information about pgwatch Kubernetes deployment, visit the [official documentation](https://pgwat.ch/latest/).

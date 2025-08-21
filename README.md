# VM SMBIOS Enforcement with Sidecars

This project enforces SMBIOS settings for VMs running on OpenShift Virtualization using hookSidecars.

## Overview

- **Purpose**: Automatically update VM SMBIOS settings (product → "KVM", family → "Virtual Machine")
- **Security**: Provides 3 different methods to prevent users from bypassing or modifying the sidecar configuration
- **Implementation**: KubeVirt hookSidecar image that injects `OnDefineDomain` binary with hardcoded SMBIOS values

## Components

- `kubevirt-sidecar-shim/`: Go-based sidecar container that modifies SMBIOS values
- `policies/`: Three different approaches for policy enforcement
- `workload/vm.yaml`: Example VM with required hookSidecar annotation

## Policy Enforcement Methods

This repository provides **3 methods** for restricting users from customizing hooks:

### 1. Kube Native (ValidatingAdmissionPolicy)
```bash
kubectl apply -f policies/kube-native/vap.yaml
kubectl apply -f policies/kube-native/vapb.yaml
```
Uses Kubernetes native ValidatingAdmissionPolicy for policy enforcement.

### 2. OPA (Open Policy Agent/Gatekeeper)
```bash
kubectl apply -f policies/OPA/constraint.yaml
kubectl apply -f policies/OPA/const-enforce.yaml
```
Uses Gatekeeper with Rego policies for advanced policy logic and enforcement.

### 3. Red Hat ACM (Advanced Cluster Management)
```bash
kubectl apply -f policies/rhacm/policy-const.yaml
```
Uses Red Hat Advanced Cluster Management for multi-cluster policy distribution and compliance.

## Usage

1. Choose one of the three policy enforcement methods above
2. Deploy VMs with the required hookSidecar annotation (see `workload/vm.yaml`)
3. SMBIOS settings are automatically enforced at VM creation time

All three methods ensure VMs must include the exact hookSidecar configuration, preventing unauthorized modifications.

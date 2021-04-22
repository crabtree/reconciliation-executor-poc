# Reconciliation executor POC

## Overview

This repository contains the POC of the reconciliation executor logic, especially focusing on the YAML rendering. One of the reconciliation steps generates valid Kubernetes object declarations. In order to preserve a flexible way of providing object definitions and their respective configuration, we should base our reconciliation logic on the existing rendering engine. This POC presents the possible implementations using Kustomize and Helm. Both approaches render simple deployment and service with support for two landscapes (dev and prod).

The Kustomize approach renders output YAML using the customization and composition approach offered by the tool. Helm example, on the other hand, shows how we can use the tool to render the output YAML based on the chart using the dry-run mode of the install command.

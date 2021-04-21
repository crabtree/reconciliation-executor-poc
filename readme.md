# Reconciliation executor POC

## Overview

This repository contains the POC of the reconciliation executor logic, especially focusing on the yaml rendering. One of the reconciliation steps requires generation of the valid kuberentes objects declarations. In order to perserve felixble way of providing objects definitions and their respective configuration we should base our reconciliation logic on some rendering engine. This POC presents the possible implementations using `Kustomize` and `Helm`. Both approaches render simple deployment and service with support for two landscapes (dev and prod).

The `Kustomize` approach renders output yaml using customization and composition approach offered by the tool. `Helm` example on the other hand, shows how we can use the tool to render the output yaml based on the defined chart using dry-run mode of the install commnad.

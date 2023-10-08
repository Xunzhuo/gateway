---
title:  Compatibility Matrix
---

# Compatibility Matrix

Envoy Gateway relies on the Envoy Proxy and the Gateway API, and runs
within a Kubernetes cluster. Not all versions of each of these products
can function together for Envoy Gateway. Supported version combinations
are listed below; **bold** type indicates the versions of the Envoy Proxy
and the Gateway API actually compiled into each Envoy Gateway release.

+--------------+-----------+-----------+-----------+----------------+
| Envoy        | Envoy     | Rate      | Gateway   | Kubernetes     |
| Gateway      | Proxy     | Limit     | API       | version        |
| version      | version   | version   | version   |                |
+--------------+-----------+-----------+-----------+----------------+
| v0.5.0       | **v1.27   | > **e     | > *       | v1.25, v1.26,  |
|              | -latest** | 059638d** | *v0.7.1** | v1.27          |
+--------------+-----------+-----------+-----------+----------------+
| v0.4.0       | **v1.26   | > **5     | > *       | v1.25, v1.26,  |
|              | -latest** | 42a6047** | *v0.6.2** | v1.27          |
+--------------+-----------+-----------+-----------+----------------+
| v0.3.0       | **v1.25   | > **f     | > *       | v1.24, v1.25,  |
|              | -latest** | 28024e3** | *v0.6.1** | v1.26          |
+--------------+-----------+-----------+-----------+----------------+
| v0.2.0       | **v1.23   |           | > *       | v1.24          |
|              | -latest** |           | *v0.5.1** |                |
+--------------+-----------+-----------+-----------+----------------+
| latest       | **dev     | > *       | > *       | v1.26, v1.27,  |
|              | -latest** | *master** | *v0.8.0** | v1.28          |
+--------------+-----------+-----------+-----------+----------------+

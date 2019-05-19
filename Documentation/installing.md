---
title: Installing
description: Learn about how to install the Dolittle Edge Agent
keywords: Edge
author: einari
weight: 4
---


## Configuration

For the edge agent to work, you'll need to have 

```json
{
    "locationId": "[Guid for location]",
    "nodeId": "[Guid for node]",
    "state": {}
}
```


## Installing as a Daemon

### Systemd

/opt/dolittle.edge

/lib/systemd/system/DolittleEdgeAgent.service

https://fabianlee.org/2017/05/21/golang-running-a-go-binary-as-a-systemd-service-on-ubuntu-16-04/
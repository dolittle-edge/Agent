---
title: Installing
description: Learn about how to install the Dolittle Edge Agent
keywords: Edge
author: einari
weight: 4
---
The Dolittle Edge Agent can be installed from the Linux binaries in our [GitHub releases](https://github.com/dolittle-edge/Agent/releases).

{{% notice information %}}
We've assumed the installation path of `/usr/bin/dolittle.edge` on your system, this includes setting up as a Daemon.
{{% /notice %}}

Pick the version you want and use that in the URL when downloading it as below:

```shell
$ sudo mkdir /usr/bin/dolittle.edge
$ sudo wget https://github.com/dolittle-edge/Agent/releases/download/[RELEASE]/DolittleEdgeAgent -O /usr/bin/dolittle.edge/DolittleEdgeAgent
```

You then have to make it executable:

```shell
$ sudo chmod +x /usr/bin/dolittle.edge/DolittleEdgeAgent
```


## Configuration

For the edge agent to work, you'll need to configure it properly. To do this, you'll need to have the information about which
location and which node it represents. This is found in the Dolittle Edge Studio.

In the folder of the agent binary (`/usr/bin/dolittle.edge`); create a file called `DolittleEdgeAgent.json`.

```json
$ sudo nano /usr/bin/dolittle.edge/DolittleEdgeAgent.json
```

Then configure it as follows:

```json
{
    "locationId": "[Guid for location]",
    "nodeId": "[Guid for node]",
    "state": {}
}
```

{{% notice information %}}
Location and Node identifiers are typically configured in the Edge Studio and is what the APIs
recognize on the other side.
{{% /notice %}}

## Installing as a Daemon

The agent should be running as a Daemon and automatically restarted if failed and also automatically
started when the operating system starts.

### User

We want to run the agent with a specific user. If you don't already have the `dolittle` user created,
add it as follows:

```shell
$ sudo useradd dolittle -s /sbin/nologin -M
```

### Systemd

Download the service definition into where `systemd` is expecting it:

```shell
$ sudo wget https://raw.githubusercontent.com/dolittle-edge/Agent/master/Source/DolittleEdgeAgent.service -O /lib/systemd/system/DolittleEdgeAgent.service
```

Then enable and start it:

```shell
$ sudo systemctl enable DolittleEdgeAgent
$ sudo systemctl start DolittleEdgeAgent
```

To make sure it is running and reporting back to the cloud:

```shell
$ sudo journalctl -f -u DolittleEdgeAgent
May 19 09:08:52 edgeproc2 DolittleEdgeAgent[6501]: Starting
May 19 09:08:52 edgeproc2 DolittleEdgeAgent[6501]: Reporting
May 19 09:08:52 edgeproc2 DolittleEdgeAgent[6501]: {"LocationId":"[some guid]","NodeId":"[some guid]","State":{"ActualMemory":13.700333,"DiskUsage":6.8786774,"Memory":95.04978,"SwapMemory":26.052002}}
```

## Upgrading

When you want to upgrade the client manually, find the version in the [GitHub releases](https://github.com/dolittle-edge/Agent/releases)
and use it in the URL in the script below.

```shell
$ sudo systemctl stop DolittleEdgeAgent
$ sudo wget https://github.com/dolittle-edge/Agent/releases/download/[RELEASE]/DolittleEdgeAgent -O /usr/bin/dolittle.edge/DolittleEdgeAgent
$ sudo systemctl daemon-reload
$ sudo systemctl start DolittleEdgeAgent
```

To make sure it is running and reporting back to the cloud:

```shell
$ sudo journalctl -f -u DolittleEdgeAgent
May 19 09:42:51 edgeproc2 DolittleEdgeAgent[8379]: Dolittle Edge Agent - (C) Dolittle
May 19 09:42:51 edgeproc2 DolittleEdgeAgent[8379]: Starting
May 19 09:42:51 edgeproc2 DolittleEdgeAgent[8379]: Reporting
May 19 09:42:51 edgeproc2 DolittleEdgeAgent[8379]: {"LocationId":"[some guid]","NodeId":"[some guid]","State":{"ActualMemory":13.8554535,"DiskUsage":6.879242,"Memory":95.291115,"SwapMemory":26.046822}}
```
---
title: Building for release
description: Learn about how to build the Dolittle Edge Agent for release
keywords: Edge
author: einari
weight: 3
---
This project uses [Gox](https://github.com/mitchellh/gox) a simplified way
to do cross compilation.

Install it as follows:

```shell
go get github.com/mitchellh/gox
```

You can then run the `build.sh` file inside the `Source` folder.
It will output the binary in the `output` folder within `Source` folder.
The output is targetting Linux.
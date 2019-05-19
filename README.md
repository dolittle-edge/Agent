# Agent

This project represents the Dolittle Edge Agent. Its job is to run at the host level and
provide information / telemetry back to the cloud environment and also provide functionality
for the cloud to call upon to perform actions on the edge.

## Getting started

This project is built using [Go](http://golang.org/).
For those unfamiliar with Go as a language, there is a great [Go by example site](https://gobyexample.com).
If you're using Visual Studio Code, there is a [great extension for Go](https://code.visualstudio.com/docs/languages/go).
A good way to keep the feedback loop tight is to enable the `go.buildOnSave` and the `go.testOnSave` options
for the extension.

## Building for release

This project uses [Gox](https://github.com/mitchellh/gox) a simplified way
to do cross compilation.

Install it as follows:

```shell
go get github.com/mitchellh/gox
```

You can then run the `build.sh` file inside the `Source` folder.
It will output the binary in the `output` folder within `Source` folder.
The output is targetting Linux.
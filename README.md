# Dolittle Edge Agent

This project represents the Dolittle Edge Agent. Its job is to run at the host level and
provide information / telemetry back to the cloud environment and also provide functionality
for the cloud to call upon to perform actions on the edge.

Read more in our [documentation](https://dolittle.io/edge/agent).

## Getting started

This project is built using [Go](http://golang.org/).
For those unfamiliar with Go as a language, there is a great [Go by example site](https://gobyexample.com).
If you're using Visual Studio Code, there is a [great extension for Go](https://code.visualstudio.com/docs/languages/go).
A good way to keep the feedback loop tight is to enable the `go.buildOnSave` and the `go.testOnSave` options
for the extension.

## Tests

Tests are written in a BDD - Specifications by Example - using the [Ginkgo](http://onsi.github.io/ginkgo/)
framework.
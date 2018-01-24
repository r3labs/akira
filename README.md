# akira

Akira provides a complete [go-nats](https://github.com/nats-io/go-nats) interface with a mock connection for testing.

# How to use it

You can use it like:
```go
fc = akira.NewFakeConnector()

// Create a fake connector
xfc := fc.(*akira.FakeConnector)

// Subscribe to a specific event stream
fc.Subscribe("environment.get", func(m *nats.Msg){ println("Hello world!")})
fc.Publish("environment.get", "{...}")

// You can get a history of events accessing xfc.Events
```

## Running Tests

```
make lint
make test
```

## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.

## Versioning

For transparency into our release cycle and in striving to maintain backward
compatibility, this project is maintained under [the Semantic Versioning guidelines](http://semver.org/).

## Copyright and License

Code and documentation copyright since 2015 r3labs.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).

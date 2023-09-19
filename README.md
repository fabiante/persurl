# PersURL

Application to manage and resolve [PURL](https://en.wikipedia.org/wiki/Persistent_uniform_resource_locator) links.

## Usage

### Configuration

The application can be configured using:

- env variables, prefixed with `PERSURL_` (example: `PERSURL_DB_DSN`)
- config files (example: `app.yml`)

Have a look at `example.config.yml` for an example configuration.

## Documentation

Until the documentation becomes large enough, this README will be used to
documentation features and concepts.

### Concept

The application handles PURL domains. A domain is a collection of PURLs which
you can think of as shortened URLs. When opened in a browser (a GET request),
PURLs forward the user to the target URL which is configured by the PURLs
maintainer.

### PURL Support

This project tries it's best to implement PURL support as it can be understood
from a couple of hours researching the topic on the web. It is possible that
some implementations deviate from what PURL was ment to do.

If you find any critical deviations or design flaws which you think should be adressed,
please open an issue to resolve that.

## Contributing

### Testing

This project tries to implement tests in a behaviour-driven style. If you add / change features, ensure that
test specifications ensure correct behaviour. Test drivers execute test specifications.

#### Load Tests

Load tests can enabled via the env variable `TEST_LOAD=1` or setting `test_load` to `1` in your config file.

These run the application and generate load by running multiple agents simulating user behaviour.
The motivation of these tests is to ensure that the application can be used for a large user base which
create large amounts of data.

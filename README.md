# PersURL

Application to manage and resolve [PURL](https://en.wikipedia.org/wiki/Persistent_uniform_resource_locator) links.

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

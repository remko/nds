# Google Cloud Datastore Go utility clients

**⚠️ This package is still in the design phase, with no working code yet. I am
still looking for interest, input, and help on this package**

Approximately source-compatible implementations of the [Google Cloud Datastore
API Go package](https://pkg.go.dev/cloud.google.com/go/datastore), with
alternative and middleware implementations for unit testing, caching, ...
Offers many features of the
[ndb](https://googleapis.dev/python/python-ndb/latest) package, and more.

Client implementations:

- [**Datastore**](https://pkg.go.dev/github.com/remko/nds/datastore): Passthrough to the
  [`cloud.google.com/datastore`](https://pkg.go.dev/cloud.google.com/go/datastore)
  package, for the actual connection to Google Cloud Datastore.
- **Memory**: in-memory implementation, for use in unit tests
- **Cache**: Middleware client that caches entities locally (in the current
  request context) and/or globally (e.g. in [Redis](https://redis.io) or global
  memory), similar to the local and global cache of the
  [`ndb`](https://googleapis.dev/python/python-ndb/latest) package.


This package also provides utility functions for working with the datastore:

- Functions to safely load and save incomplete structs from datastore. This
  avoids data loss when working with structs that do not perfectly map to the
  datastore (e.g. when working with legacy data or other clients)

## Documentation

- [Package documentation](https://pkg.go.dev/github.com/remko/nds)

## Related packages

- <https://pkg.go.dev/github.com/qedus/nds>

    - Cached datastore API implementation
    - Follows the same API as the standard datastore package
    - Only adds a cache layer (no abstraction, no memory implementation)
    - Inspired by NDB

- <https://github.com/mercari/datastore>

    - Does much more than just wrapping: different API, different implementation, ...

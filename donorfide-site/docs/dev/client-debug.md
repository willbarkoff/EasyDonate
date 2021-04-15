---
layout: docs
title: Client Debug Mode
subtitle: Load a client from the file system.
---

Donorfide provides _client debug mode_, which allows you to load a client from the file system. This is useful because it allows you to develop the client without having to rebuild Donorfide every time changes are made.

Donorfide will load the client from `./client/dist`.

To enable client debug mode, run Donorfide with the `--client-debug` flag:

```shell
$ donorfide --client-debug
```

When Donorfide starts, it will print a message indicating that client debug mode is enabled:

```
8:41PM INF Donorfide is running in client debug mode. For more information, visit https://donorfide.org/docs/client-debug
```

The API tester can also be enabled by setting the environment variable `DONORFIDE_CLIENT_DEBUG` to the value `1`.
```bash
$ DONORFIDE_CLIENT_DEBUG=1 donorfide # command-line envioment variables
$ echo "DONORFIDE_CLIENT_DEBUG=1" >> .env && donorfide # .env file
```

Client debug mode can be used efficiently in production.

---

When developing the Donorfide client, it is suggested that you use client debug mode. This can be done by running both `yarn run` and `donorfide` simultaneously. This will continuously update the client, allowing you to make changes to the TypeScript of the client without rebuilding Donorfide:

```shell
$ donorfide --client-debug &
$ cd client
$ yarn run
```
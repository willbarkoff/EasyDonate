---
layout: docs
title: API Tester
subtitle: Test the Donorfide API during development.
---

An API tester can be enabled for Donorfide development to more easily enable API testing.

To enable the API tester, run Donorfide with the `--enable-tester` flag:

```shell
$ donorfide --enable-tester
```

When Donorfide starts, it will print a message indicating that the API tester is enabled:

```
8:41PM INF The API tester is enabled. For more information, visit https://donorfide.org/docs/api-tester
```

The API tester can also be enabled by setting the environment variable `DONORFIDE_API_TESTER` to the value `1`.
```bash
$ DONORFIDE_API_TESTER=1 donorfide # command-line envioment variables
$ echo "DONOFIDE_API_TESTER=1" >> .env && donorfide # .env file
```

Though there are no additional security consequences of leaving the API tester enabled in production environments, it is recommended that you disable it.
---
layout: docs
title: Install
subtitle: Install Donorfide for your platform
---

Donorfide has a few installation options. These include:
- Binary installation
- Building from source
- and cloud provider deploy.

The easiest is binary installation.

{% include toc.md %}

## Binary Installation
<div class="field is-grouped is-grouped-multiline">
{% include badge.html value="recommended" color="primary"%}
{% include badge.html key="difficulty" value="easy" color="success" %}
</div>

Donorfide distributes several _binaries_, or single files that can be run to deploy Donorfide. These are available for the following architectures and operating systems:

### Binaries
The <strong>Linux Intel 64-bit processor</strong> binary is suitable for use on most Linux servers, including (but not limited to) Ubuntu, Debian, Red Hat, and CentOS. The <strong>Windows Intel 64-bit processor</strong> binary is suitable for use on most Windows servers.

<a class="button is-primary is-outlined">
	Linux Intel 64-bit processor
</a>
<a class="button is-primary is-outlined">
	Windows Intel 64-bit processor
</a>

### Additional Binaries
These additional binaries are useful if you have a less conventional server setup.


## Building from Source
<div class="field is-grouped is-grouped-multiline">
{% include badge.html key="difficulty" value="hard" color="warning" %}
</div>
Donorfide can be built from source using its source code, available at [GitHub](https://github.com/willbarkoff/donorfide).

First, install the required software to build Donorfide. This includes the [Go Programming Language](https://golang.org), [Node.js](https://nodejs.org), and [Yarn](https://yarnpkg.org).

Next, clone the Donorfide repository.
```shell
$ git clone https://github.com/willbarkoff/donorfide.git # clone using HTTPS, or
$ git clone git@github.com:willbarkoff/donorfide.git # clone using SSH
```

The next step is to build the Donorfide client. It is required that the client be built before the server, as the server produces a binary including the contents of the client. Move to the client's directory.

```shell
$ cd donorfide/client
```

Next, install the required dependencies.
```shell
$ yarn
```

Finally, create a production build of the client.
```shell
$ yarn build
```

This will create a build suitable for production in the directory `donorfide/client/dist`. Once the client finishes building, you need to build the server. Move to the server's directory.
```shell
$ cd ..
```

Next, install the required dependencies, and build and install Donorfide.
```shell
$ go get
$ go install
```

This will produce a Donorfide build at `$GOPATH/bin/donorfide`.
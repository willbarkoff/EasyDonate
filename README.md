<div align="center">

![Donorfide](./logotype.png)

<p>
✨ Donorfide makes it easy for your nonprofit organization to collect donations.
</p>

Learn more at [donorfide.org](https://donorfide.org).

</div>

---

> ⚠️ This is the documentation for Donorfide version 3.0, which is still very much a work in progress, and should not be used in production environments. Donorfide v3.0 is a continuation of the EasyDonate project, with an emphasis on usability. 
> 
> For **EasyDonate 2.0**, the latest stable version of this software, see the `master` branch:
> ```shell
> % git checkout -b master
> ```

Donorfide is an open-source, online payment processing and donor management platform that allows your nonprofit to process donations easily, so you can focus on doing good.

- Collect donations online in a highly customizable manner
- Manage recurring donations

---

For a sample installation see [demo.donorfide.org](https://demo.donorfide.org).

Donorfide prides itself on being easy to install. For installation instructions, see [donorfide.org/docs/install](https://donorfide.org/docs/install).

The remainder of this document focuses on technical information about Donorfide. If you're interested in that, read on! If you want to get a Donorfide installation up and running, visit our [installation instructions](https://donorfide.org/docs/install).

---

This repository is split into 3 sections
- [`donorfide-client`](./donorfide-client) holds the code related to the React client for Donorfide.
- [`donorfide-server`](./donorfide-server) holds the code related to the Go server for Donorfide.
- [`donorfide-site`](./donorfide-site) holds the code related to the Jekyll site at donorfide.org.
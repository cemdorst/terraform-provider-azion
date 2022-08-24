Azion Terraform Provider (WIP)
==================

- Documentation: https://registry.terraform.io/providers/cemdorst/azion/latest/docs

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.18 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/cemdorst/terraform-provider-azion`

```sh
$ mkdir -p $GOPATH/src/github.com/cemdorst; cd $GOPATH/src/github.com/cemdorst
$ git clone git@github.com:cemdorst/terraform-provider-azion.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/cemdorst/terraform-provider-azion
$ make
```

Using the provider
----------------------

See the [Azion Provider documentation](https://registry.terraform.io/providers/cemdorst/azion/latest/docs) to get started using the Azion provider.

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.18+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.


In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

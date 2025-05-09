A Neocities client written in Go
================================

Upload files to you [Neocities](https://neocities.org/) website from the
comfort of your own terminal.

[![Build Status](https://github.com/azhai/neocities-go/actions/workflows/test.yml/badge.svg)](https://github.com/azhai/neocities-go/actions/workflows/test.yml)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/azhai/neocities-go)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/azhai/neocities-go#license-mit)

## Installation


Or, if you have [Go](http://golang.org/) installed:

    go install github.com/azhai/neocities-go@latest

## Usage

[![Cat!](https://neocities.org/img/cat.png)](https://neocities.org/)

First you need to set two environment variables:

```bash
export NEOCITIES_USER=<user>
export NEOCITIES_PASS=<pass>
```

Alternatively you can use the `NEOCITIES_API_KEY` variable.

Then you should be able to upload files to your website:

    neogo upload foo.html bar.js folder/baz.jpg

You can also delete files:

    neogo delete foo.html folder/baz.jpg

You get a list of available commands by default:

```bash
$ neogo
usage: neogo <command> [<args>]

Commands:
    upload       Upload files to Neocities
    upload-root  Upload local files to webroot
    delete       Delete files from Neocities
    delete-all   Delete all remote files
    info         Info about Neocities websites
    key          Neocities API key
    list         List files on Neocities
    version      Show neocities client version

Help for a specific command:
   help [command]
```

## Donate

[![Support Neocities](https://neocities.org/img/support-us.png)](https://neocities.org/donate)

NeoCities is funded by donations. If you’d like to contribute, you can help to pay for server costs using Bitcoin or PayPal.

## License (MIT)

Copyright (c) 2014-2024 [Peter Hellberg](https://c7.se)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

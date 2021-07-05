mod-redirect
============

A simple `go-get` custom redirect server

[![Go Report Card](https://goreportcard.com/badge/gophers.dev/cmds/mod-redirect)](https://goreportcard.com/report/gophers.dev/cmds/mod-redirect)
[![Build Status](https://travis-ci.com/shoenig/mod-redirect.svg?branch=master)](https://travis-ci.com/shoenig/mod-redirect)
[![GoDoc](https://godoc.org/gophers.dev/cmds/mod-redirect?status.svg)](https://godoc.org/gophers.dev/cmds/mod-redirect)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/mod-redirect.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/shoenig/mod-redirect.svg)](LICENSE)

# Project Overview

Module `gophers.dev/cmds/mod-redirect` provides an an API and web server for
responding with custom redirects to `?go-get=1` from Go compilers / proxies.

# Getting Started

The `mod-redirect` package can be installed by running
```bash
$ go install gophers.dev/cmds/mod-redirect/cmd/mod-redirect@latest
```

# Example Usage

#### Add a module redirect

The `/v1/set` endpoint can be used to add a new module redirect.
The following example will create the redirect:
`<host>/pkgs/my/name/space` => `https://github.com/user/project`
```bash
$ curl -H "X-MR-Key: abc123" -XPOST "localhost:1300/v1/set" -d '{"kind":"pkgs", "namespace":"my/name/space", "vcs":"git", "destination":"https://github.com/user/project"}'
```
Then, the server will respond appropriately with meta content for `go-get=1` requests
```bash
$ curl "https://<host>/pkgs/my/name/space?go-get=1"
```
Resulting in HTML like
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta
            name="go-import"
            content="<host>/pkgs/my/name/space git https://github.com/user/project"
    >
    <title>go-get</title>
</head>
</html>
```

#### List configured module redirects

The `/v1/list` can be used to list existing configured module redirects.
```bash
$ curl -s "https://gophers.dev/v1/list" | jq .
```
Results in a JSON response
```json
[
  {
    "kind": "cmds",
    "namespace": "mod-redirect",
    "vcs": "git",
    "destination": "https://github.com/shoenig/mod-redirect"
  }
]
```

# Configuration

See [example-config.json](hack/example-config.json) for an example configuration file.

If `mod-redirect` will be running behind a reverse proxy (e.g. Caddy), it is
recommended to enable rate-limiting (to avoid DOS attacks), BUT to allow for
bursts of requests, due to the nature of how the Go compiler makes requests.
Something like
```
gophers.dev
	proxy / localhost:1300
	gzip
	log	stdout
	errors stderr
	ratelimit / 3 20
}
```

# Contributing

The `gophers.dev/cmds/mod-redirect` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `gophers.dev/cmds/mod-redirect` module is open source under the [BSD-3-Clause](LICENSE) license.

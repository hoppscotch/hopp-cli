⚠️⚠️⚠️⚠️⚠️⚠️

**This project is now archived. If you want to use Hoppscotch as a CLI client for CI/CD purposes, please use the new [Hoppscotch CLI](https://docs.hoppscotch.io/cli).**

⚠️⚠️⚠️⚠️⚠️⚠️

# Hoppscotch CLI [![hoppscotch](https://img.shields.io/badge/Made_for-Hoppscotch-hex_color_code?logo=Postwoman)](https://hoppscotch.io) [![Go Report Card](https://goreportcard.com/badge/github.com/athul/pwcli)](https://goreportcard.com/report/github.com/athul/pwcli)

Send HTTP requests from terminal and Generate API Docs. An alternative to cURL, httpie ⚡️

## Installation

### From Source

- Clone the repo

```shell
$ git clone https://github.com/hoppscotch/hopp-cli.git
```

- Build and install

```shell
$ make
$ sudo make install
```

### From Binary

- You can download prebuilt binaries from the [Releases](https://github.com/hoppscotch/hopp-cli/releases) page.
- **Supported platforms**:
  - Linux (x64, x86)
  - Mac (x64)
  - Windows (x64, x86)

> **IMPORTANT: Not tested on Windows, please leave your feedback/bugs in the Issues section**

### Arch GNU/Linux

- You can install from [AUR](https://aur.archlinux.org/)
- There are three different packages available

Name          | Link                                              | Description
------------- | ------------------------------------------------- | -----------------------------
hopp-cli-bin  | https://aur.archlinux.org/packages/hopp-cli-bin/  | Pre-built binary
hopp-cli      | https://aur.archlinux.org/packages/hopp-cli/      | Compiled from latest release
hopp-cli-git  | https://aur.archlinux.org/packages/hopp-cli-git/  | Compiled from latest commit

### Homebrew

Install by

```shell
brew install athul/tap/hopp-cli
```

### Windows

You can download pre-built binaries from the [Releases](https://github.com/hoppscotch/hopp-cli/releases) page.

Alternatively, you can install `hopp-cli` via [Scoop](https://scoop.sh/):

```shell
scoop install hopp-cli
```

## Usages

Putting Simply: **Just pass the URL to the request method**

### Basic Commands

- GET : `$ hopp-cli get <url>`
- POST: `$ hopp-cli post <url>`
- PATCH: `$ hopp-cli patch <url>`
- PUT : `$ hopp-cli put <url>`
- DELETE: `$ hopp-cli delete <url>`

Example for a POST request:

```shell
$ hopp-cli post https://reqres.in/api/users/2 -c js -b '{"name": "morp","job": "zion resident"}'
```

### Extra Commands

- `send` for testing multiple endpoints
- `gen` for generating API docs from Collection

#### SEND

This can be used to test multiple endpoints from the `hoppscotch-collection.json` file.

> The output will only be the `statuscode`

Example:

```shell
$ hopp-cli send <PATH to hoppscotch-collection.json>
```

Sample output:

![send-output](/assets/send.png)

#### GEN

The `gen` command generates the API documentation from `hoppscotch-collection.json` file and serves it as a static page on port `1341`.

Example:

```shell
$ hopp-cli gen <PATH to hoppscotch-collection.json>
```

Sample Hosted site: https://hopp-docsify.surge.sh/

Powered by [Doscify](https://docsify.js.org)

Flags:

- `browser` or `b` to toggle whether the browser should open automatically [Boolean]
- `port` or `p` for specifying the port where the server should listen to [Integer]

### There are 3 Authentication Flags

_(optional)_

- `-t` or `--token` for a Bearer Token for Authentication
- `-u` for the `Username` in Basic Auth
- `-p` for the `password` in Basic Auth

### There are 2 flags especially for the data management requests like POST, PUT, PATCH and DELETE

- `-c` or `--ctype` for the _Content Type_
- `-b` or `--body` for the Data Body, this can be of json, html or plain text based on the request.

> Enclose the body in Single Quotes(\')

**Content Types can be of**

|Short Code|Content Type|
|:---:|:---:|
|`js`|`application/json`|
|`html`|`text/html`|
|`xml`|`application/xml`|
|`plain`|`text/plain`|

### Include Arbitrary Headers

- `-H` or `--header` may be specified multiple times to include headers with the request.

Example:

```shell
$ hopp-cli get -H 'X-Api-Key: foobar' -H 'X-Api-Secret: super_secret' https://example.com/api/v1/accounts
```

### Providing a Request Body via stdin

In addition to `-b`/`--body`, you may provide a request body via stdin.\
If you combine this method with the `-b` flag, the body provided with `-b` will be ignored.

**Example with Pipes**

```shell
$ echo '{"foo":"bar"}' | hopp-cli post -c js http://example.com
```

**Example with Redirection**

```shell
$ cat myrequest.json
{
  "foo": "bar"
}

$ hopp-cli post -c js http://example.com <myrequest.json
```

### Providing a Request Body via text-editor

In addition to providing request body via `-b / --body` flag and stdin,
you can also use `-e / --editor` flag which opens default text-editor in your system.

**Example:**

```shell
$ hopp-cli post https://reqres.in/api/users/2 -c js -e
```

It will preferrably open editor based on `$EDITOR` environment variable.

**For example:**
If the environment variable is `$EDITOR=code` it will open VSCode for request-body input. Else, it will use default editor value based on the OS.

| OS      | Default Editor |
| ------- | -------------- |
| Linux   | `nano`         |
| macOS   | `nano`         |
| Windows | `notepad`      |

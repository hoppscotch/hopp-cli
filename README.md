# Hoppscotch CLI [![hoppscotch](https://img.shields.io/badge/Made_for-Hoppscotch-hex_color_code?logo=Postwoman)](https://hoppscotch.io) [![Go Report Card](https://goreportcard.com/badge/github.com/athul/pwcli)](https://goreportcard.com/report/github.com/athul/pwcli)

Send HTTP requests from terminal and Generate API Docs. An alternative to cURL, httpie ⚡️

## Installation

### From Script

```bash
$ sh -c "$(curl -sL https://git.io/getpwcli)"
```

### From Source

- Clone the repo

```
$ git clone https://github.com/hoppscotch/hopp-cli.git
```

- Build and install

```
$ make

$ sudo make install
```

### From Binary

- You can find the Binaries in Gzipped form from the [Releases](https://github.com/athul/pwcli/releases) page  
  **Supports**
- Linux(x64,x86)
- Mac(x64)
- Windows(x64,x86)

> **IMPORTANT: Not tested on Windows, please leave your feedback/bugs in the Issues section**

### Homebrew

Install by `brew install athul/tap/hopp-cli`

## Usages

Putting Simply: **Just pass the URL to the request method**

### Basic Commands

- GET : `$ hopp-cli get <url> `
- POST: `$ hopp-cli post <url> `
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

**SEND**: This can be used to test multiple endpoints from the `hoppscotch-collection.json` file.

> The output will only be the `statuscode`

Example : `hopp-cli send <PATH to hoppscotch collection.json>`  

Sample Output:
![](/assets/send.png)

---

**GEN**: Gen command Generates the API Documentation from  `hoppscotch-collection.json` file and serves it as a Static Page on port `1341`  
Example: `hopp-cli gen <PATH to hoppscotch collection.json>`

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

### There are 2 flags especially for the data management requests like POST,PUT,PATCH and DELETE

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

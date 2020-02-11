# PostWoman CLI   [![Build Status](https://travis-ci.com/athul/pwcli.svg?token=udLtq6DyJs4Gxpze9nqX&branch=master)](https://travis-ci.com/athul/pwcli)[![Postwoman](https://img.shields.io/badge/Made_for-Postwoman-hex_color_code?logo=Postwoman)](https://postwoman.io)
Use Postwoman's CLI direct from your terminal.

# Installation
### From Script
```bash
$ sh -c "$(curl -sL https://git.io/getwcli)"
```
### From Source
- Clone the repo
- Build using `go build`
- Move Binary to `/usr/local/bin/`
### From Binary
- You can find the Binaries in Gzipped form from the [Releases](https://github.com/athul/pwcli/releases) page      
**Supports**
- Linux(x64,x86)
- Mac(x64)
- Windows(x64,x86)

### Homebrew
Install by `brew install athul/tap/pwcli`

> **IMPORTANT: Not tested on Windows, please leave your feedback/bugs in the Issues section**

# Usages

Putting Simply: **Just pass the URL to the request method**


- GET : `pwcli get <url> -t/--token <token> -u <username for basic auth> -p <password for basic auth>`
- POST: `pwcli post <url> < -t/-u/-p > -c/--content type <content type> -b/--body <body>`
- PATCH: `pwcli patch <url> < -t/-u/-p > -c/--content type <content type> -b/--body <body>`
- PUT : `pwcli put <url> < -t/-u/-p > -c/--content type <content type> -b/--body <body>`
- DELETE: `pwcli delete <url> < -t/-u/-p > -c/--content type <content type> -b/--body <body>`

**Content Types can be of**
`html`   :   `text/html`
`js`     :   `application/json`,
`xml`    :   `application/xml`
`plain`  :   `text/plain`,


`` Extra
**SEND**: This can be used to test multiple endpoints from the `postwoman-collection.json` file. The output will only be the `statuscode`.       
RUN: `pwcli send <PATH to postwoman collection.json>`      
OUTPUT: 
![](/assets/send.png)
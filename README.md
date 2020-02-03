# PostWoman CLI   [![Build Status](https://travis-ci.com/athul/pwcli.svg?token=udLtq6DyJs4Gxpze9nqX&branch=master)](https://travis-ci.com/athul/pwcli)[![Postwoman](https://img.shields.io/badge/Made_for-Postwoman-hex_color_code?logo=Postwoman)](https://postwoman.io)
Use Postwoman's CLI direct from your terminal.

# Installation

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
```
NAME:
   Postwoman CLI - Test API endpoints without the hassle

USAGE:
   cli [global options] command [command options] [arguments...]

VERSION:
   0.0.1

DESCRIPTION:
   Made with <3 by Postwoman Team

COMMANDS:
   get      Send a GET request
   post     Send a POST Request
   put      Send a PUT Request
   patch    Send a PATCH Request
   delete   Send a DELETE Request
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
----
## GET
**Usage**  
```
NAME:
   cli get - Send a GET request

USAGE:
   cli get [command options] [arguments...]

OPTIONS:
   --url value    The URL/Endpoint you want to check (default: "https://reqres.in/api/users")
   --token value  Send the Request with Bearer Token
   -u value       Add the Username
   -p value       Add the Password
```
## POST
**Usage**   
```
NAME:
   cli post - Send a POST Request

USAGE:
   cli post [command options] [arguments...]

OPTIONS:
   --url value    The URL/Endpoint you want to check (default: "https://reqres.in/api/users")
   --token value  Send the Request with Bearer Token
   -u value       Add the Username
   -p value       Add the Password
   --ctype value  Change the Content Type (default: "application/json")
   --body value   Body of the Post Request
```

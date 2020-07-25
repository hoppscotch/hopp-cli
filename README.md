# PostWoman CLI  [![Postwoman](https://img.shields.io/badge/Made_for-Postwoman-hex_color_code?logo=Postwoman)](https://postwoman.io) [![Go Report Card](https://goreportcard.com/badge/github.com/athul/pwcli)](https://goreportcard.com/report/github.com/athul/pwcli)

Send HTTP requests from terminal. An alternative to cURL, httpie ⚡️

# Installation

### From Script

```bash
$ sh -c "$(curl -sL https://git.io/getpwcli)"
```

### From Source

- Clone the repo

```
$ git clone https://github.com/athul/pwcli
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

### Homebrew

Install by `brew install athul/tap/pwcli`

> **IMPORTANT: Not tested on Windows, please leave your feedback/bugs in the Issues section**

# Usages

Putting Simply: **Just pass the URL to the request method**
## Basic
- GET : `pwcli get <url> `
- POST: `pwcli post <url> `
- PATCH: `pwcli patch <url>`
- PUT : `pwcli put <url>`
- DELETE: `pwcli delete <url>`

Example for a POST request: 
`pwcli post https://reqres.in/api/users/2 -c js -b '{"name": "morp","job": "zion resident"}`

### Extra

**SEND**: This can be used to test multiple endpoints from the `postwoman-collection.json` file. The output will only be the `statuscode`.  
Example : `pwcli send <PATH to postwoman collection.json>`  
o/p:
![](/assets/send.png)


### There are 3 Authentication Flags
*(optional)*        
- `-t` or `--token` for a Bearer Token for Authentication
- `-u` for the `Username` in Basic Auth
- `-p` for the `password` in Basic Auth
### There are 2 flags especially for the data management requests like POST,PUT,PATCH and DELETE
- `-c` or `--ctype` for the *Content Type*

- `-b` or `--body` for the Data Body, this can be of json, html or plain text based on the request. 
  > Enclose the body in Single Quotes(\')

**Content Types can be of**         
`html` : `text/html`   
`js` : `application/json`   
`xml` : `application/xml`   
`plain` : `text/plain`  

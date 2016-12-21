## Standard CA Roots for Golang

This package provides an HTTP client preloaded with the [standard Mozilla CA roots](https://hg.mozilla.org/releases/mozilla-release/raw-file/default/security/nss/lib/ckfw/builtins/certdata.txt). This can be helpful for TLS clients in environments that do not provide a standard set of CA roots (e.g. extremely minimal Docker containers).

### Security warning

If you ship code that uses package then you will be shipping CA roots that will not update except when you ship updates to your code. This means that when a root CA is revealed to be untrustworthy, such as in the recent Wosign incident, your code will continue trusting an untrustworthy CA until you update it. In contrast, if you get CA roots from the underlying operating system (which is the default in Golang) then this issue will be taken care of by updates from your operating system vendor. This package should therefore only be used in situations where it will be updated regularly, and only when you cannot get CA roots from the underlying operating system.

### Quick start

First install the package:
```shell
go get github.com/alexflint/stdroots
```

Then use `stdroots.Client`, which is a `*http.Client`:
```go
resp, err := stdroots.Client.Get("https://www.google.com")
```

### Reproducibility

The certificates are embedded in the `bindata.go` file. You can reproduce this file by running `make update`, which will pull `certdata.txt` from the latest Mozilla release and generate a set of certificates under `certs/`, then finally generate `bindata.go`.

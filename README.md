## Standard CA Roots for Golang

This package provides an HTTP client preloaded with the [standard Mozilla CA roots](https://hg.mozilla.org/releases/mozilla-release/raw-file/default/security/nss/lib/ckfw/builtins/certdata.txt). This can be helpful for TLS clients in environments that do not provide a standard set of CA roots (e.g. extremely minimal Docker containers).

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

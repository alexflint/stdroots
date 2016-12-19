//go:generate go-bindata -o bindata.go -pkg stdroots certs/...

package stdroots

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
)

// Client is an HTTP client preloaded with the certificates bundled
// with this package
var Client *http.Client

// Pool is an x509 certificate pool containing each of the certifications
// bundled with this package.
var Pool *x509.CertPool

func init() {
	// get list of certs
	fs, err := AssetDir("certs")
	if err != nil {
		// assets are bundled with binary so it makes sense to panic here
		panic(fmt.Sprintf(`AssetDir("certs") failed: %v`, err))
	}

	// load and parse certs
	Pool = x509.NewCertPool()
	for _, f := range fs {
		ok := Pool.AppendCertsFromPEM(MustAsset("certs/" + f))
		if !ok {
			// assets are tested before release so it makes sense to panic here
			panic(fmt.Sprintf("could not load cert from %s: AppendCertsFromPEM returned false", f))
			log.Println("failed to append cert:", f)
		}
	}

	// setup HTTP client
	tlsConfig := &tls.Config{
		RootCAs: Pool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	Client = &http.Client{Transport: transport}
}

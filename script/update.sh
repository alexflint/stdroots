#!/bin/bash

# The raw certs from mozilla
CERTDATA="https://hg.mozilla.org/releases/mozilla-release/raw-file/default/security/nss/lib/ckfw/builtins/certdata.txt"

# Delete and recreate the directory
rm -rf certs
mkdir certs || exit 1
cd certs

# certdata.txt must be in current directory
curl -o certdata.txt $CERTDATA || exit 1

# run python script to convert to crts
python2.7 ../certdata2pem.py

# remove files that should not end up in bindata
rm *.p11-kit *.txt

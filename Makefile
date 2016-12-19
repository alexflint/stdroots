# The raw certs from mozilla
CERTDATA:="https://hg.mozilla.org/releases/mozilla-release/raw-file/default/security/nss/lib/ckfw/builtins/certdata.txt"

update:
	# Delete and recreate the directory
	rm -rf certs
	mkdir certs || exit 1

	# certdata.txt must be in current directory
	curl -o certs/certdata.txt $(CERTDATA) || exit 1

	# run python script to convert to crts
	cd certs && python2.7 ../script/certdata2pem.py || exit 1

	# remove files that should not end up in bindata
	rm certs/*.p11-kit certs/*.txt

	# generate new bindata file
	go generate


test:
	go test -v

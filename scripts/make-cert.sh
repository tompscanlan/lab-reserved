[ -f temp.key -a -f temp.crt ] || \
openssl req -new \
	-newkey rsa:4096 \
	-days 365 \
	-nodes -x509 \
	-subj "/C=US/ST=KY/L=Louisville/O=none/CN=none.com" \
	-keyout temp.key \
	-out temp.crt

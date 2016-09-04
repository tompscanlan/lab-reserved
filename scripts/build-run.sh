go build ./cmd/labreserved-server && \
  ./labreserved-server --port 2080 --tls-port 20443 --tls-key temp.key  --tls-certificate temp.crt

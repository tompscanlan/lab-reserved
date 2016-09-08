# current bug on Mac prevents running from scratch.
#FROM scratch
FROM golang:1.6
MAINTAINER Tom Scanlan <tscanlan@vmware.com>

<<<<<<< HEAD
ENV BLOB_ENDPOINT http://blobs.vmwaredevops.appspot.com/api/v1/blobs
ENV BLOB_ID 7357
EXPOSE 80 443

=======
>>>>>>> 1b69b57d27c436964a5c79c74bd51c4235900778
# Add the microservice
ADD labreserved-server /labreserved-server
ADD temp.crt /lab.crt
ADD temp.key /lab.key

<<<<<<< HEAD
CMD ["/labreserved-server", "--tls-host", "0.0.0.0", "--host", "0.0.0.0", "--port", "80", "--tls-port", "443", "--tls-key", "/lab.key", "--tls-certificate", "/lab.crt", "--blob-endpoint", "$BLOB_ENDPOINT", "--blob-id", "$BLOB_ID"]
=======
EXPOSE 80 443
CMD ["/labreserved-server", "--tls-host", "0.0.0.0", "--host", "0.0.0.0", "--port", "80", "--tls-port", "443", "--tls-key", "/lab.key", "--tls-certificate", "/lab.crt"]
>>>>>>> 1b69b57d27c436964a5c79c74bd51c4235900778

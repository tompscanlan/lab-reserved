# current bug on Mac prevents running from scratch.
FROM scratch
#FROM golang:1.6
MAINTAINER Tom Scanlan <tscanlan@vmware.com>

ENV BLOB_ENDPOINT http://blobs.vmwaredevops.appspot.com/api/v1/blobs
ENV TEAM_ID 7357
EXPOSE 80 443

# Add the microservice
ADD labreserved-server /labreserved-server
ADD temp.crt /lab.crt
ADD temp.key /lab.key

CMD ["/labreserved-server", "--tls-host", "0.0.0.0", "--host", "0.0.0.0", "--port", "80", "--tls-port", "443", "--tls-key", "/lab.key", "--tls-certificate", "/lab.crt"]

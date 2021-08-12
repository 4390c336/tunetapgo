FROM golang:1.16
WORKDIR /opt/tunetap-go/
RUN apt update && DEBIAN_FRONTEND=noninteractive apt-get install -y tshark isc-dhcp-client
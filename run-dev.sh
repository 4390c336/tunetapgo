docker run --name tunetap-go --cap-add=NET_ADMIN --device /dev/net/tun:/dev/net/tun -v $(pwd)/:/opt/tunetap-go/ -it --rm tunetap-go bash
#docker build -t admiralu82/moproxy .
FROM ubuntu:20.04
LABEL version="1.0"
LABEL description="Moproxy - transparent proxy"
LABEL org.opencontainers.image.authors="admiralu82@yandex.ru"
COPY moproxy /
COPY run.sh /
COPY out /
COPY out.png /
COPY out.html /

RUN apt update && apt install -y net-tools iproute2 iputils-ping nano traceroute iptables

CMD /run.sh

#docker network create -d macvlan --subnet=10.31.78.0/24 --gateway=10.31.78.1 -o parent=eno1.33 moproxy_vlan33
#docker run --name moproxy --network moproxy_vlan33 --ip=10.31.78.250 --cap-add=NET_ADMIN admiralu82/moproxy
#docker exec -it moproxy bash

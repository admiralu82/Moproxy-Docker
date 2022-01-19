#!/bin/bash
iptables -t nat -A OUTPUT -p tcp -m multiport --dports 80,443 -j REDIRECT --to-port 8080
iptables -t nat -A PREROUTING -p tcp -m multiport --dports 80,443 -j REDIRECT --to-port 8080

param=""
if [ -n "$1" ]
then
  param=$1
else
  param="10.0.31.52:3128"
fi
/moproxy --port 8080  --log-level warn --remote-dns --stats-bind 0.0.0.0:8081 --http $param

#!/bin/sh

VPN_USER="$1"

if [ -z "$VPN_USER" ]; then
  echo "Usage: $0 username" >&2
  echo "Example: $0 jordi" >&2
  exit 1
fi

cp /etc/ppp/chap-secrets /etc/ppp/chap-secrets.bak
sed "/\"$VPN_USER\" /d" /etc/ppp/chap-secrets.bak > /etc/ppp/chap-secrets
cp /etc/ipsec.d/passwd /etc/ipsec.d/passwd.bak
sed "/$VPN_USER:/d" /etc/ipsec.d/passwd.bak > /etc/ipsec.d/passwd

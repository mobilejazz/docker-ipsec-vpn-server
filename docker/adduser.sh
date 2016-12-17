#!/bin/sh

VPN_USER="$1"

if [ -z "$VPN_USER" ]; then
  echo "Usage: $0 username" >&2
  echo "Example: $0 jordi" >&2
  exit 1
fi

case "$VPN_USER" in
  *[\\\"\']*)
    echo "VPN credentials must not contain any of these characters: \\ \" '" >&2
    exit 1
    ;;
esac

SHARED_SECRET=$(cut -d'"' -f2 /etc/ipsec.secrets)
echo "Shared secret: $SHARED_SECRET"

VPN_PASSWORD="$(LC_CTYPE=C tr -dc 'A-HJ-NPR-Za-km-z2-9' < /dev/urandom | head -c 20)"
VPN_PASSWORD_ENC=$(openssl passwd -1 "$VPN_PASSWORD")
echo "Password for user is: $VPN_PASSWORD"

echo '"'$VPN_USER'"' l2tpd '"'$VPN_PASSWORD'"' '*' >> /etc/ppp/chap-secrets
echo $VPN_USER:$VPN_PASSWORD_ENC:xauth-psk >> /etc/ipsec.d/passwd

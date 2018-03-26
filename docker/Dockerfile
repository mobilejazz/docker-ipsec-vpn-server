FROM debian:stretch
MAINTAINER Jordi Gimenez <info@mobilejazz.com>

ENV REFRESHED_AT 2018-03-26
ENV SWAN_VER 3.23

WORKDIR /opt/src

RUN apt-get -yqq update \
    && DEBIAN_FRONTEND=noninteractive \
       apt-get -yqq --no-install-recommends install \
         wget dnsutils openssl ca-certificates kmod \
         iproute gawk grep sed net-tools iptables \
         bsdmainutils libcurl3-nss \
         libnss3-tools libevent-dev libcap-ng0 xl2tpd \
         libnss3-dev libnspr4-dev pkg-config libpam0g-dev \
         libcap-ng-dev libcap-ng-utils libselinux1-dev \
         libcurl4-nss-dev flex bison gcc make \
    && wget -t 3 -T 30 -nv -O "libreswan.tar.gz" "https://github.com/libreswan/libreswan/archive/v${SWAN_VER}.tar.gz" \
    || wget -t 3 -T 30 -nv -O "libreswan.tar.gz" "https://download.libreswan.org/libreswan-${SWAN_VER}.tar.gz" \
    && tar xzf "libreswan.tar.gz" \
    && rm -f "libreswan.tar.gz" \
    && cd "libreswan-${SWAN_VER}" \
    && sed -i '/docker-targets\.mk/d' Makefile \
    && printf 'WERROR_CFLAGS =\nUSE_DNSSEC = false\nUSE_SYSTEMD_WATCHDOG = false\n' > Makefile.inc.local \
    && make -s base \
    && make -s install-base \
    && cd /opt/src \
    && rm -rf "/opt/src/libreswan-${SWAN_VER}" \
    && apt-get -yqq remove \
         libnss3-dev libnspr4-dev pkg-config libpam0g-dev \
         libcap-ng-dev libcap-ng-utils libselinux1-dev \
         libcurl4-nss-dev flex bison gcc make \
         perl-modules perl \
    && apt-get -yqq autoremove \
    && apt-get -y clean \
    && rm -rf /var/lib/apt/lists/* \
    && rm -f /etc/ppp/chap-secrets /etc/ipsec.d/passwd /etc/ipsec.secrets

COPY ./run.sh /opt/src/run.sh
COPY ./adduser.sh /opt/src/adduser.sh
COPY ./lsusers.sh /opt/src/lsusers.sh
COPY ./rmuser.sh /opt/src/rmuser.sh
RUN chmod 755 /opt/src/run.sh /opt/src/adduser.sh /opt/src/lsusers.sh /opt/src/rmuser.sh

EXPOSE 500/udp 4500/udp

VOLUME ["/lib/modules", "/etc/ppp/chap-secrets", "/etc/ipsec.d/passwd", "/etc/ipsec.secrets"]

CMD ["/opt/src/run.sh"]
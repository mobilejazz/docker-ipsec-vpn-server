![DockerIpSec](https://raw.githubusercontent.com/mobilejazz/metadata/master/images/banners/mobile-jazz-docker-ipnsec-server.png)

﻿IPsec VPN Server on Docker
--------------------------

[![Build Status](https://travis-ci.org/mobilejazz/docker-ipsec-vpn-server.svg)](https://travis-ci.org/mobilejazz/docker-ipsec-vpn-server)

Docker image to run an IPsec VPN server, with support for both `IPsec/L2TP` and `IPsec/XAuth ("Cisco IPsec")`.

Based on [Lin Song's IPsec VPN Server on Docker](https://github.com/hwdsl2/docker-ipsec-vpn-server) modified specially for **multiple users**.

## Install Docker

Follow [these instructions](https://docs.docker.com/engine/installation/) to get Docker running on your server.

## How to use this image

### Start the IPsec VPN server

Download [this repo](https://github.com/mobilejazz/docker-ipsec-vpn-server/archive/master.zip), `unzip` it and `cd` into it.

Start a new Docker container with the following command:

```
./start.sh
```

Once you have the service up, you will need at least one user to use it.

### Add a new user

Create a new VPN user with the adduser command. For example, create a user `john-ipad` like this:

```
./adduser.sh john-ipad
```

This will generate an individual password for this user (user specific, usually called "password") and also display the shared key of the server (same for all users, usually called "PSK" or "Pre-Shared Key").

The user will be available immediately, there is no need to restart the server.

**IMPORTANT**: Due to a limitation in the IPSec protocol design, several devices can not connect to the same server behind the same NAT router. We recommend creating a separate account **for each device** a user owns. This will also make revocation of credentials easier if a user lost a device.

### List users in the system

You can list all VPNs with the command:

```
./lsusers.sh
```

### Remove a user

You can remove a user like this (following the `john-ipad` example):

```
./rmuser.sh john-ipad
```

### Check server status

To check the status of your IPsec VPN server, you can pass `ipsec status` to your container like this:

```
./status.sh
```

## Next steps

Get your computer or device to use the VPN. Please refer to:

[Configure IPsec/L2TP VPN Clients](https://github.com/hwdsl2/setup-ipsec-vpn/blob/master/docs/clients.md)   
[Configure IPsec/XAuth ("Cisco IPsec") VPN Clients](https://github.com/hwdsl2/setup-ipsec-vpn/blob/master/docs/clients-xauth.md)

If you get an error when trying to connect, see [Troubleshooting](https://github.com/hwdsl2/setup-ipsec-vpn/blob/master/docs/clients.md#troubleshooting).

Enjoy your very own VPN! :sparkles::tada::rocket::sparkles:

## Technical details

There are two services running: `Libreswan (pluto)` for the IPsec VPN, and `xl2tpd` for L2TP support.

Clients are configured to use [Google Public DNS](https://developers.google.com/speed/public-dns/) when the VPN connection is active.

The default IPsec configuration supports:

* IKEv1 with PSK and XAuth ("Cisco IPsec")
* IPsec/L2TP with PSK

The ports that are exposed for this container to work are:

* 4500/udp and 500/udp for IPsec

## Extending the configuration

The default configuration will work out of the box in most cases. However, you might want to tweak some little settings,
like the routing table, or maybe something specific to your environment. If you mount a `/pre-up.sh` script, it will be executed
before starting the VPN.

## Backing up your VPN configuration

When using the start script a new `etc` directory will be created. You can back up this directory.

## Build from source code

Advanced users can download and compile the source yourself from GitHub:

```
git clone https://github.com/mobilejazz/docker-ipsec-vpn-server.git
cd docker-ipsec-vpn-server/docker
docker build -t mobilejazz/docker-ipsec-vpn-server .
```

## See also

* [IPsec VPN Server on Ubuntu, Debian and CentOS](https://github.com/hwdsl2/setup-ipsec-vpn)
* [IKEv2 VPN Server on Docker](https://github.com/gaomd/docker-ikev2-vpn-server)

## License

    This project is distributed with a Creative Commons Attribution and Share Alike license. See LICENSE.md file to read the complete license.

Made with ❤️ from Barcelona by [Mobile Jazz](https://mobilejazz.com), the Web and App development company for startups.

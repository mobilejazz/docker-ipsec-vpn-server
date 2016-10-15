#!/bin/sh

grep -v '^#' /etc/ppp/chap-secrets | cut -d' ' -f1 | cut -d'"' -f2

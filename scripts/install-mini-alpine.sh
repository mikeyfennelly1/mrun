#!/bin/bash

arch=$(uname -m)

case "$arch" in
    x86_64)
        tarballUrl=https://dl-cdn.alpinelinux.org/alpine/v3.22/releases/x86_64/alpine-minirootfs-3.22.1-x86_64.tar.gz
        tarballName=alpine-minirootfs-3.22.1-x86_64.tar.gz
        ;;
    aarch64)
        tarballUrl=https://dl-cdn.alpinelinux.org/alpine/v3.22/releases/aarch64/alpine-minirootfs-3.22.1-aarch64.tar.gz
        tarballName=alpine-minirootfs-3.22.1-aarch64.tar.gz
        ;;
    riscv64)
        tarballUrl=https://dl-cdn.alpinelinux.org/alpine/v3.22/releases/riscv64/alpine-minirootfs-3.22.1-riscv64.tar.gz
        tarballName=alpine-minirootfs-3.22.1-riscv64.tar.gz
        ;;
    *)
        echo "Unknown architecture: $arch"
        echo "See the alpinelinux downloads page for an alternative filesystem: https://alpinelinux.org/downloads/"
        exit 1
        ;;
esac

curl -LO "${tarballUrl}" \
  && mkdir rootfs \
  && tar -xzf "${tarballName}" -C rootfs

echo "Root filesystem for your container unpacked at $(pwd)/rootfs"
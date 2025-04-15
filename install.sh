#!/bin/zsh

# Check if the script is run as root
if [[ $(id -u) -ne 0 ]]; then
    echo "You must be root to run this script. Use 'sudo'.." >&2
    exit 1
fi

mrunBinaryPath="/usr/bin/mrun"

go build -o ${mrunBinaryPath}
sudo chown 0:0 ${mrunBinaryPath}
sudo chmod +s ${mrunBinaryPath}

sudo setcap cap_sys_admin,cap_audit_write=ep ${mrunBinaryPath}


alias mrun="${mrunBinaryPath}"
source ~/.zshrc
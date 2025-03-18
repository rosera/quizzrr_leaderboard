# Leaderboard

## Create VM

A micro VM is required.

## Install packages

```bash
#!/bin/bash
# STARTUP-START
apt-get update -y

# Update package lists and install required packages
apt-get install -y curl jq git
```

## Add Go
```bash
# Download Go
curl -LO https://go.dev/dl/go1.23.4.linux-amd64.tar.gz 

# Install Go
rm -rf /usr/local/go && tar -C /tmp -xzf go1.23.4.linux-amd64.tar.gz
```

## Create a USER
```
# Env Var
export USER="api-dev"

# Create a user
useradd $USER -m -p Password01 -s /bin/bash -c 'Developer Account'
```

## Add Systemctl Process

Move the `glb` Go binary to the user account

```bash
sudo su -u api-dev
mkdir ~/glb && cd $_
```

Clone the GLB repository
```bash
git clone https://github.com/rosera/quizzrr_leaderboard.git
```

Create a file `/home/api-dev/glb/glb` and copy the binary from the quizzrr_leaderboard/glb folder.

Create a service `glb.service`

```bash
[Unit]
Description=Quizzrr Leaderboard service

[Install]
WantedBy=multi-user.target

[Service]
Type=simple
ExecStart=/home/api-dev/glb/glb
WorkingDirectory=/home/api-dev/glb
Restart=always
RestartSec=5
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=%n
```

Start the service
```bash
systemctl start glb.service
```

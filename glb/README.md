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
```bash
# Env Var
export USER="api-dev"

# Create a user
useradd $USER -m -p Password01 -s /bin/bash -c 'Developer Account'
```

## Add Systemctl Process

Switch to the developer account
```bash
sudo su -u api-dev
```

Clone the GLB repository
```bash
git clone https://github.com/rosera/quizzrr_leaderboard.git
```

```bash
mkdir ~/glb && cd $_
```

Move the `glb` Go binary to the user account
```bash
tar -xvf ~/quizzrr_leaderboard/glb/glb.tar.gz
```

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

Exit the developer account

```bash
exit
```

Enable the service
```bash
systemctl enable glb.service
```

Start the service
```bash
systemctl start glb.service
```

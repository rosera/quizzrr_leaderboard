# Leaderboard

## Transfer Data
### SCP Transfer

Transfer file to virtual machine using Ssh SCP
__NOTE:__
External IP is not enabled, so tranfer via IAP on internal IP.

```
gcloud compute scp glb arcade-leaderboard:/tmp --zone=us-central1-c --project=qwiklabs-resources
```

### Connect to Virtual Machine
Connect to the compute instance using SSH

```
gcloud compute ssh arcade-leaderboard --zone us-central1-c
```

## Create a Service

### Create VM

A micro VM is required.

### Install packages

```bash
#!/bin/bash
# STARTUP-START
apt-get update -y

# Update package lists and install required packages
apt-get install -y curl jq git
```

### Add Go
```bash
# Download Go
curl -LO https://go.dev/dl/go1.23.4.linux-amd64.tar.gz 

# Install Go
rm -rf /usr/local/go && tar -C /tmp -xzf go1.23.4.linux-amd64.tar.gz

# Use Go Build for Linux Platform
CGO_ENABLED=0 GOOS=linux go build -o glb
```

Amend main.go variable to match:

- [x] Internal IP of Host Machine
- [x] PORT of Service e.g. 8080


Build the binary

### Create a USER
```bash
# Env Var
export USER="api-dev"

# Create a user
useradd $USER -m -p Password01 -s /bin/bash -c 'Developer Account'
```

### Add Systemctl Process

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

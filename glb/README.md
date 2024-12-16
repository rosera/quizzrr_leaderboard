# Leaderboard

A quick in-memory leaderboard api to capture scores.

| Tool | Version |
|------|---------|
| Go   | 1.23    |

## Configure
```bash
go build
```

## Run
```bash
./glb
```


## Run as Systemd

Use the following configuration to create a service on a Debian virtual machine.

```
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

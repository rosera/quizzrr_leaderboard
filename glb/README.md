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


## Environment 

A Terraform configuration is included to setup the environment.
The compute instance will run a startup-script and install the Go executables.

1. Move to the `tf` folder

2. Initiate Terraform
```bash
terraform init
```

3. Validate the Terraform
```bash
terraform validate
```

4. Create Terraform state 

Add:

| Field | Description |
|-------|-------------|
| project | PROJECT_ID |
| region  | PROJECT_ID |

```bash
terraform plan --out tf.state
```

5. Initialise the resource using state

```bash
terraform apply
```

## Build Application



## Run as Systemd

The virtual machine includes a user account `api-dev`.
Run the application under this account.
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

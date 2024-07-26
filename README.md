# HBaaS

HBaaS (Hackathon Backend as a Service) is a service that provides general-purpose backend functionalities tailored for hackathons.
You can use HBaaS to quickly build a backend for your hackathon project.

We provide SDKs for various platforms such as Web, iOS, and Android.

## Features

- User Authentication
- Realtime Messaging
- Key-Value Store
- File Storage
- And more...

## Getting Started

### Prerequisites

- Golang
- Docker
- Google Cloud SDK

### Installation

```shell
# Prepare tools
make install-tools
docker-compose up -d

# Prepare application config
export HBAAS_CONFIG_FILEPATH=$(pwd)/config/default.json
# Optional: Use custom config file
cp config/default.json config/{custom_config_name}.json  
export HBAAS_CONFIG_FILEPATH=$(pwd)/config/{custom_config_name}.json
```

### Start Server

```shell
make run-api-server
```

# Profile Service

## Introduction

A Profile Service is a API service that provides user profile management. It allows users to create, update, delete and view their profiles, add, remove and view user devices.

# Features

- Create a user profile
- Get a user profile
- Update a user profile
- Delete a user profile
- Add the multiple devices to a user profile
- Get the devices of a user profile
- Update the devices of a user profile
- Remove devices from a user profile

# Configuration

The service is configured with a [config.ini](config.ini) file found at `/app/config.ini` or another file specified using the --config flag.

## Installation

To install the service, you need to have Go installed on your machine. You can download and install Go from the official website: [https://golang.org/](https://golang.org/).

After installing Go, clone the repository and build the service using the following commands:

```bash
git clone https://github.com/dredfort42/rupl-user-profile.git
cd rupl-user-profile
go build -o profile ./cmd/profile/main.go
```

## Usage

### Running the service

To start the service, run the following command:

```bash
./profile
```

The service will start and listen on the host and port specified in the configuration file.

### Running the service in Debug mode

To run the service in Debug mode, set the DEBUG environment variable to true before starting the service:

```bash
env DEBUG=true ./profile
```

The service will start in Debug mode and print additional information to the console.

### Running the service with a specific configuration file

To run the service with a specific configuration file, set the --config flag with the path to the configuration file while starting the service:

```bash
./profile --config /path/to/my_config.ini
```

### Running the service in Docker

To run the service in Docker, build the Docker image using the following command:

```bash
docker build -t profile .
```

After building the Docker image, run the service using the following command:

```bash
docker run -p 4242:4242 profile
```

The service will start and listen on port 4242.

### Running the Service with Docker Hub Images

To run the service using Docker Hub images, use the following command:

```bash
docker run -p 4242:4242 dredfort/profile:latest
```

This will download the service from Docker Hub and start it, listening on port 4242.

## API

The API endpoints are described in the [openapi.yaml](/api/openapi.yml) file.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

# tempDB Key-Value Store HTTP API Service
This is a Go-based key-value store HTTP API service that allows you to set, retrieve, and search for key-value pairs. It also includes Prometheus monitoring for latency, HTTP status codes, and the total number of keys in the database.

# Table of Contents

* Prerequisites
* Getting Started
   - Build and Push Docker Image
   - Deploy on Minikube
* Usage
   - Set a Key-Value Pair
   - Get the Value for a Key
   - Search for Keys
* Monitoring with Prometheus
* Contributing

# Prerequisites
## Before you begin, ensure you have the following tools installed:
* Go
* Docker
* Minikube

# Getting Started

# Build and Push Docker Image

1. Clone this repository:

    ```bash
    git clone git@github.com:TirushV/tempDB.git
    cd tempDB
    ```

2. Build your Docker image:

    ```bash
    docker build -t <image-name>:latest
    ```

# Deploy on Minikube

1. Deploy the Kubernetes resources:

    ```bash
    kubectl apply -f deployment.yaml
    kubectl apply -f service.yaml
    kubectl apply -f ingress.yaml
    ```

2. Map your domain to Minikube's IP in /etc/hosts: (If testing from local)

    ```bash
    echo "$(ip) your-domain.com" | sudo tee -a /etc/hosts
    ```

3. Access your application at http://localhost:8080.

# Usage
## Set a Key-Value Pair
To set a key-value pair, make a POST request to /key/set with a JSON payload:

```bash
{
  "key": "your-key",
  "value": "your-value"
}
```

```bash
curl -X POST http://localhost:8080/key/set -H "Content-Type: application/json" -d '{
  "key": "your-key",
  "value": "your-value"
}'
```

## Get the Value for a Key
To retrieve the value for a key, make a GET request to /get/your-key. The value will be returned as a plain string.

```bash
curl http://localhost:8080/get/<key> 
```

## Search for Keys
You can search for keys using the /search endpoint with the prefix and suffix query parameters:

* /search?prefix=your-prefix returns keys with the specified prefix.
* /search?suffix=your-suffix returns keys with the specified suffix.

```bash
curl "http://localhost:8080/search?prefix=abc"
curl "http://localhost:8080/search?suffix=-1"
```

## Monitoring with Prometheus
This service includes Prometheus monitoring for latency, HTTP status codes, and the total number of keys in the database. You can access the Prometheus dashboard at http://localhost:8080/metrics.
(Yet to update this)

** Metrics to watch for
- total_keys_in_db
- http_status_codes_total
- http_request_duration_seconds

Contributing
Contributions are welcome! Please read the Contributing Guidelines for details.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Feel free to customize this README with specific installation, configuration, and usage details relevant to your project. Additionally, you may want to include information about API documentation, testing procedures, and any other relevant project-specific details.
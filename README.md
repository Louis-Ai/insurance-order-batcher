# Insurance Order Batcher

This is a simple Go microservice that receives insurance order requests and batches them into CSV files.

## Overview

The application exposes an HTTP POST endpoint at `/orders` where other services can submit order data in JSON format. It validates the postcode of the order and then adds the order to an in-memory batch. Once the batch reaches a configured size, the orders are written to a CSV file in a specified directory.

## Prerequisites

* Go 1.24 or higher installed.
* Docker installed.

## Running with Docker

1.  Clone the repository
    ```bash
    git clone <repository_url>
    cd insurance-order-batcher
    ```

2.  Build the Docker image
    ```bash
    docker build -t order-batcher .
    ```

3.  Run the Docker container
    ```bash
    docker run -p 8080:8080 \
        -v /tmp/docker_orders:/data/orders \
        -e OUTPUT_DIRECTORY=/data/orders \
        -e BATCH_SIZE=5 \
        order-batcher
    ```
    Adjust the volume mapping (`-v`) to a local directory where you want the CSV files to be saved. You can also override the environment variables here.

## Testing API Endpoint

### `POST /orders`

http://localhost:8080/orders 

Accepts JSON requests with the following structure:

```json
{
  "customerID": "UW234",
  "address": {
    "addressLineOne": "123 street",
    "townCity": "town",
    "postcode": "SW10 0RA"
  },
  "order_timestamp": "2025-01-01 00:00:00"
}
```
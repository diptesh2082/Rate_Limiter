# Tollbooth Rate Limiting API

Welcome to the Tollbooth Rate Limiting API project! This project implements a simple HTTP API with rate limiting using the [Tollbooth](https://github.com/didip/tollbooth) library. The API provides a `/ping` endpoint that responds with a JSON message, demonstrating how to limit the number of requests a client can make.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the API](#running-the-api)
- [Testing the API](#testing-the-api)
- [API Endpoints](#api-endpoints)
- [Code Structure](#code-structure)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Features

- **Rate Limiting**: Control the number of requests to the `/ping` endpoint.
- **JSON Responses**: Clear and structured responses for both successful and failed requests.
- **Simple Code Structure**: Easy to understand and extend.

## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (version 1.16 or later)
- A working Go environment

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/tollbooth-rate-limiting.git
   cd tollbooth-rate-limiting
   ```

2. Install the required dependencies:

   ```bash
   go get github.com/didip/tollbooth
   ```

### Running the API

To run the API, execute the following command in your terminal:

```bash
go run tollbooth/main.go
```

The server will start on `http://localhost:8080`.

## Testing the API

You can test the `/ping` endpoint using `curl`:

```bash
curl -i http://localhost:8080/ping
```

To test the rate limiting, you can send multiple requests in a loop:

```bash
for i in {1..6}; do curl -i http://localhost:8080/ping; done
```

## API Endpoints

### `GET /ping`

- **Description**: Returns a JSON response indicating the status of the API.
- **Response**:
  - **Status 200**: Successful response with a message.
  - **Status 429**: Too Many Requests if the rate limit is exceeded.

**Example Response**:
```json
{
  "status": "success",
  "body": "pong"
}
```

## Code Structure

- `tollbooth/main.go`: The main application file containing the API logic and rate limiting setup.

## Contributing

Contributions are welcome! To contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Make your changes and commit them (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

Please ensure your code adheres to the existing style and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Tollbooth](https://github.com/didip/tollbooth) for the rate limiting library.
- Special thanks to all contributors and the open-source community for their support and inspiration.
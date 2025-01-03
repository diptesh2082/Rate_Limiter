# Tollbooth Rate Limiting API

This project implements a simple HTTP API with rate limiting using the [Tollbooth](https://github.com/didip/tollbooth) library. The API provides a `/ping` endpoint that responds with a JSON message, demonstrating how to limit the number of requests a client can make.

## Features

- Rate limiting for the `/ping` endpoint.
- JSON responses for successful and failed requests.
- Simple and easy-to-understand code structure.

## Getting Started

### Prerequisites

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

```

The server will start on `http://localhost:8080`.

### Testing the API

You can test the `/ping` endpoint using `curl`:

```

To test the rate limiting, you can send multiple requests in a loop:

```

### API Endpoints

#### `GET /ping`

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

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Tollbooth](https://github.com/didip/tollbooth) for the rate limiting library.
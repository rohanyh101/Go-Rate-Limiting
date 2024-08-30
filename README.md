# Go Rate Limiting Implementations

This repository showcases three different rate-limiting implementations in Go:

1. **Token Bucket**: A leaky bucket rate limiter that allows bursts of requests and then refills over time.
2. **Per Client Rate Limiting**: Individual rate limiting based on the client IP or identifier.
3. **Tollbooth**: An external rate-limiting library used to throttle requests.

## Project Structure

Each implementation is housed in its own directory:

- `token-bucket/`
- `per-client-rate-limiting/`
- `tollbooth/`

Each directory contains:

- **`rateLimiter` Middleware**: Checks user eligibility and determines how many requests can be sent per burst.
- **`endpointHandler` HTTP Handler**: The main endpoint executed after rate limiting checks. 

###  Running the Project
1. Clone the repository:

  ```bash
  git clone https://github.com/rohanyh101/go-rate-limiting.git
  cd go-rate-limiting
  ```
2. Navigate to the desired rate-limiting implementation:

  ```bash
  cd token-bucket  # or per-client-rate-limiting, tollbooth
  ```

3. Run the Go application:

  ```bash
  go run main.go
  ```

### Some CURL to test,

- **Request,**
  ```bash
      >for i in {1..10}; do echo ""; curl --location --request GET 'http://localhost:8080/ping' \
      --header 'Content-Type: application/json'; done
  ```

- **Response,**
  ```bash
    {"status":"success","body":"HI, you have reached the endpoint!, how may I help you ?"}
    
    {"status":"success","body":"HI, you have reached the endpoint!, how may I help you ?"}
    
    {"status":"success","body":"HI, you have reached the endpoint!, how may I help you ?"}
    
    {"status":"success","body":"HI, you have reached the endpoint!, how may I help you ?"}
    
    {"status":"request denied","body":"reached api capacity, try some time later"}
    
    {"status":"request denied","body":"reached api capacity, try some time later"}
    
    {"status":"request denied","body":"reached api capacity, try some time later"}
    
    {"status":"request denied","body":"reached api capacity, try some time later"}
    
    {"status":"request denied","body":"reached api capacity, try some time later"}
    
    {"status":"request denied","body":"reached api capacity, try some time later"}
  ```

Each implementation will start a server and apply the rate-limiting logic before allowing access to the `endpointHandler`.

## Contributing

Feel free to open issues or submit pull requests if you find any bugs or have suggestions for improvements.

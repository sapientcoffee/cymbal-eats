# Coffee Service

A simple Go service for getting a list of coffees.

## Running the service

1.  **Build the Docker image:**
    ```bash
    docker build -t coffee-service .
    ```

2.  **Run the Docker container:**
    ```bash
    docker run -p 8080:8080 coffee-service
    ```

3.  **Access the API:**
    Open your browser or use curl to access the API:
    ```bash
    curl http://localhost:8080/coffees
    ```

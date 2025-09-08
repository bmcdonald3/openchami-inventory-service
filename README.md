# Hardware Inventory API

This service provides a REST API for storing and managing hardware inventory, location, and event data.

## Building and Running

### Commands
First, build the executable file. Then, run the file to start the server.

# Build the executable from the project root
```bash
go build -o inventory-api ./cmd/openchami-inventory-service/
```

# Run the server. It will listen on port 8080.
```bash
./inventory-api
```

## Testing Endpoints

Once the server is running, you can test the mock endpoints using `curl` from a separate terminal.

### List All Devices
```bash
curl -i http://localhost:8080/inventory/v1/devices
```

### Get a Specific Device by ID
```bash
curl -i http://localhost:8080/inventory/v1/devices/c3d4e5f6-a1b2-4c1d-8e9f-0c1d2e3f4a5b
```

## API Specification

The formal API definition is written in `TypeSpec` and is forthcoming...

### Generating the OpenAPI Specification

TODO
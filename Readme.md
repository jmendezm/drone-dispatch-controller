# Drone Dispatch Controller

API REST service for communication with the drone fleet

### How build
```bash
go mod vendor
go build -o drone-dispatch-controller
```

### How run
```bash
./drone-dispatch-controller --config=[/path/to/config/file]
```

### Configuration file
There is an example of a configuration file in the config folder

#### Example:
```json
{
  "log_level": 1,
  "show_logs": true,
  "listen_port": 9000
}
```

### Run using docker
#### Build the image
```bash
docker build -t drone-dispatch-controller .
```
#### Run the image
```bash
docker run -p 8080:8080 drone-dispatch-controller
```

### Storage
The service uses an in-memory database
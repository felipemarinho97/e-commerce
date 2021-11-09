# e-commerce REST
e-commerce REST service.

## Configuration

### Rest API
| Name | Type | Default | Description |
|:-----|:----:|:-------:|:------------|
| HASH_PORT | int | 8080 | Port used by web rest API to receive requests. |
| HASH_HOST | str | localhost | Address used by web rest receiver to receive requests. |
| HASH_GRPC_PORT | int | 50051 | Port used by gRPC server to receive requests. |
| HASH_GRPC_HOST | str | localhost | Address used by ecommerce-rest API to connect with gRPC server. |

## Run
Inside [root](./..) folder:

```bash
go run main.go
```

Inside [rest](rest/) folder:

```bash
go run main.go
```

## Docker
### Build
Build gRPC server:

```bash
docker build -t hash/ecommerce ./..
```

Build Rest server:

```bash
docker build -t hash/ecommerce-rest .
```

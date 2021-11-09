# Ecommerce API
## - ecommerce

Simple API written in golang to handle the checkout functionality.
It consists in a gRPC micro-service that only handles the method "Checkout".

### Envs
| Name | Type | Default | Description |
|:-----|:----:|:-------:|:------------|
| HASH_PORT | int | 8080 | Port used by gRPC API to receive requests. |
| HASH_HOST | str | localhost | Host used by gRPC API to receive requests. |
| HASH_DISCOUNT_ADDR | str | localhost:5050 | Address used by gRPC API to comunicate with discount micro-service. |
| HASH_PRODUCTS_MOCK_FILE | str | ./products.json | Path to the file containing the products. |

## - ecommerce-rest

This is a proxy to the gRPC service. The mapping to REST is done by the `ecomerce-rest` micro-service using the grpc-gateway plugin. 

The `ecommerce-rest` implement the following endpoints:
 - POST /v1/products

See more info about this service [here](/rest/).

---

## Running the stack
Inside the project folder **e-commerce** run the following command:

```bash
docker-compose up -d --build
```

### Manual testing of the service

    $ curl -X POST -H "Content-Type: application/json" http://localhost:80/v1/products --data-raw '{
    "products": [
            {
                "id": 2,
                "quantity": 1
            }
        ]
    }'

## API Swagger
After running the docker-compose stack, the Swagger UI will be available at [http://localhost:8080/](http://localhost:8080/).

For more information about the backend project's Swagger API, go to [this document](doc/openapiv2/protos/api/ecommerce.swagger.json). This file is generated by the `protoc-gen-openapiv2` plugin and the API definitions can be seen and edited in the [`ecommerce.proto`](protos/api/ecommerce.proto) file.


## Generating proto and OpenAPI files

Make sure you have the following binaries installed.
- `protoc-gen-grpc-gateway`
- `protoc-gen-openapiv2`
- `protoc-gen-go`
- `protoc-gen-go-grpc`

Then type:

    $ buf generate


## Tests
Unit tests cover the following percentages of the packages as shown below:
- db: 100%
- discount: 83.3% (This percentage ignores the protobuf generated code)
- api: 64%

### Running the unit tests
Run the following command at the root of the project:

    $ go test -v -coverprofile=coverage.out ./...


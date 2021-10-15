# go-supermarket

There are 4 models:

- Customer
- Product
- Transaction
- Purchase

A supermarket sells a variety of `Product`s. Only registered `Customer`s can shop at the supermarket. A transaction has `Purchases`, which is a list of items purchased by a customer. Each `Purchase` has information like `Quantity`, `PricePerUnit` and the `Product`'s information.

Refer to the `struct` types in the `model` directory for more detailed information.

## Initialization

Execute `go run cmd/generatedb/main.go` to generate some dummy data in a `sqlite` database called `test.db` in the project root directory.

Execute `go run main.go` to start the server. By default the server starts on port 80.

## Todo

- Deep dive on validator

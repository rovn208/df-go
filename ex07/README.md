# Assignment for Day 7

## Simple E-commerce
Simple e-commerce project with basic functionalities 
- Cart management: 
  - Add to cart
  - Remove from cart
  - Checkout
- Product management APIs
  - Create a new product
  - Update a product's details
  - Delete a product by its ID
  - Retrieve a list of all products


## Authentication
Basic Auth with 2 defined authenticated users

| Username | Password |
|----------|----------|
| admin    | admin    |
| dfgo     | awesome  |

Using [Basic Auth Header Generator](https://www.debugbear.com/basic-auth-header-generator) to convert the credential above
in order to use the authenticated endpoints

## Endpoints
BASE_URL:
* Local: localhost:8080/

Basic endpoints:
* Products - Authenticated [via Basic Auth](#authentication)
  * GET - `{{BASE_URL}}`/api/v1/products
  * POST - `{{BASE_URL}}`/api/v1/products
  * DELETE - `{{BASE_URL}}`/api/v1/products/`:productId`
  * PUT - `{{BASE_URL}}`/api/v1/products/`:productId`
* Cart
  * POST - `{{BASE_URL}}`/api/v1/carts/add
  * DELETE - `{{BASE_URL}}`/api/v1/carts/remove
  * POST - `{{BASE_URL}}`/api/v1/carts/checkout
## Running local
- Start database
> docker compose up -d
- Start application
> go run cmd/main.go

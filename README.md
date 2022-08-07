# Propery Finder Final Project

This project is the implementation of a basic basket system based on the requirements of the final project specification.

## Technologies/Libraries Used

* Go as language
* Gin for http request handling
* SQLc for database actions
* Docker for local development and database

## Running locally
Project requires a Postgres database to run. Make sure you have one running.

Build the project and run using
```
go build cmd/main/main.go
./main
```

## Example usage of endpoints
**Note for queries that requires a user: Since there is not an implemented user system, user id is passed via query params. In a real application this would be either a cookie or a token in the Authorization header.**

## List products
```GET /api/v1/products```

Example response:
```json
[
    {
        "id": 1,
        "name": "test",
        "price": 1000,
        "vat": 18
    },
    {
        "id": 2,
        "name": "test2",
        "price": 100,
        "vat": 8
    },
]
```

## Get a product
```GET /api/v1/products/{id}```

Example response:
```json
{
    "id": 1,
    "name": "test",
    "price": 1000,
    "vat": 18
}
```


## List Cart
```GET /api/v1/cart```

Example response:
```json
{
    "id": 3,
    "items": [
        {
            "id": 37, // ID of each item in cart
            "product": {
                "id": 2,
                "name": "test2", 
                "price": 100,
                "vat": 18
            },
            "sell_price": 118, // price + tax
            "tax": 18, // Actual calculated tax amount
            "quantity": 1,
            "total": 118 // Price * quantity
        }
    ],
    "summary": {
        "product_total": 118, // Original price of the cart including tax
        "final_price": 109.8, // Final price after discounts and tax changes
        "tax_total": 16.2, // New tax amount after discounts
        "discount_amount": 10 // How much actual discount is applied
    }
}
```

## Add to cart
```POST /api/v1/cart?user_id={id}```

Example response:
```json
{
    "id": 37,
    "product": {
        "id": 1,
        "name": "test",
        "price": 1000,
        "vat": 18
    },
    "cart_id": 3
}
```

## Remove from cart
```DELETE /api/v1/cart/item/{item_id}?user_id={id}```

Example response:
```json
{
    "id": "43",
    "status": "Success"
}
```

## Complete order 
```/api/v1/order/complete?user_id={}```

Example response:
```json
{
    "id": 19,
    "user_id": 3,
    "ordered_at": "2022-08-07T20:47:26.093433Z",
    "total": 2478
}
```

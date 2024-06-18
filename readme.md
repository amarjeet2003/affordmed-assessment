# Top Products HTTP Microservice

This microservice provides a RESTful API to retrieve top products from multiple e-commerce companies based on specified categories, price range, and sorting criteria. It aggregates data from various e-commerce APIs and presents a unified interface for users to compare products across different companies.

## Features

- **Fetch Top Products**: Retrieve the top N products within a specified category and price range.
- **Sorting**: Sort products by rating, price, company, or discount in ascending or descending order.
- **Pagination**: Support pagination when fetching more than 10 products.
- **Product Details**: Retrieve details of a specific product by its ID.

## Project Structure

The project follows the MVC (Model-View-Controller) pattern:

- **`controllers/`**: HTTP request handlers (controllers) for different endpoints.
- **`models/`**: Data structures used in the application.
- **`routes/`**: Registers HTTP routes and ties them to corresponding controllers.
- **`services/`**: Business logic to fetch and process data from e-commerce APIs.
- **`utils/`**: Utility functions and configurations.
- **`main.go`**: Entry point of the application where the HTTP server is initialized.

## Installation and Setup

1. **Clone the repository**:

   ```bash
   git clone <repository-url>
   cd top-products-microservice
   ```

2. **Install dependencies**:

   Ensure Go is installed. Then, install dependencies using:

   ```bash
   go mod tidy
   ```

3. **Run the server**:

   Start the server on `localhost:8080`:

   ```bash
   go run main.go
   ```

4. **Endpoints**:

   - **Fetch top products**: `GET /categories/{categoryname}/products`
     Example: `http://localhost:8080/categories/Laptop/products?top=10&minPrice=1&maxPrice=10000`
   
   - **Fetch product by ID**: `GET /categories/{categoryname}/products/{productid}`
     Example: `http://localhost:8080/categories/Laptop/products/1`

## Usage

1. **Fetch Top Products**:
   
   Make a GET request to the endpoint `/categories/{categoryname}/products` with query parameters:
   - `top`: Number of top products to fetch.
   - `minPrice`, `maxPrice`: Price range for products.
   - Additional optional parameters for sorting (`sort=rating|price|company|discount`) and pagination (`page`).

2. **Fetch Product by ID**:

   Make a GET request to `/categories/{categoryname}/products/{productid}` to retrieve details of a specific product identified by `productid`.

3. **Response Format**:

   Responses are returned in JSON format, structured as per the `Product` model defined in `models/product.go`.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

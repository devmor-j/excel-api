# Excel API

<!-- TODO: add api endpoints guide here -->

## Desicions Made

- Since http2 is not required, [fiber](https://github.com/gofiber/fiber) is chosen to run our http server but Go's http package could also serve well for this project provided that number of endpoints are kept small
- Since no mock data is provided, a seed generator was developed to create unique data and insert them into mongodb
- Student and seed student types are separated to allow furthur modification and type conversion during seed process (ex: human date string to golang Time type)
- Writing to spreadsheet performance can be faster using concurrency techniques

## Objective

The goal of this task is to assess the candidate's ability to build a RESTful service using Golang, interact with a MongoDB database, generate an Excel file, and return the generated file.

## Task Description

1. Build a RESTful API Service:

   - Endpoint: Create a REST API endpoint, such as GET /export-data.
   - Functionality:
     - Upon receiving a request to this endpoint, the service should fetch specific documents from a MongoDB collection.
     - Process the fetched data to generate an Excel file.

2. Data Fetching from MongoDB:

   - The service should connect to a MongoDB database.
   - Fetch documents based on specified criteria or parameters (e.g., date range, specific fields).
   - Ensure secure and efficient data retrieval, handling potential errors.

3. Excel File Generation:

   - Use a Golang library to generate an Excel file from the fetched data.
   - Structure the Excel file with appropriate headers and data formatting.

4. File Handling:

   - Return the created file.

## Submission Requirements

- Source code for the API service.
- A sample Excel file generated by the service.
- A brief explanation of the chosen architecture and design patterns.

_Note: If needed, you can add other features to the program to improve efficiency, security and speed._

## Development

Use docker or docker-compose to run a MongoDB instance.
Also Provide a username and password for the database inside `.env` file.

```sh
docker run --name mongodb -p 27017:27017 -d mongo:8.0.0-rc15-jammy
# or
docker-compose up -d
```

You can run server using `go run main.go` or `make`

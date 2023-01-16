# Product

### Repository for my GO lang + Lambda + API Gateway Project


## Development commands:
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
Add main file to main.zip
Create Lambda function with runtime as GO 1.x.
Upload main.zip as source code.
Point Main as the handler in runtime settings of Lambda function

GO to API Gateway in AWS
Create API --> HTTP API
Select Lambda function as the Integration
Configure routes with the necessary(or default) HTTP Methods, Resource Paths.



Test the functions using the following API Gateway routes:



### For Get request
GET https://uwdy3hxjl1.execute-api.us-east-1.amazonaws.com/product




###

POST https://uwdy3hxjl1.execute-api.us-east-1.amazonaws.com/product
Content-Type: application/json

{
    "firstName":"Narasimha",
    "lastName":"Konangi",
    "age":40
}


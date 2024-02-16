# Go Serverless API

Go Serverless lambda with DynamoDB  

## Pre-Requisites

Before you begin, ensure you have met the following requirements:

- [Go](https://golang.org/doc/install) (above v1.21.x)
- [Node.js and npm](https://nodejs.org/en/download/)
- [Serverless Framework](https://www.serverless.com/framework/docs/getting-started/)
- Serverless Offline, `npm install serverless-offline -g`

## Getting Started

To get a local copy up and running, follow these simple steps.

### Installation

1. Clone the repository
    ```sh
    git clone <url>
    cd go-Serverless-api
   ```
2. Install NPM dependencies
    ```sh
    npm install
    ```
3. Set Environment
    ```bash
    export ENV=local
    ```
4. Install Go dependencies and tidy up the module
    ```sh
    go mod tidy
    ```

### Running the application

To run the functions locally, you can use the following npm scripts:

- Start the Serverless application locally:

    ```sh
    npm start
    ```

  _Local execution will use the `serverless-offline` plugin to emulate the AWS Lambda environment and API Gateway locally.
  Since the go function will be build everytime you hit the local endpoint, the function will take some time to execute._


- Watch for changes and restart the service automatically (using nodemon):

    ```sh
    npm run watch
    ```

Note: Ensure you have Docker running if you are using the Docker-related commands.

## Deployment

### Serverless Configuration Changes

There are few changes to be made in the serverless file to deploy,
1. Update the timeout of functions to be `30`
2. Update the runtime to `provided.al2`

### Deploy

To deploy the application to your default AWS stage, use the following command:

```sh
serverless deploy --stage <your-stage>
```

## Configuration

All function configuration is given in the `serverless.yml` file. The following environment variables are required:
- ENV
- DYNAMO_TABLE

Refer .env file for examples


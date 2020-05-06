# jwt-go Microservice

## Description

jwt-go is a microservice that provides API for creating JWT Token and verifing the JWT token. 


## Requirements

  - Go 1.12 or higher
  
  - Do NOT disable [Go modules](https://github.com/golang/go/wiki/Modules) (`export GO111MODULE=on`)
  
## Installation

There are 2 methods explained in this document to install jwt-go:

1. Manual Installation Steps
2. Docker Based Installation

Please do **only one** method, do not do both methods, unless of course you know what you're doing.


### Method 1: Manual Installation Steps

1.  Clone the jwt-go into your go path and change directory to cloned repository.

    ```bash
    
    git clone https://@github.com/tiniyo/jwt-go.git
    
    cd jwt-go
    ```
    
2.  Build the jwt-go microservice

    ```bash
    
    go build -o jwt-go main.go
    
    ```
    
3. Set the configuration as per your requirement

    ```bash
    
    vim config/config.toml
    
    ```
    
 4. Now you are ready the run the jwt-go microservice
 
    ```bash
    ./jwt-go
    ```
    
### Method 2: Docker Based Installation

1.  Clone the jwt-go into your go path and change directory to cloned repository.

    ```bash
    
    git clone https://@github.com/tiniyo/jwt-go.git
    
    cd jwt-go
    ```
 
2.  Build the jwt-go microservice

    ```bash
    
    docker build -t jwt-go .
    
    ```

3. Now you are ready the run the jwt-go microservice
 
    ```bash
    
    docker run jwt-go
    
    ```

# Go-Classical-WebServer

## Overview

This Go program is a basic web server that serves static files and handles form submissions. 
It includes two endpoints: "/hello" for a simple "Hello!" message and "/form" for submitting a 
form with name and address fields.


## Prerequisites

Go (Golang) should be installed on your system.

## Usage

1. Clone or download the program's source code.
2. Open a terminal and navigate to the directory containing the program.
3. Compile the program using the following command: `go build main.go`
4. Run the compiled program: `./main`
5. The program will start a web server on port 8080.
6. Access the web application by opening your web browser and visiting http://localhost:8080.

## Endpoints
`/hello`
- Access http://localhost:8080/hello to display a "Hello!" message.
`/form`
- Access http://localhost:8080/form to submit a form with name and address fields.
- Upon submission, the program will display the submitted name and address.

## Static Files
- Place any static HTML files in the "static" directory, and the web server will serve them.

## Customization
- Customize the HTML files in the "static" directory to create your own web pages.
- Extend the program to add more endpoints and functionality as needed.

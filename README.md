# httpServer

This is a simple HTTP server written in Go that responds with the message "You are in server" to clients who visit the root URL path ("/").

## Usage

To use this server, you must have Go installed on your machine. Once you have Go installed, follow these steps:

- Save the code snippet to a file called `main.go`.
- Open a terminal window and navigate to the directory containing the `main.go` file.
- Run the command go run main.go to start the server.
- Open a web browser and navigate to `http://localhost/.`
  You should see the message "You are in server" displayed in the browser window.

## Implementation Details

The server is implemented using the Go standard library's net/http package. The http.HandleFunc() function is used to register a function that will handle incoming requests to the root URL path ("/"). This function simply writes the message "You are in server" to the response writer.

The http.ListenAndServe() function is used to start the server on port 80 and use the default http.ServeMux to handle incoming requests. Once started, the server will continue running and listening for incoming requests until it is explicitly stopped or terminated.

License
This code is released under the MIT license. See the LICENSE file for more information.

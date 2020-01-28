package books
​
// Book is a struct which represents a record of book
type Book struct {
	// ID is an unique value in the library
	ID       int `json:"id,omitempty"`
	Revision int `json:"revision,omitempty"`
	// ISBN number validation has to be performed
	// can contain hyphens. you can use an existing library
	// for the validation
	ISBN string `json:"isbn,omitempty"`
	// Can contain commas
	Title string `json:"title,omitempty"`
	// Can contain commas
	Author string `json:"author,omitempty"`
}
​
/*
	The above struct represents a single book and we need to manage
	a catalogue of books. The struct has the validation requirements written
	above each field and the json tags are also provided for the same. This
	will be the payload your server will receive`
*/
​
/*
	** Task 1 **
	Write a webserver in golang which implements endpoints for managing these records in a CSV file.
	Need to support 3 basic operations.
	HTTP GET /book - get all books
	HTTP GET /book/$ID - get specific book
	HTTP POST /book - insert new book
	HTTP PUT /book/$ID - update book details, can update specific detail or all details, except ID
	HTTP DELETE /book/$ID - delete a book
​
	The response time taken to process each of these requests must be logged to a file in the server
	For POST request simulation you can use curl or Postman`
*/
​
/*
	** Task 2 **
	Implement the same server as above using a database as the backend. Similar to above
	task log the time taken for processing the request.`
​
*/
​
/*
	** Task 3 **
	Implement a proxy server which accepts the above endpoints but uses paths to identify which server to call
	for example:
	/v1/book should get all books in the file
​
	/v2/book should get all the books from the database.
​
	This is the third server you will build. This server connects to the first 2 servers
	using a http client request and propogates the reponse to the final client`
*/
​
/*
	General Guidelines
	The project should be following the project structure in
	https://github.com/golang-standards/project-layout
​
	You can use https://github.com/gorilla/mux for the defining HTTP routes
​
	Use only the folders which make sense, cmd and pkg folders would be needed for sure
	And each of the server would be a package under the pkg folder.
	FYI: You are not being checked on how many folders you have :)`
​
*/
​
/*
	** EXTRA STEPS **
	Maintain the manipulation to file in memory and flush to file at specific time interval
​
	In case you finish the task early and want to do somthing more, then
	create docker files for running these 3 components and run them in a docker compose`
*/
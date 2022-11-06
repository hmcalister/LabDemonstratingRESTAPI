# Lab Demonstrating RESTful API
#### Hayden McAlister <haydenmcalister49@gmail.com>

Implementation of a RESTful API for lab demonstrating using Golang. Components are split into modules (which may more may not be best Golang practice?) and are described below. 

The intention of this project is to expose a web interface for a lab marking application. This requires state, hence, a REST api over a database. The end goal is to have an executable that spins up a webserver hosting both a web app frontend and REST api to interact with the database. Most of the CRUD operations will have implementations (see database section below).

## Running this project
To see this project in action, grab one of the releases (including an executable binary and the web app interface) and run it using `./rest_api`. If you need to build the application for a different platform, pull the source code and run `go build ./main`. Run the executable (use `-h` to get a summary of command line flags and options) and follow the link printed to start marking labs.

## Modules
### Models
This module just houses the model structs used throughout the project. Because we are using Gorm for object-relation mapping these structs also forms the schema of the database. The structs used are

- User: Intended for future implementation of authentication to the api/front end. Not currently used.
- Student: Representation of a single student. A student must exist before they can complete labs. Students can be created, deleted, and updated.
- Lab: Representation of a single lab. A lab must exist to be completed. Labs can be created, deleted, and updated.
- LabCompletion: Representation of a student completing a lab. Existence of an instance of this struct implies the referenced student completed the referenced lab. Can be created and deleted.

### Database
The core database functions. Uses SQLite and Gorm to achieve persistent state. The SQLite database file will be set using a command line argument to the executable (if multiple papers/labs are required, for example). While we are using Gorm to map the models into the database, we do not want to expose a raw connection from this module for security purposes. For this reason we have many wrapper functions that perform the functionalities for us (and are also a near one-to-one mapping of the API endpoints).

### API
Defines the API endpoints using GIN. API endpoints are near one-to-one mappings to exported functions from Database module, e.g. a `GET` request to `/student/:id` will return JSON of the student struct with the given ID (if it exists!).

This module will also eventually handle the front end serving of the webserver, as the GIN router is already defined here. Perhaps in future this responsibility will be shifted to a new module and the API can simply act on a separate port?

Something to look into in future is finding the correct HTTP response code for various events. Currently, most requests return 200 (OK) or an internal service error if something goes wrong. Some functions return a response code based on the most likely error (e.g. if a `PUT` request failed it's probably because the primary key is not unique).

### Main
A (hopefully!) very simple module to glue to above modules together. Handles getting command line arguments, setting up the database module, starting the API and any shutdown.


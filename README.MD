# Company Service
This service is use to company. It includes the following operations

- Add company
- List company
- Update company
  - Remove company

### Architecture

![alt text](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

##### Use Cases
The software in this layer contains application specific business rules. It encapsulates and implements all of the use cases of the system. 

##### Interface Adapters
The software in this layer is a set of adapters that convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the Database or the Web.

##### Infrastructure / Frameworks and Drivers.
The outermost layer is generally composed of frameworks and tools such as the Database, the Web Framework, etc. 

### Setup

Create database first and update the credentials in `config/config.go ` file.

After updating the file use the following command to download packages

`make deps`

NOTE : Go modules should be enabled.
### Run
Use one of the following commands to run the app

`make run`

or

`go run main.go`


## Documentation


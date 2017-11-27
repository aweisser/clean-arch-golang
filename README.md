# What's in here?
An example go application with a folder structure that reflects the Clean Architecture as proposed by Uncle Bob.

## Requirements

### Clean Architecture Layers
An application that follows the Clean Architecture by Uncle Bob consists out of several layers:

- Frameworks & Drivers - Describes anything thats outside the actual application process. Such as an explicit DB or restless webservice or a file system or something like that. 
- Interface Adapters - Connects the external stuff with the actual application. In this layer you should see the technical building blocks of the application such as 'Server', 'Client', 'Gateway', 'Proxy', 'Factory', 'Controller', 'Presenter' ...
- Application Business Rules - Holds the pure business rules or use cases that are only relevant for this special application. 
- Enterprise Business Rules - Holds the common business rules, use cases or domain model that can be reused in different applications.

### Components and Dependencies
Each of these layers contains several components which must follow a strict dependency rule. As an application developer I want to to be able to identify the boundaries of the components, so that I can easily meet the single responsibility principle and manage the dependency rules of the Clean Architecture.

### Screaming Architecture
As a new developer on the project I also want the folder structure and source file names scream there application purpose to me. It should be very easy to understand, that this is a Healthcare application or ERP system.

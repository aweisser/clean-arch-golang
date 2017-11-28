# What's in here?
An example go application with a folder structure that reflects the Clean Architecture as proposed by Uncle Bob.

## Requirements

### Use Clean Architecture Layers
An application that follows the Clean Architecture by Uncle Bob consists out of several layers:

- Frameworks & Drivers - Describes anything thats outside the actual application process. Such as an explicit DB or restless webservice or a file system or something like that. 
- Interface Adapters - Connects the external stuff with the actual application. In this layer you should see the technical building blocks of the application such as 'Server', 'Client', 'Gateway', 'Proxy', 'Factory', 'Controller', 'Presenter' ...
- Application Business Rules - Holds the pure business rules or use cases that are only relevant for this special application. 
- Enterprise Business Rules - Holds the common business rules, use cases or domain model that can be reused in different applications.

### Reflect Components, Boundaries and Dependencies
Each of these layers contains several components which must follow a strict dependency rule. As an application developer I want to to be able to identify the boundaries of the components, so that I can easily meet the single responsibility principle and manage the dependency rules of the Clean Architecture.

### Follow the Screaming Architecture
As a new developer on the project I also want the folder structure and source file names scream there application purpose to me. It should be very easy to understand, that this is a Healthcare application or ERP system.

### Use Go best practices
I also want to follow best practices of Go apps such as a ```cmd``` folder holding all executable ```main.go``` files.

## The Structure
The top level structure reflects the different layers of the Clean Architecture as well as folder ```cmd```. 
As a developer you can easily identifiy the different layers and their concerns. Additionally you can easily verify the dependency rules of the Clean Architecture.

(Note that the layer 'Frameworks & Drivers' is not reflected in the folder structure, because this layer only contains external processes that can be adapted by the application.)

The second level is all about the application purpose and follows the Screaming Architecture.
As a developer you can easily figure out what this application is all about and how to execute and access it.

```
.
+-- cmd                     (Contains variuous executable entry points to the application.)
|   +-- go-shopping-cobra   
|   +-- go-shopping-flag    
+-- domain                  (Enterprise Business Rules. This is the most inner layer of the Clean Architecture. Only incoming dependencies are allowed.)
|   +-- shop                (The domain seems to be a shop.)
+-- usecase                 (Application Business Rules. This is middle layer of the Clean Architecture. Only incoming dependencies from io and outgoing dependencies to domain are allowed)
|   +-- welcome             (The components within this layer may access other usecase components within this layer, as long as there are no cyclic dependencies)
|   |   +-- customer.go     (There seems to be a use case that welcomes customers.)
+-- io                      (Interface Adapters. This is the most outer layer and provides adapters to external processes or ports. Only outgoing dependencies to usecase and domain components are allowed.)
|   +-- cobracli            (Incoming adapter: a cobra based cli)
|   +-- flagcli             (Incoming adapter: a flag based cli)
|   +-- stdout              (Outgoing adapter: a stdout printer)
+-- factory.go              (A global factory assembles concrete components via dependency injection and provides valid application setups. To do this the factory may access every component without any dependency restriction.)
```

In this structure every go package can be seen as a single component with a clearly seperated concern, organized in concentric layers and follows the dependency rules of the Clean Architecture.
# Book Store

It is a store for ordering/selling books.

It is a customizable open source application for integration.


### objective
Create an API that allows requests and that in the end the customer can order and buy your book.

Bookstore Components
- Store
- Client
- Product
- Category
- Order
- Order Item

Application developed in Golang language.

Concept used for business development is DDD.

Pattern implementations to solve typical and repetitive problems.

Integrated SonarCloud in CI/CD pipeline is being used to review automated clean code checks.

## Getting Started
    Required
        - Golang

# Features
- [ ] client [TODO] CRUD
    - api, service, repository
        - [ ] create
        - [ ] update
        - [ ] delete
        - [ ] get
        - [ ] list
- [ ] product [TODO] CRUD
    - api, service, repository
        - [ ] create
        - [ ] update
        - [ ] delete
        - [ ] get
        - [ ] list
- [ ] category [TODO] CRUD
    - api, service, repository
        - [ ] create
        - [ ] update
        - [ ] delete
        - [ ] get
        - [ ] list
- [ ] order [TODO] CRUD
    - api, service, repository
        - [ ] create
        - [ ] update
        - [ ] delete
        - [ ] get
        - [ ] list
- [ ] order item [TODO] CRUD
    - api, service, repository
        - [ ] create
        - [ ] update
        - [ ] delete
        - [ ] get
        - [ ] list
- [ ] Persistent Storage
    - [ ] Repository Pattern
    - [ ] Postgres
- [ ] CI/CD
    - [ ] Github Actions
    - [ ] SonarCloud
    - [ ] Docker
- [ ] Tests
    - [ ] Unit Tests
    - [ ] Integration Tests
    - [ ] E2E Tests
- [ ] Observability
    - [ ] Prometheus
    - [ ] Grafana
    - [ ] Jaeger
    - [ ] Zipkin
    - [ ] Loki
    - [ ] Tracing
    - [ ] Metrics
    - [ ] Logs
- [ ] Security
    - [ ] Authentication
    - [ ] Authorization
    - [ ] Encryption
    - [ ] Hashing
    - [ ] JWT
- [ ] API Documentation
    - [ ] Swagger
- [ ] Cache
    - [ ] Redis

## Project Structure

### ***Domain*** is the core of the application, it contains the business rules and the entities that represent the business.

- aggregate root - root of the aggregate
- entity - object with identity
- factory - object that creates other objects
- repository - object that manages the persistence of entities
- service - object that implements business rules
- value object - object that does not have identity

### ***Application*** is the layer that connects the domain with the external world, it is responsible for receiving requests and returning responses.

- api - is the layer that receives requests and returns responses
- service - is the layer that implements the business rules

### ***Infrastructure*** is the layer that connects the application with the external world, it is responsible for receiving requests and returning responses.

- cache - is a cache that stores data in a structured format
- database - is a database that stores data in a structured format
- logger - is a log aggregator that collects logs from different sources and stores them in a central location
- metrics - is a monitoring tool that collects metrics from different sources and stores them in a central location
- tracing - is a tool that collects traces from different sources and stores them in a central location



## DDD Domain Drive Design
This project uses a lot of the ideas introduced by Eric Evans in his book [Domain Driven Design](https://www.domainlanguage.com/), which is a great book that I recommend to anyone who wants to learn more about the subject.

### ***Strategic Design***

![Ubiquitous Language](docs/ddd/img/ddd-Ubiquitous%20Language.drawio.png)
![Bounded Context](docs/ddd/img/ddd-Bounded%20Context.drawio.png)
![Context Map](docs/ddd/img/ddd-Context%20Map.drawio.png)

## Contributing

For major changes, open a problem first to discuss what you would like to change.

### Codification

> Attention, do not write, change or delete anything in the code from the main branch (default), for any new functionality you must create a new branch 

- To create a new branch run the command `git checkout -b <branch-name>`
    - To switch from one branch to another existing in your local repository run `git checkout <branch-name>` without `-b`

- Once you have finished coding on the new branch, push it to GitHub with the command `git push -u origin <branch-name>`
    - After this step, open a pull-request for the main branch and wait for the Admin to analyze and merge the pull-request




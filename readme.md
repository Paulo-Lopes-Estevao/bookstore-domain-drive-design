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

## DDD Domain Drive Design

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




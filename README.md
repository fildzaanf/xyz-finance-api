# xyz-finance-api

## 🚀 Tools and Technologies

* Go Programming Language
* Echo Go Framework
* GORM for Object Relational Mapping
* PostgreSQL for Relational Database
* JSON Web Token (JWT) for Authentication
* Docker for Containerization
* CI/CD Pipeline using GitHub Actions
* Amazon Web Services (AWS)
  * Amazon Simple Storage Service (Amazon S3)
* Postman for API Testing
* Webhook for Midtrans Payment Gateway 

## 🏛️ System Design and Architecture

* Clean Architecture
* Hexagonal Architecture
* Domain-Driven Design (DDD)
* Command Query Responsibility Segregation (CQRS)
* Microservices Architecture
* REST API

## ✨ Features

#### User Management

| Feature          | Description                                                             |
| ---------------- | ----------------------------------------------------------------------- |
| Authentication   | Handles user registration, login, and retrieval of profile information  |
| Profile          | Allows users to update or retrieve their personal details               |

#### Transaction Management

| Feature               | Description                                                         |
| --------------------- | ------------------------------------------------------------------- |
| Transaction Creation  | Enables users to create new financial transactions                  |
| Transaction Retrieval | Provides access to all transactions or specific transactions by ID  |

#### Payment Management

| Feature             | Description                                                          |
| ------------------- | -------------------------------------------------------------------- |
| Payment Creation    | Allows users to create new payments                                  |
| Payment Retrieval   | Enables retrieval of all payments or details of a specific payment   |
| Payment Integration | Supports real-time payment updates via Midtrans webhook integration  |

#### Loan Management

| Feature                | Description                                           |
| ---------------------- | ----------------------------------------------------- |
| Loan Creation          | Enables users to create new loans                     |
| Loan Retrieval         | Provides access to all loans or specific loans by ID  |
| Loan Status Management | Allows updating the status of loans as they progress  |

#### Installment Management

| Feature                       | Description                                                         |
| ----------------------------- | ------------------------------------------------------------------- |
| Installment Creation          | Enables creation of new installments for loans                      |
| Installment Retrieval         | Provides access to all installments or specific installments by ID  |
| Installment Status Management | Allows updating the status of installments to track payments        |


## 📡 API Endpoints

#### Users

| Method | Endpoint        | Description                 |
| ------ | --------------- | --------------------------- |
| POST   | /users/register | Register a new user         |
| POST   | /users/login    | Login user                  |
| GET    | /users/:user_id | Retrieve user profile by ID |

#### Transactions

| Method | Endpoint          | Description                   |
| ------ | ----------------- | ----------------------------- |
| POST   | /transactions     | Create a new transaction      |
| GET    | /transactions     | Retrieve all transactions     |
| GET    | /transactions/:id | Get transaction details by ID |

#### Payments

| Method | Endpoint                   | Description                                  |
| ------ | -------------------------- | -------------------------------------------- |
| POST   | /payments                  | Create a new payment                         |
| POST   | /payments/midtrans/webhook | Receive Midtrans webhook for payment updates |
| GET    | /payments                  | Retrieve all payments                        |
| GET    | /payments/:id              | Retrieve payment by ID                       |

#### Installments

| Method | Endpoint          | Description                     |
| ------ | ----------------- | ------------------------------- |
| GET    | /installments     | Retrieve all installments       |
| GET    | /installments/:id | Get installment details by ID   |
| POST   | /installments     | Create a new installment        |
| PUT    | /installments/:id | Update installment status by ID |

#### Loans

| Method | Endpoint   | Description              |
| ------ | ---------- | ------------------------ |
| GET    | /loans     | Retrieve all loans       |
| GET    | /loans/:id | Get loan details by ID   |
| POST   | /loans     | Create a new loan        |
| PUT    | /loans/:id | Update loan status by ID |


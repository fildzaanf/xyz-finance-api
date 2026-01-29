# XPay : Finance Solutions

## 📝 Project Overview
XPay is a digital financing platform for electronics, vehicles, and household appliances, enabling users to apply for loans, make transactions, and manage installment payments

## 🎯 Problem Statement & Solution

#### Problem Statement
Many users face challenges in managing purchases of electronics, vehicles, and household appliances due to limited access to flexible financing options, difficulty tracking installment payments, and the complexity of managing multiple transactions manually. These issues can lead to missed payments, financial confusion, and reduced access to essential goods.

#### Solution
XPay provides a digital financing platform that simplifies loan applications, transaction management, and installment tracking. The platform allows users to:

- Apply for loans quickly for electronics, vehicles, and household appliances
- Create and track financial transactions in real time
- Manage installment payments efficiently with clear status tracking
- Integrate with payment gateways (e.g., Midtrans) for secure and automated payments

By centralizing all financial activities in a single, user-friendly platform, XPay helps users access financing easily, maintain accurate payment records, and improve financial management.

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


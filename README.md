
# Bank Account Web Service

A simple microservice that simulates a "Bank Account" with endpoints to check balance, deposit funds, and withdraw funds. The service has restrictions on the maximum deposit and withdrawal amounts, as well as transaction limits per day. 

This application is built using **GoLang**, **Chi** router for handling HTTP requests, and **GORM** with **PostgreSQL** for managing the database.

## Features
- **Balance**: Retrieve the current balance of the account.
- **Deposit**: Deposit money into the account with restrictions.
- **Withdraw**: Withdraw money from the account with restrictions.
  
### Endpoints:
1. **GET /balance** – Returns the current account balance.
2. **POST /deposit** – Deposits money into the account.
3. **POST /withdraw** – Withdraws money from the account.

## Approach & Trade-offs

### 1. **Framework Choice**
   - **GoLang**: Chosen for its simplicity, concurrency, and performance. Go provides efficient handling of high-concurrency workloads, making it suitable for small, high-performance services.
   - **Chi Router**: The **Chi** framework was chosen for simplicity and lightweight routing. It is an easy-to-use alternative to heavier frameworks like Gin, while still being fast and flexible.
   
   **Trade-off**: Using a lightweight router like Chi reduces the overhead of a heavy framework (like Gin ), but also requires more manual setup of middleware, such as error handling and request validation.

### 2. **Database Choice**
   - **PostgreSQL**: Chosen for its reliability, and support for complex queries. It is ideal for managing transactional systems like a bank account service.

### 3. **Transaction Restrictions**
   - **Deposit**: Maximum of **$150,000 per day**, **$40,000 per transaction**, and **4 deposits per day**.
   - **Withdrawal**: Maximum of **$50,000 per day**, **$20,000 per transaction**, and **3 withdrawals per day**.

   These restrictions are based on the instuctions given in the assessment

### 4. **No Authentication (for simplicity)**
   - The service is designed to be open for testing purposes, meaning that no authentication is required to access the endpoints. Adding auth was optional for this assessment

   **Trade-off**: This makes the service simple to use for testing but is insecure for production use. In production, you'd likely want to add JWT or another form of authentication to secure the endpoints.

## How to Run the Application

### Prerequisites
- **GoLang** (version 1.18 or higher)
- **PostgreSQL** (or a Dockerized PostgreSQL instance)
- **Docker** (optional, if you prefer running the app in containers)

### 1. Clone the repository

```sh
git clone https://github.com/Rebeccah5/bank-service.git
cd bank-service
```

### 2. Set up environment variables

Create a `.env` file in the root of the project to store your database credentials and other configuration settings. Check `.env.example` file

### 3. Install Go dependencies

Make sure Go modules are initialized and dependencies are installed:

```sh
go mod tidy
```

### 4. Run the application locally

To run the application locally:

```sh
go run cmd/main.go
```

The application will start on the default port **8080**. You can visit the following endpoints:

- `http://localhost:8080/balance`
- `http://localhost:8080/deposit`
- `http://localhost:8080/withdraw`

# Internship-CRUD microservice

## Description

This is a CRUD microservice for the Intership project. This microservice is able to create, read and update users, internships and all related data.

### Technology stack

- The microservice is written in **GO**.
- This service communicates through a **RabbitMQ** message broker using **Remote Procedure Calls (RPC)**.
- To use **RabbitMQ**, we use the **amqp** library for **GO**.
- To communicate with a **Postgres** database we use the **pgx** library.

### Deployment

This service is deployed and available through the **API gateway**.<br>
We use **Coolify** to automatically deploy the microservice inside a **Docker container**.

## Installation

Here is an installation guide on how to run this project locally.

### Requirements

- **Go 1.19.3** installed on your machine or higher.
- Access to a **RabbitMQ message bus**.
- Access to a **PostgreSQL** database.

### Installation

1. Clone the repository to your local machine. This can be done through the following command:

```bash

git clone git@git.ti.howest.be:TI/2022-2023/s5/trending-topics-sd/students/mars05/internship-microservice.git

```

2. Navigate to the root of the project and open a command line prompt.

3. install the dependencies by running the following command:

```bash

go mod download

```

4. Create a **.env** file in the root of the project. This file should contain the variables that are used in the **.env.example** file.

5. Run the following command to start the microservice:

```bash

go run src/main.go

```

### Usage

This microservice only communicates over **RabbitMQ** message queues with the use of RPC calls.

The format of the messages is as follows:

```json

{
    "action": "get | create | update | customTag",
    "data": { ... }
}

```

To easily communicate with this microservice, we use an [API gateway](https://git.ti.howest.be/TI/2022-2023/s5/trending-topics-sd/students/mars05/api-gateway) that will publish an RPC call to the correct microservice.











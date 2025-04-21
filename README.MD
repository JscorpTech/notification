# Notification Service

A microservice for handling and delivering notifications through various channels like SMS and email using RabbitMQ as a message broker.

## Overview

This notification service is designed as a standalone microservice that consumes notification requests from a RabbitMQ queue and routes them to the appropriate notification provider based on the notification type. Currently, it supports SMS and email notifications.

## Features

- Message consumption from RabbitMQ
- Support for multiple notification channels (SMS, email)
- Extensible architecture for adding new notification types
- Asynchronous notification handling

## Architecture

The notification service follows a clean architecture approach:

- **Domain Layer**: Contains core business logic and port interfaces
- **Infrastructure Layer**: Implements the ports with concrete adapters
- **RabbitMQ**: Used as a message broker for consuming notification requests

## Installation

### Prerequisites

- Go 1.x+
- RabbitMQ server

### Setup

1. Clone the repository:
```bash
git clone https://github.com/JscorpTech/notification.git
cd notification
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o notification-service ./cmd/main.go
```

## Configuration

Configure your RabbitMQ connection and other settings in the appropriate configuration files.

## Usage

### Running the service

```bash
./notification-service
```

This will start the notification consumer that listens for incoming notification requests.

### Sending a notification

Notifications should be published to the RabbitMQ exchange with the following JSON format:

```json
{
  "type": "email",
  "message": "Hello, this is a test notification.",
  "to": ["user@example.com"]
}
```

Python example
```python
from kombu import Connection, Exchange, Producer

# RabbitMQ ulanishi
rabbit_url = 'amqp://guest:guest@127.0.0.1:5672/'
connection = Connection(rabbit_url)
channel = connection.channel()

exchange = Exchange('notification', type='direct')

# Producer yaratish
producer = Producer(channel, exchange=exchange, routing_key="notification")

# Xabar yuborish
message = {'type': 'sms', 'message': "classcom.uz sayti va mobil ilovasiga ro'yxatdan o'tishingingiz uchun tasdiqlash kodi: 1234", "to": ["+998888112309", "+998943990509"]}
producer.publish(message)

print("Message sent to all workers!")
```

Available notification types:
- `email`: For email notifications
- `sms`: For SMS notifications

## Project Structure

```
notification/
├── cmd/
│   └── main.go           # Entry point
├── internal/
│   ├── domain/
│   │   └── ports.go      # Interfaces
│   ├── notifier/
│   │   ├── email.go      # Email notification implementation
│   │   └── sms.go        # SMS notification implementation
│   ├── rabbitmq/
│   │   └── connection.go # RabbitMQ connection handling
│   └── consumer/
│       └── consumer.go   # Implementation of the notification consumer
└── README.md
```

## Contributing

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin feature/my-new-feature`
5. Submit a pull request

## License

[Add your license here]

## Contact

JscorpTech - [GitHub](https://github.com/JscorpTech)

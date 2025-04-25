# Notification Service

A flexible notification service that supports multiple message brokers (Redis, RabbitMQ) and notification channels (SMS, Email). This service is designed to handle asynchronous notification delivery in your applications.

## Features

- **Multiple Brokers**: Support for both Redis and RabbitMQ as message brokers
- **Multiple Notification Channels**: SMS and Email notification support
- **Containerized**: Ready to deploy with Docker
- **Extensible Architecture**: Easy to add new notification channels or message brokers

## Architecture

The service follows a clean architecture pattern with the following components:

- **Broker**: Handles message subscription from different sources (Redis/RabbitMQ)
- **Notifier**: Implements different notification channels (SMS/Email)
- **Services**: Contains the business logic for each notification channel
- **Domain**: Defines interfaces and data models

## Prerequisites

- Go 1.24 or higher
- Redis (for Redis broker)
- RabbitMQ (for RabbitMQ broker)
- Docker (optional, for containerized deployment)

## Configuration

Copy the provided `.env.example` to `.env` and update with your configuration:

```bash
cp .env.example .env
```

### Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| BROKER | Message broker to use (redis or rabbitmq) | redis |
| TOPIC | Topic/queue name for notifications | notification |
| REDIS_ADDRESS | Redis server address | 127.0.0.1:6379 |
| REDIS_PASSWORD | Redis password (if any) | |
| REDIS_DB | Redis database number | 0 |
| RABBITMQ_URL | RabbitMQ connection URL | amqp://guest:guest@localhost:5672/ |
| ESKIZ_DOMAIN | Eskiz SMS API domain | https://notify.eskiz.uz/api |
| ESKIZ_USER | Eskiz SMS API username | admin@example.com |
| ESKIZ_PASSWORD | Eskiz SMS API password | password |
| ESKIZ_FROM | Eskiz SMS sender ID | 4546 |
| MAIL_DOMAIN | SMTP server domain | smtp.gmail.com |
| MAIL_USER | SMTP username | notification@example.com |
| MAIL_PASSWORD | SMTP password | yourpassword |
| MAIL_PORT | SMTP port | 587 |

## Installation

### Local Development

1. Clone the repository
   ```bash
   git clone https://github.com/JscorpTech/notification.git
   cd notification
   ```

2. Install dependencies
   ```bash
   go mod download
   ```

3. Build and run the application
   ```bash
   go build -o notification ./cmd/main.go
   ./notification
   ```

### Docker Deployment

Build and run using Docker:

```bash
docker build -t notification-service .
docker run -p 8080:8080 --env-file .env notification-service
```

Or using Docker Compose:

```bash
docker-compose up -d
```

## Usage

### Message Format

The service expects messages in the following JSON format:

```json
{
  "type": "email",
  "message": "Subject: Welcome\r\n\r\nHello, welcome to our service.",
  "to": ["user@example.com"]
}
```

For SMS notifications:

```json
{
  "type": "sms",
  "message": "Your verification code is 1234",
  "to": ["+998901234567"]
}
```

### Sending Messages

#### Using Redis

```python
import redis
import json

r = redis.StrictRedis(host='127.0.0.1', port=6379, db=0)

message = {
    'type': 'email',
    'message': "Subject: Welcome\r\n\r\nWelcome to our service!",
    'to': ["user@example.com"]
}

r.rpush('notification', json.dumps(message))
```

#### Using RabbitMQ

```python
from kombu import Connection, Exchange, Producer

rabbit_url = 'amqp://guest:guest@127.0.0.1:5672/'
connection = Connection(rabbit_url)
channel = connection.channel()

exchange = Exchange('notification', type='direct')
producer = Producer(channel, exchange=exchange, routing_key="notification")

message = {
    'type': 'sms',
    'message': "Your verification code is 1234",
    'to': ["+998901234567"]
}
producer.publish(message)
```

## Adding New Notification Channels

1. Create a new notifier implementation in `internal/notifier/`
2. Implement the `domain.NotifierPort` interface
3. Add the new notifier type to the `Handler` function in `internal/consumer/notification.go`

## License

MIT License - See [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-feature`)
5. Create a new Pull Request
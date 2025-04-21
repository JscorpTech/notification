# rabbitmq.py

from kombu import Connection, Exchange, Producer

rabbit_url = 'amqp://guest:guest@rabbitmq:5672/'
_connection = None
_channel = None
_exchange = Exchange('notification', type='direct')
_producer = None


def get_connection():
    global _connection
    if _connection is None or not _connection.connected:
        _connection = Connection(rabbit_url)
        _connection.ensure_connection(max_retries=3)
    return _connection


def get_producer():
    global _producer, _channel
    if _producer is None:
        conn = get_connection()
        _channel = conn.channel()
        _producer = Producer(_channel, exchange=_exchange, routing_key="notification")
    return _producer


def send_notification(message: dict):
    producer = get_producer()
    producer.publish(message, serializer='json')

# rabbitmq.py

from kombu import Exchange, Producer
from .rabbitmq import get_connection

_channel = None
_exchange = Exchange('notification', type='direct')
_producer = None

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

# rabbitmq.py

from kombu import Connection

rabbit_url = 'amqp://guest:guest@rabbitmq:5672/'
_connection = None


def get_connection():
    global _connection
    if _connection is None or not _connection.connected:
        _connection = Connection(rabbit_url)
        _connection.ensure_connection(max_retries=3)
    return _connection


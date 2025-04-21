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

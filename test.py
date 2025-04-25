# from kombu import Connection, Exchange, Producer

# # RabbitMQ ulanishi
# rabbit_url = 'amqp://guest:guest@127.0.0.1:5672/'
# connection = Connection(rabbit_url)
# channel = connection.channel()

# exchange = Exchange('notification', type='direct')

# # Producer yaratish
# producer = Producer(channel, exchange=exchange, routing_key="notification")

# # Xabar yuborish
# message = {'type': 'email', 'message': "Subject: test\r\n\r\nclasscom.uz sayti va mobil ilovasiga ro'yxatdan o'tishingingiz uchun tasdiqlash kodi: 1234", "to": ["JscorpTech@gmail.com", "admin@jscorp.uz"]}
# producer.publish(message)

# print("Message sent to all workers!")


import redis
import json

# Redis ulanishi
r = redis.StrictRedis(host='127.0.0.1', port=6379, db=0)
# Subject: tasdiqlash ko'di\r\n\r\n
# Xabar tayyorlash
message = {
    'type': 'sms',
    'message': "Assalomu alaykum samandar sizni https://classcom.uz o’qituvchining virtual kаbinetida muallif sifatida tasdiqlanganingiz bilan tabriklaymiz!!!",
    'to': ["+998888112309"]
}

# Xabarni JSON formatga o‘tkazib, Redis listga push qilish
r.rpush('notification', json.dumps(message))

print("Message pushed to Redis list!")

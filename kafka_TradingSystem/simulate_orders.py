import time
import random
from kafka import KafkaProducer
import json

def generate_order_commands():
    order_types = ['buy', 'sell']
    symbols = ['AAPL', 'GOOGL', 'MSFT', 'AMZN', 'TSLA']
    while True:
        order_id = random.randint(1000, 9999)
        user_id = random.randint(500, 999)
        symbol = random.choice(symbols)
        quantity = random.randint(1, 100)
        order_type = random.choice(order_types)
        timestamp = int(time.time())
        yield {
            'order_id': order_id,
            'user_id': user_id,
            'symbol': symbol,
            'quantity': quantity,
            'order_type': order_type,
            'timestamp': timestamp
        }

def main():
    producer = KafkaProducer(
        bootstrap_servers='localhost:9092',
        value_serializer=lambda v: json.dumps(v).encode('utf-8')
    )
    for order_command in generate_order_commands():
        producer.send('order-commands', order_command)
        print(f"Sent order command: {order_command}")
        time.sleep(random.uniform(0.5, 1.5))  # Simulate variability in order submission rate

if __name__ == '__main__':
    main()

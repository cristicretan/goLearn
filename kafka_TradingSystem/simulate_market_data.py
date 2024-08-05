import time
import random
from kafka import KafkaProducer
import json

def generate_market_data():
    symbols = ['AAPL', 'GOOGL', 'MSFT', 'AMZN', 'TSLA']
    while True:
        symbol = random.choice(symbols)
        price = round(random.uniform(100, 1500), 2)
        timestamp = int(time.time())
        yield {
            'symbol': symbol,
            'price': price,
            'timestamp': timestamp
        }

def main():
    producer = KafkaProducer(
        bootstrap_servers='localhost:9092',
        value_serializer=lambda v: json.dumps(v).encode('utf-8')
    )
    for market_data in generate_market_data():
        producer.send('market-data', market_data)
        print(f"Sent market data: {market_data}")
        time.sleep(1)  # Simulate time delay between data points

if __name__ == '__main__':
    main()

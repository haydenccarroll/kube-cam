import time
import pika

def main():
    print("Hello! Container is running...")

    # Establish a connection to the RabbitMQ server
    connection = pika.BlockingConnection(pika.ConnectionParameters('rabbitmq-service.camera.svc.cluster.local'))

    # Create a channel on the connection
    channel = connection.channel()

    # Declare a queue on the channel (RabbitMQ will create the queue if it doesn't exist)
    channel.queue_declare(queue='my_queue')

    # Publish a message to the queue
    channel.basic_publish(exchange='', routing_key='my_queue', body='Hello, World!')

    print(" [x] Sent 'Hello, World!'")
    while True:
        time.sleep(1)

if __name__ == "__main__":
    main()
FROM python:3.9-alpine

WORKDIR /app

# used to interact with RabbitMQ
RUN pip install pika

COPY . .

# so python doesn't buffer prints to stdout/stderr
ENV PYTHONUNBUFFERED=1

ENTRYPOINT ["python3", "main.py"]
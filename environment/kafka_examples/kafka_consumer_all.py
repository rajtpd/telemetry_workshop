#!/usr/bin/python2

from kafka import KafkaConsumer

# Create a new Kafka consumer and ask for the latest set of data being stored in the Kafka bus.
# Note that you can use auto_offset_reset='earliest' for ask for all the available data.

#consumer = KafkaConsumer(group_id='different',bootstrap_servers='192.168.10.3:9093', auto_offset_reset='earliest')
consumer = KafkaConsumer(group_id='different',bootstrap_servers='192.168.10.3:9093', auto_offset_reset='latest')

# Subscribe to the telemetry topic, we are exporting from Pipeline
consumer.subscribe(['telemetry'])

# Print every message received from the telemetry topic
for message in consumer:
    print message

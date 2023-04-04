# AngelOwl Loyalties

- The `docker-compose.yml` file can be used to compose up the backend of the application, 
- The `test.py` file can be used to test out user interactions, such as
    - Creating a couple of exclusions and campaigns while the rewarder is running. This tests that the rewarder is appropriately watching for changes on etcd.
    - Triggers the lambda to parse a sample csv in our s3 bucket and push it to our production Kafka
---
The 4 main microservices are 
- Campaignex
    - CRUD operations for merchant managers to create and manage Campaigns and Exclusions. This service is responsible to write these into etcd, as well as its own PostgreSQL database as a layer of redundancy, or for auditing needs.
- Rewarder
    - This service isn’t a REST API, but acts as an orchestrator to process transactions it consumes from Kafka and to determine the appropriate reward a transaction will receive. It is also responsible for writing these processed transactions into the Cassandra database.
- Informer
    - This service exposes an API that the frontend web client calls to retrieve the rewards history pertaining to a requested card from the Cassandra database.
- Profiler
    - This service exposes REST API endpoints to perform CRUD operations on User and Customer information. It allows a customer to obtain information on their user profile as well as the cards in their ownership.

Ingestor is a lambda that ingests a CSV file from S3 and parses rows and acts as a producer to into Kafka.

## Images

### Kafka UI - Consumer Group `reward`

This screenshot shows the number of messages the cluster of rewader service pods is yet to consume.

![Kafka UI Consumer Screen](docs/kafkaui_consumer.png)

### Kafka UI - Topic `transaction`

This screenshot shows the number of messages present in the topic of concern: `transaction`

![Kafka UI Topic](docs/kafkaui_topic.png)

### Web App UI - Admin View

This screenshot shows the main page of UI where admin can view prevailing campaigns and exclusions, and add their own

![UI Rewards](docs/admin.png)

### Web App UI - Upload files

This screenshot shows a part of the admin portion of the UI where admins can drop bulk transaction CSV, or add users via a CSV.

![UI Rewards](docs/upload.png)
# AngelOwl Loyalties

<!-- Doesnt work for private repositories, uncomment when public -->
<!-- <table style="width: 100%; border: none;" cellspacing="0" cellpadding="0" border="0">
  <tr>
    <td>![Authorizer](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/actions/workflows/ci_authorizer.yml/badge.svg)</td>
    <td>![Campaignex](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/actions/workflows/ci_campaignex.yml/badge.svg)</td>
    <td>![Emailer](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/actions/workflows/ci_emailer.yml/badge.svg)</td>
    <td>![Informer](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/actions/workflows/ci_informer.yml/badge.svg)</td>
    <td>![Ingestor](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/actions/workflows/ci_ingestor.yml/badge.svg)</td>
    <td>![Passworder](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/actions/workflows/ci_passworder.yml/badge.svg)</td>
    <td>![User Ingestor](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/actions/workflows/ci_pet-rock.yml/badge.svg)</td>
    <td>![Profiler](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/actions/workflows/ci_profiler.yml/badge.svg)</td>
    <td>![Publish Single](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/actions/workflows/ci_publish-single.yml/badge.svg)</td>
    <td>![Rewarder](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/actions/workflows/ci_rewarder.yml/badge.svg)</td>
  </tr>
</table> -->

## Team Members

[Aw Khai Loong](https://github.com/awwkl)

[Goh Wen Liang](https://github.com/wenlianggg)

[Lam Xi Kai Justin](https://github.com/iPhantasmic)

[Lye Jian Yi](https://github.com/lyejy)

[Omer Wai Yan Oo](https://github.com/omerwyo)

[Ong Chi Kiat Nicholas](https://github.com/oversparkling)

[Owyong Jian Wei](https://github.com/alvinowyong)

## Directory Structure

As we have followed a monorepo strategy, you'll find one folder for each service. The following folders are not directly related to a service but are crucial to our deployments

- [`.chart/`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/.chart)
    - These consist of Helm charts that are used in our Github Action for deployment. This folder specifies yamls that describe our Kubernetes resources, secrets and others so we can declare how we want our cluster to look like at the end of a deployment.
- [`.github/workflows/`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/.github/workflows)
    - This file consists of 2 reusable workflows, `workflow_container.yml` and `workflow_lambda.yml`, and `ci_` files that reference the worflows. These CI files help automate our CI/CD pipeline, from linting to auto-deploying to production.
    - Some miscellaneous files include a yaml for sending notifications to our Telegram group when a PR or a code review is submitted
- [`.infra/`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/.infra)
    - Consists of all the Terraform files that declare the goal state of our deployments. One could use the Terraform commands `init`, `plan`, `apply`/`destroy` to make changes to our deployed infrastructure on AWS.
- [`docs/`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/docs)
    - This folder stores some images that you can see at the bottom of this README

The remaining folders are either containerised services, or containerised lambdas that we will explore in the next section.

## Microservices Architecture

Here we cover the functionality each service/lambda serves and how they could be deployed.

Every subdirectory mentioned can be deployed locally as a container, and one could do this using a docker-compose file at the root of the repository. The Dockerfile and the corresponding docker-compose at each subdirectory provides a good reference on how this can be done.


We can first describe some of the common environment variables we utilise in the app

|Environment Variable|Description|
|:---:|:---:|
|DB_CONN_STRING|This is generally used with regards to our RDS Postgres databases, specifying the host, port, username and password for connection when a service starts up. In some cases, it could be the URL of our Cassandra database, also used for connection purposes|
|ETCD_ENDPOINTS|This is used to specify the host URL of ETCD, whether it is deployed or locally.|
|BROKER_HOST|Specifies the URL of the Kafka broker|
|TOPIC|Kafka topic to write into/read from|
|DB_KEYSPACE|Cassandra DB keyspace|
|PROFILER_ENDPOINT|This is specified in the user-aggregator (pet-rock) lambda, so as to make a POST request to the profiler service to create a batch of users from a CSV file of new user information|
|INFORMER_ENDPOINT|This is specified in the emailer lambda, so as to make a GET request to the informer endpoint to retrieve rewards of the day|
|AWS_ACCESS_KEY_ID|This is added mainly in the Lambdas, due to the various roles they perform (put files into S3, send emails). In production lambdas, these environment variables are added by default, so these are mainly added for local testing purposes.|
|AWS_SECRET_ACCESS_KEY|Similar reasons as the above|


### Containerised Services
---
For all non-lambda services, a sample `.env` file can be found at `/config/dev.env`. A `.env` file with your own environment variables could be used to pass in relevant values into the docker-compose file.
- [`campaignex`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/campaignex)
    - This service uses the Golang REST framework Gin, exposing CRUD operations for merchant managers to create and manage Campaigns and Exclusions. This service is responsible to write these into etcd, as well as its own PostgreSQL database as a layer of redundancy, or for auditing needs.
- [`rewarder`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/rewarder)
    - This service isn’t a REST API, but acts as an orchestrator to process transactions it consumes from Kafka and to determine the appropriate reward a transaction will receive. It is also responsible for writing these processed transactions into the Cassandra database.
- [`informer`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/informer)
    - This service uses the Golang REST framework Gin, exposing Read operations that the frontend web client calls to retrieve the rewards history pertaining to a requested card from the Cassandra database.
- [`profiler`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/profiler)
    - This service uses the Golang REST framework Gin, exposing CRUD operations on User and Customer information. The frontend and some other services call on this service to obtain user data.

### Lambdas
---
For all lambda services, an `.env.example` is provided to the same effect. For lambda services, the Dockerfiles and the corresponding docker-compose.yml follow AWS' guide on deploying lambdas locally so that they replicate the behaviour of production Lambdas, using AWS Lambda Runtime Interface Emulator (RIE). The AWS Lambda RIE is a proxy for the Lambda Runtime API that allows you to locally test your Lambda function packaged as a container image.

The lambdas are 
- [`authorizer`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/authorizer)
    - A serverless function for authorization needs, it is linked directly to our API gateway to authorise relevant requests
- [`emailer`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/emailer)
    - This lambda runs as a job scheduled from AWS EventBridge once every night to send emails to users who have earned rewards amounting to more than $0, also informing them of the transactions that led to the corresponding rewards.
-  [`ingestor`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/ingestor) (Batch transaction upload)
    - This lambda takes in the file of batch transactions from S3, triggered when the pre-signed URL directs the user's file into our S3 bucket. It validates data (i.e valid MCC, valid length CardPan); and sends them to Kafka. This service is highly performant and crucial to the functionality of the application.
- [`pet-rock`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/pet-rock) (Batch user creation)
    - This lambda, similar to ingestor ingests a csv file of new users that an admin uploaded. It processes them line by line and creates a user on the backend for them.
- [`passworder`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/passworder)
    - This lambda is called directly after the user-aggregator, sending emails to the newly created users, sending them their first-time password so they can login. They will then be prompted by the UI to change their password after first login.
- [`publish-single`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/publish-single) (Transaction API)
    - Similar functionality and code to Ingestor, except instead of a file, this lambda takes in a JSON body to send a single transaction to Kafka. This lambda is triggered solely from the API Gateway, using the JSON body of the /publish API to send a single transaction to Kafka to be processed by rewarder.

### Deploying services locally
As mentioned, each folder has a Dockerfile, which allows us to start all the services up from a docker-compose.yml file, by specifying a Dockerfile target. You can use the following syntax, using the informer service as an example

```
  informer-service:
    build: 
      context: ./informer
      dockerfile: Dockerfile.prod
    
    ...other properties
```

 If you'd like to deploy all services, you can use the [`docker-compose.yml`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/docker-compose.yml) at the root of the repository as a reference. 

While that file runs select services, one could add the remaining ones or start them up as containers separately. To compose up, use:

```bash
docker compose up -d --build
```

To compose down, use:
```bash
docker compose down
```

### Frontend
---

- The [`/ui`](https://github.com/cs301-itsa/project-2022-23t2-g1-t2/tree/master/ui) folder is the folder containing the frontend. We utilise the NextJS framework that provides us with powerful server-side rendering capabilities using React and Typescript. For the UI library, we mainly use Chakra UI to create a pleasant User Experience.
- We have configured an auto deploy to our production URL as well as staging environments using Vercel, a popular site hosting SaaS. We've configured auto-deployments on Vercel's end, hence automating the CI/CD of our frontend.
- The client can be deployed locally with the following instructions

```bash
npm i && npm run dev
```

## Important Libraries/Frameworks used
---

- [Gin](https://github.com/gin-gonic/gin)
    - A web framework library, the choice of many Golang developers for writing performant RESTful APIs
- [Gorm](https://github.com/go-gorm/gorm)
    - Library in Golang, allowing us to keep Object Relational Mappings (ORMs) cleanly defined right within the structs defined in a service. Gorm also provides database connection functionality and connection pooling out of the box.
- [Sarama](https://github.com/Shopify/sarama)
    - Highly performant, open-sourced Kafka client for Golang. Made setting up performant Kafka producers and consumers in our services a breeze
- [Etcd Client](https://github.com/etcd-io/etcd)
    - A Golang client to connect to etcd. This library was particularly useful for us to watch for changes in etcd instead of continually polling, improving performance significantly.
- [Next.js](https://github.com/vercel/next.js/)
    - React framework for the web.
- [Swaggo](https://github.com/swaggo/swag)
    - Golang library to convert annotations to Swagger 2.0 Documentation. Aided our development and automated the process of generating API documentation.
- [Viper](https://github.com/spf13/viper)
    - Popular configuration & secrets management library for Golang. This library parses our environment variables neatly on application startup



## Images of our app

Here are some selected screengrabs of facets of the app

### Web App UI - Admin View

This screenshot shows the main page of UI where admin can view prevailing campaigns and exclusions, and add their own

![UI Rewards](docs/admin.png)

### Web App UI - Upload files

This screenshot shows a part of the admin portion of the UI where admins can drop bulk transaction CSV, or add users via a CSV.

![UI Rewards](docs/upload.png)

### Kafka UI - Consumer Group `reward`

This screenshot shows the number of messages the cluster of rewader service pods is yet to consume.

![Kafka UI Consumer Screen](docs/kafkaui_consumer.png)

### Kafka UI - Topic `transaction`

This screenshot shows the number of messages present in the topic of concern: `transaction`

![Kafka UI Topic](docs/kafkaui_topic.png)
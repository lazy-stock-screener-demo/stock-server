# Stock Server

![Preview](https://drive.google.com/uc?export=view&id=147OYUxLMpBoquxsEgs6ivZMLPwaXFXSp)

## What is it?

It is a golang server based on the idea of clean architecture and DDD best practices with event-driven architecture. This project is aimed to be executed in terms of both microservice and monolithic since each context has been separated into individual folder. Due to the fact of traffic is in low amount, monolithic is preferred in current version. However, it is always to free to expand to microservice or even CQRS architecture.

One more thing, a central event mediator that be implemented inside golang server can be replaced by Kafka if traffic grow.

## Why did I share this project in public?

The full product is still under constructing and the full product is committed to gitlab. I only share part of the use-case in order to demonstrate how I built up the stock-server.

## What features do this service have?

- Auto-complete search
- Read stock
- Sign-up/Login in

## How did I build this project?

- Based on DDD tactic context and user story, each user story is going to turn into use-case in application layer.
- Use data-transfer object (dto) to define application interface.
- Leverage mapper object to transform data to domain object, schema in persistence or view object.

### Overview

- Everything is use-case, no matter it is worker, event handler or normal api handler.
- Workers are executed immediately when server start, e.g. building ternary tree for auto-complete searching feature.

- Clean architecture

  - References
    [Clean Architecture and Modular pattern](https://en.bbo.com.ph/tech/clean-architecture-and-modular-pattern/)
    ![Clean Architecture](https://drive.google.com/uc?export=view&id=1P0Zl80q2_FLIRAE2KU5GxDRz848syNMu)

- Controller layer define various of adapters firing use-case, e.g. api controller, graphql controller or worker controller.

### File Structure

- application
  - dto
    - CatalogDTO.go
  - mapper
    - Catalog
      - CatalogViewMap.go
      - ProfitMap.go
  - useCase
    - readStockByTicker
      - index.go
      - readStockByTickerController.go
      - readStockByTickerDTO.go
      - readStockByTickerUseCase.go
      - readStockByTickerErrors.go
- domain
  - Stock.go
  - StockTicker.go
  - StockView.go
- infra
  - http
    - restful
    - graphql
  - repo
    - mongodriver
    - schema

## TechStacks

- Golang
- FastHTTP router
- jwt-go
- redis
- mongodriver
- gorm
- multi-stage dockerfile

### Env

```
HOST_URL=0.0.0.0
TEST_PORT=4999
KAFKA_DEV_HOST=localhost
KAFKA_DEV_PORT=9092
KAFKA_PROD_HOST=localhost
KAFKA_PROD_PORT=9092
IDENTITY_DEV_HOST=0.0.0.0
IDENTITY_DEV_PORT=6379
IDENTITY_DEV_DB_NAME=0
IDENTITY_DEV_USER=admin
IDENTITY_DEV_PASSWORD=admin
REDIS_PROD_HOST=localhost
REDIS_PROD_PORT=6379
CUSTOMER_SELF_DEV_DB_NAME=customer-db
CUSTOMER_SELF_DEV_USER=postgres
CUSTOMER_SELF_DEV_PASSWORD=postgres
CUSTOMER_SELF_DEV_HOST=0.0.0.0
CUSTOMER_SELF_DEV_PORT=5435
CATALOG_DEV_DB_NAME=stock-db
CATALOG_DEV_USER=admin
CATALOG_DEV_PASSWORD=admin
CATALOG_DEV_HOST=0.0.0.0
CATALOG_DEV_PORT=27017
CATALOG_CACHE_DEV_HOST=localhost
CATALOG_CACHE_DEV_PORT=6380
CATALOG_CACHE_DEV_DB_NAME=0
CATALOG_CACHE_DEV_USER=admin
CATALOG_CACHE_DEV_PASSWORD=admin
```

## References

- Inspiration of structure

  - [https://github.com/stemmlerjs/ddd-forum](https://github.com/stemmlerjs/ddd-forum)

- How to build a project with microservice architecture from the scratch
  - [http://mikedutuandu.blogspot.com/2019/11/how-to-build-project-with-microservice.html](http://mikedutuandu.blogspot.com/2019/11/how-to-build-project-with-microservice.html)
- JWT in go
  - **[https://dev.to/stevensunflash/a-working-solution-to-jwt-creation-and-invalidation-in-golang-4oe4](https://dev.to/stevensunflash/a-working-solution-to-jwt-creation-and-invalidation-in-golang-4oe4)**
  - **[https://github.com/auth0/go-jwt-middleware](https://github.com/auth0/go-jwt-middleware)**
- Use Auth0 to authenticate user with JWT in Go
  - **[https://auth0.com/blog/authentication-in-golang/#Authorization-with-Golang](https://auth0.com/blog/authentication-in-golang/#Authorization-with-Golang)**
- Concurrency in go
  - [https://dev.to/stevensunflash/using-golang-concurrency-in-production-3ma4](https://dev.to/stevensunflash/using-golang-concurrency-in-production-3ma4)
- DDD in typescript
  - [https://github.com/stemmlerjs/ddd-forum](https://github.com/stemmlerjs/ddd-forum)
- **Auto complete feature in Tenary Search Tree**
  - [https://iq.opengenus.org/autocomplete-with-ternary-search-tree/](https://iq.opengenus.org/autocomplete-with-ternary-search-tree/)

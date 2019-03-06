# sample-go-api

This repo acts as a starter kit to create a REST api to perform basic operations like add, get and delete.

## Setup
Before running make, make sure you install `postgres`, create a database and update the details in `application.yml`

``` sh
make
```

## Start service
```sh
make start
```

## Test
Mocks are generated using mockery. Generated files are not checked in. 

```sh
make test
```
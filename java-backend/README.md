# Java-Backend

## Prerequisites

- [JDK 21](https://www.oracle.com/java/technologies/javase/jdk21-archive-downloads.html) or higher
- [Maven 3.8.6](https://maven.apache.org) or higher
- [PostgreSQL 14.5](https://www.postgresql.org/docs/14/release-14-5.html)

### Build the application

You can build the project and run the tests by running

```shell
mvn clean package
```

Or you can build without running unit tests

```shell
mvn clean package -Dmaven.test.skip=true
```

### Start application

```shell
mvn spring-boot:run
```

### Get information system health

```
http://{host}/status
```

### Get detail information system health

```
http://{host}/status?isDetail=true
```

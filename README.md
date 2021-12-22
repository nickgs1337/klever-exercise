# Hello,
This project implements gRPC with Golang, as per the proposed challenge.

### Considerations

Unfortunately, due to the holiday season, the combined development time with Klever was 24 hours (overnight). Which for normal conditions is more than enough time, but since this is my first Go project, time was pretty tight between learning Go, gRPC and running the exercise.

The project was impacted, unfortunately this project does not contain tests and any extras from the exercise. Which bothers me, because **I don't lack experience with tests or with Cloud**. I would like future opportunities to demonstrate my expertise in these areas.
## Instructions to project execution
The execution of this project is like any other project in Go, just turn on the database, download the dependencies and run the server.go file.

If you need more information on how to start the database, the docker command is provided at the end of this file.

### Observations
This repository includes an example for each method implemented in gRPC, you can find these examples in the **examples** folder.

### Starting database (MongoDB)
```shell script
docker run -d --name klever  -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=klever -e MONGO_INITDB_ROOT_PASSWORD=klever mongo
```





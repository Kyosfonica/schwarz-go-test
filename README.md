# Schwarz IT Code Review Repository

This repository is meant to be used in the onboarding process of Go developers
applying to Schwarz IT.

Code smells and security issues are made on purpose. All the code to be reviewed is
in the [review](review) directory.


# About the code reviewed

I have fixed all the security issues from the dependencies and updated the go version to 1.23.

I fixed the code smells, refactored some code to make it more readable and added some tests. I also fixed a race condition that was present in the code.

I have added a few tests with dependency mocks to ensure the logic is working as expected.

I moved the CPU check init to the main.go file, as it makes more sense to me to have it there1. I decided to leave it as I am not sure if that check is a requirement for this project.


## Project Structure

I tried to follow the folder structure provided, and I moved the config file to the API folder. It makes sense to me all API related files to be in the same folder.


## Docker

I also modified the dockerfile to use the go 1.23-alpine image. I added a Makefile to easily run the project, tests, and linters.
Unfortunately I couldn't verify why on Windows the docker container is unable to connect to the socket. I don't have other machines to test this issue, so I left it as it is.

If you run the code in local it should work fine, the server should be running on localhost:8080 and shutdown gracefully after the given period of time. By default, 1 year. I recommend to modify this value in main.go to a shorter period of time.

## Execution

To run the code in docker you can use the following make commands:
1. run the command ```make docker-build``` to build the docker image
2. run the command ```make docker-up``` to start the docker container
3. run the command ```make docker-down``` to shut down the docker container

To run the code in local I have added a make rule that automatically runs the server.
1. run the command ```make run``` to start the server from your local machine. Default host is localhost and port is 8080.

To run the tests you can use the following make commands:
1. run the command ```make test``` to run the tests with coverage

To run the linter you can use the following make commands:
1. run the command ```make lint``` to run the linter

## Final thoughts
I really enjoyed working on this test and I hope you like the changes I made.
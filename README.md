# Pento tech challenge

This project consist in a small full stack web app that satisfy the following user stories:

- As a user, I want to be able to start a time tracking session
- As a user, I want to be able to stop a time tracking session
- As a user, I want to be able to name my time tracking session
- As a user, I want to be able to save my time tracking session when I am done with it
- As a user, I want an overview of my sessions for the day, week and month
- As a user, I want to be able to close my browser and shut down my computer and still have my sessions visible to me when I power it up again.

## Deployment Instructions

### Requirements

For the backend, you need to have installed and configured [Go](https://golang.org).

### Dependencies

To install dependencies for the backend, run the following commands inside the project folder:
`make deps`

### Production build

A Dockerfile has been provided that will build the application and create an docker image ready to deploy it.

Inside the project folder, run the following command to build the image:

`docker build -t pento-task .`

Then, you can run a container:

`docker run -it -p 3000:8080 --rm pento-task`

The project will be available at http://localhost:3000/.

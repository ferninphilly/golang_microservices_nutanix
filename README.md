# Golang Microservices

# Lab Setup

## Setting up your local development environment with Docker

### Downloading and installing docker on your local machine

Obviously- before we go on we need to ensure that everyone is using the same operating system. In this class we will be utiliing [Docker](https://www.docker.com/) to create "mini vms" for all of us.
Yes...I'm aware that docker containers are *not* vms. What I'm actually after here is simply a thin image that allows us all to be on the same OS. This should save time for us down the road as we all won't have to download Vagrant and heavy ISO images to get everything running.
To download Docker on your local machine go [here](https://www.docker.com/get-started)
Once it is installed you should see a docker image in your "running apps" bar (on mac and windows they are different).
Open up a command line editor and type in `docker info`. If you get a bunch of stuff- congrats! You've got DOCKER!

#### Creating your docker container

1. Now we'll need to get an image from the docker repo. For golang let's use the standard GO [image](https://hub.docker.com/_/golang). Go ahead and download that with `docker pull golang`

2. Once you've downloaded the container run the following command:
`docker run -ti -v $HOME/{pathtothisrepo}/src:/go golang` 

3. Now that we're in the docker container navigate to the `src` folder. You should see a `testme.go` file in there (if you don't check your volume mount path from the previous step).

4. Run the following command: `go run testme.go`. 

5. Did you get an output indicating that your test has passed? If so then **Congratulations**. You are set up! 


Now that we're set up for success let's get started!



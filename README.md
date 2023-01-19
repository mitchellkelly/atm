# ATM Coding Exercise

This project mimics the functionality of an ATM in a command line program.

To run this program, compile the project to a binary and run the binary via the command line.

## Docker

To build this project as a Docker image:

cd into the project directory and build the image using the following command:

```
docker build -t mitchellkelly/atm .
```

To run the project in the docker container run the following command after building the image:

```
docker run -it mitchellkelly/atm
```

This project uses a finite state machine to manage the ATM state. A diagram of the ATM FSM is below.

![ATM FSM](../media/fsm.png?raw=true)

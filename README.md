# Wallet

A simple application to track your expenses :)

## Usage

Create the docker image:

```sh
docker image build -t wallet-go .
```

Run the container:

```sh
docker run -it --rm -p 8080:8080 --name wallet-go-v0 wallet-go
```

## Example requests

Create a new user registration. A user registration represents the intent to register
a new user. The registration must be confirmed in order to create an actual account and
profile.

```sh
curl -X POST -v http://localhost:8080/api/users --data '{"username": "torreao", "password": "myPassw0rd$", "phone": "+5511916547815"}'
```


# Ditto - Distributed Tangle of Technological Organisms

Ditto is a distributed Ledger inspired by the Ethereum Project. I must say the name is a bit misleading, as Ditto is, unlike IOTA, not a _tangle_, but rather a chain of nodes.

## Installation

To run a Ditto Node, Docker must be installed on the Machine. Get Docker [here!](https://docs.docker.com/install/)

After installing Docker, open a command prompt and run the following command to download and run the latest container image of the project

```bash
docker run -p 42000:42000 garritfra/ditto:latest
```

Open your Browser and navigate to `http://localhost:42000`.

Voila, your personal Blockchain! Your JSON Response might look something like this:

```json
[
  {
    "Timestamp": "2018-11-28T21:44:25.1737195Z",
    "Hash": "00000f1908e761bc6cf8088aa8e50bb254221fe191fb3d91ac04ca1ab3da2847",
    "PreviousHash": "0",
    "Data": [],
    "Nonce": 107703
  }
]
```

## Interacting with your blockchain

The REST API currently exposes these Endpoints:

| Method | Route       | Description                                |
| :----- | :---------- | :----------------------------------------- |
| GET    | /           | Get full blockchain                        |
| POST   | /           | Add pending transaction                    |
| GET    | /mine_block | Mine a block with all pending transactions |

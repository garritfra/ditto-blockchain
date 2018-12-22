# Ditto - Distributed Tangle of Technological Organisms

Ditto is a distributed Ledger inspired by the Ethereum Project. I must say the name is a bit misleading, as Ditto is, unlike IOTA, not a _tangle_, but rather a chain of nodes.

## Seriously, another Blockchain?

**YES**, Blockchain is awesome, and we cannot have enough of them. I have started this project with the intention to learn the more advanced concepts of Blockchain development as well as to share my approach of a simple blockchain, to give newcomers to this field an overview of the internals of a blockchain, without needing to dig through a huge codebase full of bloat and complex abstractions. In the following chapter I will guide you through the installation of the blockchain server, so you can follow along with the code. If you otherwise know what to do with the source code, go ahead and compile your own server!

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
| GET    | /is_valid   | Check, if the chain has any faulty blocks  |

Refer to the [wiki](https://github.com/garritfra/blockchain-project/wiki) to get more information about the REST API
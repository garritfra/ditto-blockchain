# Ditto - Distributed Tangle of Technological Organisms

Ditto is a distributed Ledger inspired by the Ethereum Project. I must say the name is a bit misleading, as Ditto is, unlike IOTA, not a _tangle_, but rather a chain of nodes.

## Installation

To run a Ditto Node, Docker must be installed on the Machine. Get Docker [here!](https://docs.docker.com/install/)

After installing Docker, open a command prompt and run the following command

```bash
docker pull garritfra/ditto:latest
```

```bash
docker run garritfra/ditto:latest
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

## HD wallet with golang

### Solution:

#### Storing data:

For quickly, I store data directly on global variables in process.

But it should be stored in a relation ship database. The table diagrams can be more detail in postgres.sql file in folder setup

#### Flow of process:

1. For building wallet, we will generate master private key (mnemonic) and master public key
2. When add wallet we will send master public key such as a request
3. And receive back the id of the new wallet
4. For creating new address, we will send id of wallet and currency (for more detail of currency type check below)

#### Generate address with only master public key:



#### API docs:

***Generate address with only master public key:***

`URL`: `/v1/wallet`

`Method`: `POST`

`Header`: `json`

`Request`:

```json
{
	"master_public_key": "xpub661MyMwAqRbcEchqhvjmDox2feNsqWwWTw4hWTemBx9asYS8r9SYKUkZr4a2DkEE7UpeZMhqcX4EsxDpEQBJTjn2ytnZ5EBnj6H8gf4S9P6",
	"path": "m/0"
}
```

`Response`:

```json
{
    "addresses": [
        {
            "type": "ETH",
            "address": "0xB6ff335fa47F0e3dC6491f21a02527dDc3bd21d3"
        },
        {
            "type": "BTC",
            "address": "1CshzWY19rwvNFFnfcCgJQ3yyUs5RZZi2Z"
        }
    ]
}
```

***Generate master keys:***

`URL`: `/key`

`Method`: `POST`

`Header`: `json`

`Request`: 

```json
{
    "password": string
}
```

`Response`:

```json
{
    "id": string,
    "private_key": string,
    "public_key": string,
    "mnemonic": string
}
```

***Add wallet:***

`URL`: `/wallet`

`Method`: `POST`

`Header`: `json`

`Request`: 

```json
{
    "master_public_key": string
}
```

`Response`:

```json
{
    "id": string
}
```

***Add address:***

`URL`: `/retrieve-address`

`Method`: `POST`

`Header`: `json`

`Request`: 

```json
{
    "currency": int,
	"wallet_id": string
}
```

`Response`:

```json
{
    "id": string,
    "index": int,
    "currency": int,
    "address": string
}
```

### Run:

1. For golang build: run command: `make build-run`

2. For binary file only: run command: `make run-build`

3. For docker run: run command: `make docker-run`

### Checking CI/CD setup in .circleci folder:

#### Steps:

1. 3 steps: test, build, deploy
2. Test step: install pkgs and run `go test ./...`
3. Build step: build docker image by Dockerfile
4. Build step: Push docker image to docker hub
5. Deploy step: Ssh to server 
6. Deploy step: Pull image from server and run it to container

For more detail checking file config.yml in folder .circleci

### Checking for production server:

Checking in IP: `http://35.198.235.51:6000/`

### Note:

#### Currency Type:

0. BTC - 0
1. BCH - 1
2. LTC - 2
3. DOGE - 3
4. ETH - 4
5. ETC - 5


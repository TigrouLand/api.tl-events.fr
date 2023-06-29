# api.tl-events.fr

A micro-application providing a public API of various TL events data, which is used in the main site application: [tl-events.fr](https://github.com/TigrouLand/tl-events.fr).

## Used technologies

Thanks to the use of [Golang](https://go.dev/), this micro-application is powerful, stable and extremely low cost in machine resources. In order to ensure a continuous deployment and to secure this application, being hosted on our servers, we use [Docker](https://www.docker.com/) in order to build an image at each new commit on the main branch.
## Run locally

```bash
  # First, clone this repository on your machine:
  git clone https://github.com/TigrouLand/api.tl-events.fr.git

  # Go to the project directory:
  cd api.tl-events.fr

  # Install dependencies:
  go get

  # Compile the application:
  go build
  
  # Run the application:
  ./api
```

## Environment variables

| Name                   | Description |
|------------------------|-------------|
| `PORT`                 | The port on which the application will be listening                     |
| `MONGO_URI`            | The URI of the MongoDB database |
| `MONGO_DATABASE`       | The name of the MongoDB database |
| `DISCORD_REDIRECT_URL` | The URL to which the Discord OAuth2 authentication will redirect the user |
| `DISCORD_CLIENT_ID`    | The Discord OAuth2 client ID |
| `DISCORD_CLIENT_SECRET` | The Discord OAuth2 client secret |
| `COOKIE_SECRET`        | The secret used to encrypt the cookies |
| `FRONTEND_URL`         | The URL of the frontend application |


## License
This project is licensed under the [GNU GPL 3](https://github.com/TigrouLand/api.tl-events.fr/blob/main/LICENCE).


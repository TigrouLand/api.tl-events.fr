# api.tl-events.fr

A micro-application providing a public read-only API of various TL events data, which is used in the main site application: [tl-events.fr](https://github.com/TigrouLand/tl-events.fr).

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

## License
This project is licensed under the [GNU GPL 3](https://github.com/TigrouLand/api.tl-events.fr/blob/main/LICENCE).


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [available-velibs](#available-velibs)
  - [Assignment](#assignment)
    - [Infos](#infos)
    - [Tasks](#tasks)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Run](#run)
  - [Description](#description)
  - [Data model](#data-model)
  - [API routes](#api-routes)
  - [Contributing](#contributing)
  - [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# available-velibs

Provides available Velib's around a given location - Splio technical test

## Assignment

In order to promote cycling and make life easier for our employees, we want to provide them, in real time, with the number of Vélib' available in the stations closest to our headquarters.

To do so, we will use the public data set of Vélib' stations availabilities: [https://opendata.paris.fr/explore/dataset/velib-disponibilite-en-temps-reel/information/](https://opendata.paris.fr/explore/dataset/velib-disponibilite-en-temps-reel/information/)

### Infos

- Splio HQ : 27 Boulevard des Italiens, 75002 Paris
- Vélib' data are refreshed every minute

### Tasks

- [x] The project will be written in Go with or without a framework
- [x] The results can be displayed in the application's console (for the most advanced version) or in an html page (for the most advanced)
- [ ] The project must be sent in zip format
- [x] Do not hesitate to send all the necessary details to understand your code.

Original assignment (in french) can be found [here](./docs/Splio_-_Go_Tech_Homework_-_fr.pdf).

## Prerequisites

- [docker](https://docs.docker.com/get-docker/) - version 19.03.13 was used for this project
- [docker-compose](https://docs.docker.com/compose/install/) - version 1.27.4 was used for this project
- [make](http://ftp.gnu.org/gnu/make/) - version 3.81 was used for this project
- [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) - version 2.26.2 was used for this project

## Installation

```bash
git clone https://github.com/Geoffrey42/available-velibs.git
cd available-velibs/
git checkout main # Production branch
```

Once at repository's root, create a **.env** file based on [.env.sample](./.env.sample) and fill it according to your environment:

```shell
HTTP_PROXY=  # only if your are behind a corporate proxy
HTTPS_PROXY= # only if your are behind a corporate proxy
CLIENT_PORT=
SERVER_PORT=
```

## Run

To run the application, simply do:

```bash
docker-compose up -d
```

## Description

**available-velibs** is a web application that shows currently available Velib's around Splio's HQ in real time (data is refreshed every minute).

The project is a docker composition of two main services:

- A web server written in Golang that fetches velibs related data every minutes from opendata's dataset.
- A frontend web client in React that displays the data.

## Data model

Here is a JSON representation of the data returned by the Golang server:

```json
{
    "Distance": 500, // distance around search position in meters
    "total": 14, // sum of numbikesavailable
    "nhits": 2,
    "records": [
        {
            "fields": {
                "name": "Jouffroy d'Abbans - Wagram",
                "numbikesavailable": 7,
                "coordonnees_geo": [
                    48.8819732984,
                    2.30113215744
                ]
            }
        },
        {
            "fields": {
                "name": "Courcelles - Pierre Demours",
                "numbikesavailable": 7,
                "coordonnees_geo": [
                    48.8837697333,
                    2.29858466343
                ]
            }
        }
    ]
}
```

See a full JSON sample [here](./assets/sample.json).

## API routes

Server endpoint is **/api**.

**GET**:

- **/api/fetch**: Get fresh data from opendata's dataset. See [Data model](#data-model) section for return value.

## Contributing

Pull requests are welcome.
For more details, please refers to our [contributing file](.github/CONTRIBUTING/contributing.md).

## License

[MIT](./LICENSE)
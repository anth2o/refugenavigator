# Refuge Navigator

[![Live Site](https://img.shields.io/badge/Live%20Site-refugenavigator.com-brightgreen)](https://refugenavigator.com)
[![coverage](https://codecov.io/gh/anth2o/refugenavigator/branch/main/graph/badge.svg?token=J6KEKKL5ZE)](https://codecov.io/gh/anth2o/refugenavigator)
[![pipeline](https://github.com/anth2o/refugenavigator/actions/workflows/pipeline.yml/badge.svg?query=branch%3Amain)](https://github.com/anth2o/refugenavigator/actions/workflows/pipeline.yml?query=branch%3Amain)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 🎯 About

Refuge Navigator exports data from [www.refuges.info](https://www.refuges.info/) (which is only available online) to GPX files, enabling offline access to refuge and water source information for hiking trips. The exported GPX files can be imported into navigation apps like [Iphigénie](https://www.iphigen.ie/) or any other application that supports GPX waypoints.

## 📋 Credits

Refuge Navigator uses data sourced from [www.refuges.info](http://www.refuges.info), licensed under [CC BY-SA 2.0](https://creativecommons.org/licenses/by-sa/2.0/).

- Content from Refuges.info is attributed to `©Les contributeurs de Refuges.info`
- Specific contributions such as comments and photos are attributed to their individual authors as specified on the site.

Thus, the data coming from the use of Refuge Navigator is also licensed under [CC BY-SA 2.0](https://creativecommons.org/licenses/by-sa/2.0/).

## 🛠️ Tech Stack

This project is also an opportunity for me to work on:

- **Backend**: Go with Gin framework - my first Go project since I've mostly worked with Python in all my previous projects
- **Frontend**: React with Leaflet for the interactive map, statically served by the Go backend
- **DevOps**: A CI/CD pipeline, with automated deployments to Google Cloud Run on each release [here](https://refugenavigator.com), and on each push to the main branch or a pull request [here](https://dev.refugenavigator.com). Note: deploying on every push isn't ideal for parallel contributions, but works fine for now.

## 🤝 Contributing

Pull requests are welcome! To get started:

### Backend (Go)

```bash
cd backend-go
go mod tidy
go run main.go
```

### Frontend (React)

```bash
cd frontend
yarn
yarn dev
```

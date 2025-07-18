# Refuge Navigator

[![coverage](https://codecov.io/gh/anth2o/refugenavigator/branch/main/graph/badge.svg?token=J6KEKKL5ZE)](https://codecov.io/gh/anth2o/refugenavigator)
[![pipeline](https://github.com/anth2o/refugenavigator/actions/workflows/pipeline.yml/badge.svg?query=branch%3Amain)](https://github.com/anth2o/refugenavigator/actions/workflows/pipeline.yml?query=branch%3Amain)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

You can access the latest version of Refuge Navigator at [https://refugenavigator.com](https://refugenavigator.com).

The purpose of this project is to export the data from the awesome website [www.refuges.info](https://www.refuges.info/) as GPX so that I can bring them offline when hiking. I personally import the downloaded GPX files to [Iphigénie](https://www.iphigen.ie/), but any app with support for GPX waypoints should do it.

It's also the opportunity for me to work on:

- My first Go project since I've mostly worked with Python in all my previous projects
- A CI/CD pipeline, with automated deployments to Google Cloud Run on each release [here](https://refugenavigator.com), and on each push to the main branch or a pull request [here](https://dev.refugenavigator.com)
- A little bit of frontend with React and Leaflet to have an interactive map
- In the future, integration with LLMs to be able to summarize unstructured data (such as the comments on water points) or to be able to select relevant pictures about waypoints

![refuges.info](images/refuges.info.png)

## Credits

Refuge Navigator uses data sourced from [www.refuges.info](http://www.refuges.info), licensed under [CC BY-SA 2.0](https://creativecommons.org/licenses/by-sa/2.0/).

- Content from Refuges.info is attributed to `©Les contributeurs de Refuges.info`
- Specific contributions such as comments and photos are attributed to their individual authors as specified on the site.

Thus, the data coming from the use of Refuge Navigator is also licensed under [CC BY-SA 2.0](https://creativecommons.org/licenses/by-sa/2.0/).

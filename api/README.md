# Managym API

## 📋 Table of Contents

- [Managym API](#managym-api)
  - [📋 Table of Contents](#-table-of-contents)
  - [📖 Summary](#-summary)
  - [🛠️ Technologies](#️-technologies)

## 📖 Summary

This is the repository containing the source code for our REST API. This branch features the application rewrite using Go, Fiber, go-pg and Air for hot reloading. It aims to manage the back-end of Managym app which interacts with the database and the front-end.

## 🛠️ Technologies

These are the technologies used in this project, along with their versions.

|           Name           | Version |
| :----------------------: | ------- |
|         [Go][go]         | 1.22.4  |
| [PostgreSQL][postgresql] | 16.3    |
|      [Fiber][fiber]      | 2.x     |
|      [go-pg][gopg]       | 10.13   |
|        [Air][air]        | 1.52.3  |

We recommend installing these tools in the same versions or higher to ensure
compatibility.

[go]: https://go.dev
[gopg]: https://github.com/go-pg/pg
[fiber]: https://docs.gofiber.io/
[postgresql]: https://www.postgresql.org/download/
[air]: https://github.com/air-verse/air

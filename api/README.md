# Managym API

## 📋 Table of Contents

- [Managym API](#siege-cloud-api)
  - [📋 Table of Contents](#-table-of-contents)
  - [📖 Summary](#-summary)
  - [🛠️ Technologies](#️-technologies)

## 📖 Summary

This is the repository containing the source code for our REST API. This branch features the application rewrite using Go, Fiber and go-pg. It aims to manage the back-end of Siege Cloud which interacts with the database and the front-end.

## 🛠️ Technologies

These are the technologies used in this project, along with their versions.

|           Name           | Version |
| :----------------------: | ------- |
|         [Go][go]         | 1.22.4  |
| [PostgreSQL][postgresql] | 16.3    |
|      [Fiber][fiber]      | 2.x     |
|      [go-pg][gopg]       | 10.13   |

We recommend installing these tools in the same versions or higher to ensure
compatibility.

[go]: https://go.dev
[gopg]: https://github.com/go-pg/pg
[fiber]: https://docs.gofiber.io
[postgresql]: https://www.postgresql.org/download/

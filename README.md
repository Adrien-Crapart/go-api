[![Create Go App][repo_logo_img]][repo_url]
# Go API services


[![Go version][go_version_img]][go_dev_url] [![Code coverage][go_code_coverage_img]][repo_url] [![Wiki][repo_wiki_img]][repo_wiki_url] [![License][repo_license_img]][repo_license_url]

This api allows to know the exposure of plots or a location on a set of spatial data.
<!-- 
## 📖 Documentation
Check database datas
Check routes by API -->

## Getting Started
> ❗️ **To install Go for the first time and set up this API, you can follow these steps:**
> Download Go from the official website: https://golang.org/dl/
Choose the appropriate installer for your operating system and architecture, and download it.
> 
> Run the installer and follow the instructions to install Go on your machine.

<details>
<summary>More details</summary>

Set up your Go workspace by creating a directory where you will store your Go code. For example, you can create a directory called go-api:
```
mkdir go-api
```
Navigate to the go-api directory:
```
cd go-api
```
Clone the API repository into the go-api directory:
```
git clone <repository_url>
```
Replace <repository_url> with the URL of the API repository.
Install the necessary dependencies for the API by running the following command:
```
go mod download
```
Once the dependencies are installed, you can build and run the API using the following command:
```
go run main.go
```
This will start the API server, and you can access it at http://localhost:8080.
That's it! You have now installed Go and set up the API for the first time.
</details>

### ⚡️ Run Locally
```bash
  go run main.go
```
### 🐳 Run in docker
```bash
  docker-compose up --build
```

## ⚙️ Architecture of this services
```
project
│   main.go
|   .env
└───api
│   └───handlers
│       │   user_handler.go
│   └───middlewares
│       │   auth_middleware.go
└───db
│   │   postgres.go
│   │   models.go
└───config
    │   config.go
```
<details>
<summary>More details</summary>

> #### General part (main.go)
> Entry point of your application. It initializes your server, router, database connection, middleware, and starts listening for incoming requests.
> #### Handlers part
> Handlers (api/handlers): Contains HTTP request handlers. For instance, user_handler.go would handle requests related to user management (registration, login, profile update, etc.).
> #### Middlewares part
> Middlewares (api/middlewares): Contains middleware functions that perform tasks such as authentication, logging, error handling, etc. auth_middleware.go would handle authentication.
> #### Configuration part
> Config (config): Contains configuration settings for your application.
> #### Database part
> Database (db): Contains database-related logic. postgres.go would handle PostgreSQL database connections and querying, while models.go would define your database schema as Go structs.
</details>

## 📝 API Reference
### Get all items


## Contributing
Contributions are always welcome!
See `contributing.md` for ways to get started.
Please adhere to this project's `code of conduct`.

## Authors
- [@AdrienCrapart](https://github.com/AdrienCrapart)

<!-- Go -->

[go_download_url]: https://golang.org/dl/
[go_install_url]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_version_img]: https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none
[go_report_url]: https://goreportcard.com/report/github.com/create-go-app/cli
[go_code_coverage_img]: https://img.shields.io/badge/code_coverage-100%25-success?style=for-the-badge&logo=none
[go_dev_url]: https://pkg.go.dev/github.com/create-go-app/cli/v4

<!-- Repository -->

[repo_url]: https://github.com/create-go-app/cli
[repo_logo_url]: https://github.com/create-go-app/cli/wiki/Logo
[repo_logo_img]: https://github.com/create-go-app/cli/assets/11155743/95024afc-5e3b-4d6f-8c9c-5daaa51d080d
[repo_license_url]: https://github.com/create-go-app/cli/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_wiki_url]: https://github.com/create-go-app/cli/wiki
[repo_wiki_img]: https://img.shields.io/badge/docs-wiki_page-blue?style=for-the-badge&logo=none
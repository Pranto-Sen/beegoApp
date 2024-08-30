# Beego Cat API
This is a Beego-based project that provides a simple API for managing cat breeds, favorites, and voting. The project also includes basic views and static assets.

## Prerequisites
Make sure you have the following installed:

- Go (version 1.16 or later)
- Beego framework
- Git

## Installation

1. **Clone the repository:**
    ```
    git clone https://github.com/Pranto-Sen/beegoApp
    cd beegoApp
    ```
    **Folder structure**
    ```
    beegoApp/
    ├── conf/
    │   └── app.conf
    ├── controllers/
    ├── routers/
    ├── static/
    ├── tests/
    ├── views/
    ├── .gitignore
    ├── go.mod
    ├── go.sum
    └── main.go

    ```


2. **Create config file**
     In the root directory, under the `conf` folder, create a file  named  `app.conf`. In this file, set your `catapi_key` as follows.

    ```bash
    appname = beegoApp
    httpport = 8080
    runmode = dev
    catapi_key = your_api_key
    ```

3. **Install the dependencies:**
    ```
    go mod tidy
    ```

4. **Running the Application**

    ```
    bee run
    ```
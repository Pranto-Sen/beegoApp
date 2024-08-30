# Beego Cat API
This is a Beego-based project that provides a simple API for managing cat breeds, favorites, and voting. The project also includes basic views and static assets.

## Prerequisites
Make sure you have the following installed:

- Go (version 1.16 or later)
- Beego framework
- Git

## Installation
1. **Install Go**

    - Ensure that Go is installed on the new device. You can check if Go is installed and verify the version with: `go version`
    - If Go is not installed, download and install it from the official Go website.

2. **Install Beego and Bee CLI**

    - Install Beego and the Bee CLI tool globally if they are not already installed:

        ```bash
        go get github.com/beego/beego/v2
        go install github.com/beego/bee/v2@latest
        ``` 
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


4. **Create config file**
    - In the root directory, under the `conf` folder, create a file  named  `app.conf`. In this file, set your `catapi_key` as follows.
    - Visit The Cat API website, register with your email, and check your inbox for the API key they send.

        ```bash
        appname = beegoApp
        httpport = 8080
        runmode = dev
        catapi_key = your_api_key
        ```

5. **Install Project Dependencies**
    <!-- - Fetch and install all project dependencies: `go get ./...` -->
    
    - Then tidy up the go.mod file: 
        ```
        go mod tidy
        ```

    

6. **Running the Application**

    ```bash
    bee run
    ```

7. You can access the application in your browser at http://localhost:8080
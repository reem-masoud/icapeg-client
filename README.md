# icapeg-client

icapeg-client is application used to test the ICAP server 

## Things to have

1. **Golang**(latest enough to be able to use go mod)

     ***A sample installation of go version 1.14***:

     Prepare the apt packages    
  ```bash
    sudo apt update

  ```

  ```bash
    sudo apt upgrade

  ```

Link of download of version 1.14
    https://dl.google.com/go/go1.14.linux-amd64.tar.gz

Use the command
  ```bash
    wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz

  ```
Untar in /usr/local

  ```bash
    tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz

  ```

Add /usr/local/go/bin to the PATH environment variable:

  ```bash
    export PATH=$PATH:/usr/local/go/bin

  ```
> Note: this command adds the PATH temporarily to the environment variables, the path is removed if SSH is broken, or system reboots.

To discover the go version you installed

  ```bash
    go version

  ```


## How to run icapeg-client!!

1. Install the icap-client library

  ```bash
    go get -u github.com/egirna/icap-client

  ```

2. Clone the ICAPeg-client repository

  ```bash
    https://github.com/egirna/icapeg-client.git

  ```

3. change the directory to the repository

  ```bash
    cd icapeg-client

  ```

4. Run the application

  ```bash
    go run main.go

  ```

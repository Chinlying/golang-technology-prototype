## [Download and install](https://go.dev/doc/install)

+ install
```go
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz
```

+ configure environment variables
  You can do this by adding the following line to your ~/.bashrc or /etc/profile (for a system-wide installation):
```shell
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
`Note:`  Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as `source /etc/profile`
+ verification
```shell
go version
```
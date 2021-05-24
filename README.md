# whats my ip

![White circle with "whats my ip" and a purple question mark](./logo.png)

## Motivation

I need a service to retrieve the external ip address of my services.
This is a simple golang lambda which achieves that.

## Requirements

* [go toolchain](https://golang.org/dl/)
* [up binary in $PATH](https://apex.sh/docs/up/configuration/)

## Deploy
```
$ make deploy
```

## Usage

### Running the server
```
$ make run
```

### Testing the remote
```
$ make test-remote
```

## License

This project is licensed under the GPL-3 license.

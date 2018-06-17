# Yoo [![JavaScript Style Guide](https://img.shields.io/badge/code_style-standard-brightgreen.svg)](https://standardjs.com)

A runtime and compiler of [TypeScript](typescriptlang.org) built on [Golang](golang.org). **WIP**

## Install

```bash
go get -u github.com/Apisium/Yoo/...

cd $GOPATH/src/github.com/Apisium/Yoo

npm install
```

## Features

- No `package.json`. No npm.

- Imports reference source code URLs only.

- No `prototype`, only `class`.

- No `var`. No hoisting.

- No `with`. No `arguments`. No `eval`.

- Error stacks of `Promise` always are printed.

- Full asynchronous API, based on `Promise`.

## Usage

### Compile

```bash
npm i -g ts-node

ts-node complier/index.ts
```

### Execute

```bash
go run command/main.go
```

## Author

[Shirasawa](https://github.com/ShirasawaSama)

## License

[MIT](./LICENSE)

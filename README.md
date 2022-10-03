# Joker

Get random Chuck Norris jokes in terminal. Written in Go.

This project fetch the [chucknorris.io API](https://api.chucknorris.io).

![Screenshot](./screenshot.png)

## Installation

```bash
go get github.com/Doaxan/joker_go
```

Or:

```bash
go install github.com/Doaxan/joker_go@latest
```

## Usage
Retrieve a random chuck joke
```bash
joker_go random
```
Get 5 random unique jokes for each of the existing categories and saves them to text files in ./jokes/ directory - one for each of the categories
```bash
joker_go dump
```
You can specify the number of jokes, but note that in some categories(career, fashion, for ex.), jokes are less than 5. 
```bash
joker_go dump -n 7
```

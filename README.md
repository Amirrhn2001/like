# like
like service

## contetnts
* [General info](#general-info)
* [Technologies](#technologies)
* [Architecture](#architecture)
* [How to filter](#how-to-filter)
* [Setup](#setup)

## General info
this is a service to record likes and dislikes of an entity.

## Technologies
* Gin Framework
* MongoDB

## Architecture
* Hexagonal

## How to filter
To use filter use this json array of object in query string:\n
[{"field":"field name", "value":"value","operator":"mongoDB operator"}]\n
also pagination is required in query string:\n
{"skip":0,"limit":1}

### Operator you allow to use
* $eq
* $ne
* $gt
* $gte
* $lt
* $lte
* $in
* $nin
* $exists
* $type

## Setup
To run this project:
```
$ go mod download
$ go mod tidy
$ go run .
```

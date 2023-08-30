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

## Ho to filter
To use filter use this json array of object in query string:
[{"field":"field name", "value":"value","operator":"mongoDB operator"}]

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

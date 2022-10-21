# Zenhub-Integration-POC

This is a playground to determine how to use the ZenHub API in order to automate adding a sprint
to a new issue when that sprint begins.

## Usage

`go run cmd/cmd.go`

## Current state

Currently, this project is just a test in running GraphQL commands with GoLang.  To practice,
this project currently runs a GraphQL query against the [SpaceX API](https://studio.apollographql.com/public/SpaceX-pxxbxen/home?variant=current)

Future Plans:
    - Get the current pipeline (In Progress/Backlog/etc) for each issue
    - Add a sprint to each issue through the API
    - Add a sprint to only the issues that are in progress and 


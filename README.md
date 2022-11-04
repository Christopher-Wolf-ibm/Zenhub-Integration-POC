# Zenhub-Integration-POC

This is a playground to determine how to use the ZenHub API in order to automate adding a sprint
to a new issue when that sprint begins.

## Setup

You will need a few environment variables to make this application work.

### ZenHub API Key
You will need an API key to authorize access to ZenHub.  It is expected to be stored as the `ZENHUB_KEY` environment variable.  Documentation to acquire a ZenHub API key (for IBM users only)
can be found [here](https://zenhub.ibm.com/settings/tokens).

## GitHub Repository ID
You will need the repository ID for the GitHub repo you are trying to work with.  It is expected
to be stored as `GITHUB_REPO_ID`.  You can acquire the repo ID by running the below command

```
curl https://api.github.com/repos/<username>/<reponame>
```

It will be stored as the `id` field in the returned JSON.
NOTE: You may need to include a GitHub API key in your curl header if you are working with
a private repository.

## Usage

`go run cmd/cmd.go`

## Current state

Currently, this project is just a test in running GraphQL commands with GoLang.  To practice,
this project currently runs a GraphQL query against the [SpaceX API](https://studio.apollographql.com/public/SpaceX-pxxbxen/home?variant=current)

Future Plans:
    - Get the current pipeline (In Progress/Backlog/etc) for each issue
    - Add a sprint to each issue through the API
    - Add a sprint to only the issues that are in progress and 


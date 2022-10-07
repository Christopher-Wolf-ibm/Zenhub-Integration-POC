#! /bin/bash

# gh api /octocat --method GET
gh api --header 'Accept: application/vnd.github.v3+json' --method GET /repos/Christopher-Wolf/Zenhub-Integration-POC/issues | jq .[].title
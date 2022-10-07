# Zenhub-Integration-POC

This is a playground to determine how to use the ZenHub API in order to automate adding a sprint
to a new issue when that sprint begins.

## Usage

In order to use this service, you must log in to the GitHub CLI.  The easiest way is to run 
`gh auth login`

Then you can run this service with the command:
`bash scripts/get_issues.sh`

This will provide you a JSON payload with data regarding all of the issues currently open in this ZenHub repo.

Future Plans:
    - Get the current pipeline (In Progress/Backlog/etc) for each issue
    - Add a sprint to each issue through the API
    - Add a sprint to only the issues that are incomplete
    - Automate adding a sprint whenever it begins to incomplete issues


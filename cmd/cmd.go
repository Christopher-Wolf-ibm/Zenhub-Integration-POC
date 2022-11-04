package main

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
)

func main() {
	// ENDPOINT := "https://api.spacex.land/graphql/"
	ENDPOINT := "https://api.zenhub.com/public/graphql"

	GITHUB_REPO_ID := os.Getenv("GITHUB_REPO_ID")
	// ZENHUB_KEY := os.Getenv("ZENHUB_KEY")

	graphqlClient := graphql.NewClient(ENDPOINT)
	graphqlRequest := graphql.NewRequest(`
	query getIssueInfo($repositoryGhId: Int!, $issueNumber: Int!, $workspaceId: ID!) {
		issueByInfo(repositoryGhId: $repositoryGhId, issueNumber: $issueNumber) {
		  id
		  repository {
			id
			ghId
		  }
		  number
		  title
		  body
		  state
		  pipelineIssue(workspaceId: $workspaceId) {
			priority {
			  id
			  name
			  color
			}
			pipeline {
			  id
			  name
			}
		  }
		  estimate {
			value
		  }
		  sprints(first: 10) {
			nodes {
			  id
			  name
			}
		  }
		  labels(first: 10) {
			nodes {
			  id
			  name
			  color
			}
		  }
		}
	  }
	`)
	graphqlRequest.Var("issueNumber", 1)
	graphqlRequest.Var("repositoryGhId", GITHUB_REPO_ID)

	var graphqlResponse interface{}
	err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if err != nil {
		fmt.Println("could not complete request")
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(graphqlResponse)
	}
}

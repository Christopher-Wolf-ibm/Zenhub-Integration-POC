package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func main() {
	ENDPOINT := "https://api.spacex.land/graphql/"

	graphqlClient := graphql.NewClient(ENDPOINT)
	graphqlRequest := graphql.NewRequest(`
	{
		launchesPast(limit: 10) {
		  mission_name
		  launch_date_local
		  launch_site {
			site_name_long
		  }
		  links {
			article_link
			video_link
		  }
		  rocket {
			rocket_name
		  }
		}
	  }
	`)

	var graphqlResponse interface{}
	err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if err != nil {
		fmt.Println("could not complete request")
	} else {
		fmt.Println(graphqlResponse)
	}
}

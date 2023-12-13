package main

import (
	"context"
	"fmt"
	"html/template"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

const sponsorsOutput = `{{range .}}
<a href="https://github.com/{{.Login}}">
	<img src="{{.AvatarURL}}" class="img github-avatar" alt="{{.Name}}">
</a>
{{end}}`

var sponsorsOutputTpl = template.Must(template.New("sponsorsOutputTpl").Parse(sponsorsOutput))

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	var query struct {
		Organization struct {
			Sponsors struct {
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
				Edges []struct {
					Node struct {
						User struct {
							Login     string
							Name      string
							AvatarURL string
						} `graphql:"... on User"`
						Organization struct {
							Login     string
							Name      string
							AvatarURL string
						} `graphql:"... on Organization"`
						Sponsorable struct {
							Sponsorship struct {
								Edges []struct {
									Node struct {
										IsActive bool
										Tier     struct {
											MonthlyPriceInCents int
										}
									}
								}
							} `graphql:"sponsorshipsAsSponsor(maintainerLogins:\"syncthing\", first:5)"`
						} `graphql:"... on Sponsorable"`
					}
				}
			} `graphql:"sponsors(first:100, after:$cursor)"`
		} `graphql:"organization(login:\"syncthing\")"`
	}
	vars := map[string]any{
		"cursor": (*githubv4.String)(nil),
	}

	sponsors := []map[string]any{}

	for {
		if err := client.Query(context.Background(), &query, vars); err != nil {
			fmt.Println(err)
			return
		}

		for _, sponsor := range query.Organization.Sponsors.Edges {
			for _, sponsorship := range sponsor.Node.Sponsorable.Sponsorship.Edges {
				if sponsorship.Node.Tier.MonthlyPriceInCents >= 100*100 {
					s := map[string]any{
						"Name":      sponsor.Node.User.Name,
						"Login":     sponsor.Node.User.Login,
						"AvatarURL": sponsor.Node.User.AvatarURL,
						"Amount":    sponsorship.Node.Tier.MonthlyPriceInCents / 100,
					}
					sponsors = append(sponsors, s)
				}
			}
		}

		if !query.Organization.Sponsors.PageInfo.HasNextPage {
			break
		}
		vars["cursor"] = githubv4.NewString(query.Organization.Sponsors.PageInfo.EndCursor)
	}

	sponsorsOutputTpl.Execute(os.Stdout, sponsors)
}

package main

import (
	"cmp"
	"context"
	"fmt"
	"html/template"
	"net/url"
	"os"
	"slices"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

const sponsorsOutput = `{{range .}}
<a href="{{.LinkURL}}" rel="nofollow sponsored">
	<img src="{{.AvatarURL}}" class="img github-avatar" alt="{{.Name}}">
</a>
{{end}}`

var sponsorsOutputTpl = template.Must(template.New("sponsorsOutputTpl").Parse(sponsorsOutput))

// Static sponsors (predate Github Sponsors)
var staticSponsors = []sponsor{
	{
		Amount:    1000 * 100,
		Name:      "Kastelo",
		AvatarURL: "https://cdn.kastelo.net/prm/civ3ci-shield-blue-padded-white-512.png",
		LinkURL:   "https://kastelo.net/",
	},
	{
		Amount:    100 * 100,
		Name:      "REEF Solutions",
		AvatarURL: "/img/reefsol.png",
		LinkURL:   "https://reefsolutions.com/",
	},
}

type sponsor struct {
	Amount    int
	Name      string
	Login     string
	AvatarURL string
	LinkURL   string
}

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
							Login      string
							Name       string
							AvatarURL  string
							WebsiteURL string
						} `graphql:"... on User"`
						Organization struct {
							Login      string
							Name       string
							AvatarURL  string
							WebsiteURL string
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

	sponsors := staticSponsors
	for {
		if err := client.Query(context.Background(), &query, vars); err != nil {
			fmt.Println(err)
			return
		}

		for _, spons := range query.Organization.Sponsors.Edges {
			for _, sponsorship := range spons.Node.Sponsorable.Sponsorship.Edges {
				if sponsorship.Node.Tier.MonthlyPriceInCents >= 100*100 {
					s := sponsor{
						Name:      spons.Node.User.Name,
						Login:     spons.Node.User.Login,
						AvatarURL: spons.Node.User.AvatarURL,
						Amount:    sponsorship.Node.Tier.MonthlyPriceInCents,
					}
					s.LinkURL = urlFrom(spons.Node.User.WebsiteURL, "https://github.com/"+spons.Node.User.Login+"/")
					sponsors = append(sponsors, s)
				}
			}
		}

		if !query.Organization.Sponsors.PageInfo.HasNextPage {
			break
		}
		vars["cursor"] = githubv4.NewString(query.Organization.Sponsors.PageInfo.EndCursor)
	}

	// Sort by amount, highest first, then by name
	slices.SortFunc(sponsors, func(a, b sponsor) int {
		if v := cmp.Compare(b.Amount, a.Amount); v != 0 {
			return v
		}
		return cmp.Compare(a.Name, b.Name)
	})

	sponsorsOutputTpl.Execute(os.Stdout, sponsors)
}

// urlFrom returns the first reasonable, working URL in the list
func urlFrom(urls ...string) string {
	for _, alt := range urls {
		if alt == "" {
			continue
		}
		if u, err := url.Parse(alt); err == nil {
			if u.Scheme == "" {
				u.Scheme = "https"
			}
			if u.Host == "" && u.Path != "" {
				u.Host = u.Path
				u.Path = ""
			}
			if u.Path == "" {
				u.Path = "/"
			}
			return u.String()
		}
	}
	return ""
}

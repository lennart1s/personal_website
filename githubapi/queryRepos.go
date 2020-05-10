package githubapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
)

func GetGithubRepos(username string) []Repository {
	resp, err := http.Get("https://api.github.com/users/" + username + "/repos")
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	var repos []Repository
	json.Unmarshal(body, &repos)

	return repos
}

func GetLatestGithubRepos(username string, n int) []Repository {
	repos := GetGithubRepos(username)
	sort.Sort(byLastUpdate(repos))

	if len(repos) >= n {
		return repos[:n]
	}

	return repos
}

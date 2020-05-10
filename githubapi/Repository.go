package githubapi

type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"html_url"`
	Owner       Owner  `json:"owner"`
	Description string `json:"description"`
	Language    string `json:"language"`
	LastUpdate  string `json:"updated_at"`
}

type Owner struct {
	Name string `json:"login"`
	URL  string `json:"html_url"`
}

type byLastUpdate []Repository

func (repos byLastUpdate) Len() int {
	return len(repos)
}

func (repos byLastUpdate) Swap(i int, j int) {
	repos[i], repos[j] = repos[j], repos[i]
}

func (repos byLastUpdate) Less(i int, j int) bool {
	for ci := 0; ci < len(repos[i].LastUpdate) && ci < len(repos[j].LastUpdate); ci++ {
		if repos[i].LastUpdate[ci] > repos[j].LastUpdate[ci] {
			return true
		} else if repos[i].LastUpdate[ci] < repos[j].LastUpdate[ci] {
			return false
		}
	}
	return false
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

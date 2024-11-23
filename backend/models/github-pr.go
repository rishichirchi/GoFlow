package models

type GithubPR struct{
	Url string `json:"url"`
	User struct {
		Login string `json:"login"`
	}
	Body string `json:"body"`
}
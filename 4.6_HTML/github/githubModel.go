package github

import (
	"fmt"
	"time"
)

type User struct {
	AvatarURL  string  `json:"avatar_url"`
	HTMLURL    string  `json:"html_url"`
	ID   	   int
	Login      string
}

type Milestone struct {
	Description  string
	HTMLURL      string  `json:"html_url"`
	ID           int
	State        string
	Title        string
}

type Issue struct {
	Number    int
	HTMLURL   string  `json:"html_url"`
	Title     string
	State     string
	User      *User
	Assignees []*User
	CreatedAt time.Time  `json:"created_at"`
	Body      string
	UpdatedAt time.Time  `json:"updated_at"`
	Milestone *Milestone
}

func getIssueURL(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?state=all", owner, repo)
}

// 回傳是否可以認為他等於給定的 user
func (u *User) Equals(x *User) bool {
	return u.ID == x.ID
}

func (m *Milestone) Equals(x *Milestone) bool {
	return m.ID == x.ID
}

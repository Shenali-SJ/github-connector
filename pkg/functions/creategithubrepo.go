package functions

import (
	"fmt"
	"github-connector/pkg/model"
	"github.com/go-resty/resty/v2"
)

func CreateGithubRepo(gitjsonmodel *model.Requestjsonmodel) (string, error) {

	repoData := model.Createrepodata{
		Name:              "test-repo-connector",
		Autoinit:          true,
		Private:           true,
		Gitignoretemplate: "nanoc",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "token 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").
		SetBody(repoData).
		Post("https://api.github.com/user/repos")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status())

	return resp.Status(), nil
}

package cli

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
)

func openGitRepo() (*git.Repository, error) {
	repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func GetGitHead() (string, error) {
	repo, err := openGitRepo()
	if err != nil {
		return "", err
	}
	return getGitHead(repo)
}

func getGitHead(repo *git.Repository) (string, error) {
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	return firstNChars(ref.Hash().String(), 10), nil
}

func getRepoBaseUrl(repo *git.Repository) (string, error) {
	remotes, err := repo.Remotes()
	if err != nil {
		return "", err
	}
	if len(remotes) == 0 {
		return "", fmt.Errorf("no remotes found")
	}
	remote := remotes[0]
	git_origin := strings.TrimSpace(remote.Config().URLs[0])
	git_origin = regexp.MustCompile(`(?:\.git|/)$`).ReplaceAllString(git_origin, "")                           // remove trailing .git and trailing /
	git_origin = regexp.MustCompile(`\w+(:\w+)?@github\.com`).ReplaceAllString(git_origin, "github.com") // remove usernames and tokens from url
	return git_origin, nil
}

func GetCommitURL() (string, error) {
	repo, err := openGitRepo()
	if err != nil {
		return "", err
	}
	baseUrl, err := getRepoBaseUrl(repo)
	if err != nil {
		return "", err
	}
	head, err := getGitHead(repo)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v/commit/%v", baseUrl, head), nil
}

func GetCompareURL(hash string) (string, error){
	repo, err := openGitRepo()
	if err != nil {
		return "", err
	}
	baseUrl, err := getRepoBaseUrl(repo)
	if err != nil {
		return "", err
	}
	head, err := getGitHead(repo)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v/compare/%v...%v", baseUrl, hash, head), nil
}

func firstNChars(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

package gitclient

import (
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type GitConnectionOption interface {
	apply(*GitClient)
}

type GitUploadOption interface {
	GitConnectionOption
	apply(*GitClient)
}

type optionFunc func(*GitClient)

func (of optionFunc) apply(cfg *GitClient) { of(cfg) }

func OptionAuthor(name, email string) GitConnectionOption {
	return optionFunc(func(cfg *GitClient) {
		cfg.Author = object.Signature{
			Name:  name,
			Email: email,
			When:  time.Now(),
		}
	})
}

func OptionBranch(branch string) GitConnectionOption {
	return optionFunc(func(cfg *GitClient) {
		cfg.Branch = branch
	})
}

func OptionDepth(depth int) GitUploadOption {
	return optionFunc(func(cfg *GitClient) {
		cfg.Depth = depth
	})
}

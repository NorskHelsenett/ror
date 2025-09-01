// Package gitclient provides a simple Git client for updating files in remote repositories using in-memory operations.
//
// It uses go-git and go-billy's memfs to perform all git operations (clone, update, commit, push) without writing to disk.
// This is useful for ephemeral or serverless environments, or when you want to avoid filesystem side effects.
//
// Example usage:
//
//	func main() {
//	    client := &GitClient{
//	        RepoURL: "https://github.com/youruser/yourrepo.git",
//	        Branch:  "main",
//	        Token:   "<your-git-token>",
//	    }
//	    err := client.UpdateFile("path/to/file.txt", "new file content", "Update file.txt via automation")
//	    if err != nil {
//	        panic(err)
//	    }
//	    fmt.Println("File updated and pushed successfully!")
//	}
package gitclient

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/dotse/go-health"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

type GitClient struct {
	RepoURL string
	Branch  string
	Token   string
	Author  object.Signature
}

func NewGitClient(repoURL, branch, token string, author object.Signature) *GitClient {
	return &GitClient{
		RepoURL: repoURL,
		Branch:  branch,
		Token:   token,
		Author:  author,
	}
}

// UpdateFile clones the repo into an in-memory filesystem, updates a file, commits, and pushes the change.
// No files are written to disk; all operations are performed in memory using go-billy/memfs.
//
// Arguments:
//
//	filePath:   Path to the file to update (relative to repo root)
//	newContent: New content to write to the file
//	commitMsg:  Commit message for the change
//
// Returns an error if any git operation fails.
func (c *GitClient) UploadFile(filePath string, newContent []byte, commitMsg string) error {
	// Use in-memory filesystem
	fs := memfs.New()

	repo, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:           c.RepoURL,
		ReferenceName: plumbing.NewBranchReferenceName(c.Branch),
		SingleBranch:  true,
		Auth: &http.BasicAuth{
			Username: "token",
			Password: c.Token,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to clone repo: %w", err)
	}

	// Write new content to the file in memory
	absPath := filepath.Join("/", filePath)
	f, err := fs.Create(absPath)
	if err != nil {
		return fmt.Errorf("failed to create file in memfs: %w", err)
	}
	if _, err := f.Write(newContent); err != nil {
		return fmt.Errorf("failed to write file in memfs: %w", err)
	}
	f.Close()

	// Add the file
	wt, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}
	_, err = wt.Add(filePath)
	if err != nil {
		return fmt.Errorf("failed to add file: %w", err)
	}

	// Commit
	_, err = wt.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  c.Author.Name,
			Email: c.Author.Email,
			When:  time.Now(),
		},
	})
	if err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	// Push
	err = repo.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: "token",
			Password: c.Token,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to push: %w", err)
	}
	return nil
}

// GetFile clones the repo into an in-memory filesystem and returns the contents of the specified file as []byte.
// If the file does not exist, it returns nil and no error.
func (c *GitClient) GetFile(filePath string) ([]byte, error) {
	fs := memfs.New()

	_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:           c.RepoURL,
		ReferenceName: plumbing.NewBranchReferenceName(c.Branch),
		SingleBranch:  true,
		Auth: &http.BasicAuth{
			Username: "token",
			Password: c.Token,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to clone repo: %w", err)
	}

	absPath := filepath.Join("/", filePath)
	f, err := fs.Open(absPath)
	if err != nil {
		// File does not exist
		return nil, nil
	}
	defer f.Close()

	content := make([]byte, 0)
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			content = append(content, buf[:n]...)
		}
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
	}
	return content, nil
}

// CheckConnection attempts to clone the repository to verify access and authentication.
// Returns nil if successful, or an error if the connection/auth fails.
func (c *GitClient) CheckConnection() error {
	fs := memfs.New()
	_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:           c.RepoURL,
		ReferenceName: plumbing.NewBranchReferenceName(c.Branch),
		SingleBranch:  true,
		Auth: &http.BasicAuth{
			Username: "token",
			Password: c.Token,
		},
		Depth: 1, // shallow clone for speed
	})
	if err != nil {
		return err
	}
	return nil
}

// CheckHealth checks the health of the git connection and returns a health check
func (c *GitClient) CheckHealth() []health.Check {
	check := health.Check{}
	if err := c.CheckConnection(); err != nil {
		check.Status = health.StatusFail
		check.Output = fmt.Sprintf("Could not connect to git: %v", err)
	}
	return []health.Check{check}
}

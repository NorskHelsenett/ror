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

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

type GitClient struct {
	RepoURL string
	Branch  string
	Token   string
}

func NewGitClient(repoURL, branch, token string) *GitClient {
	return &GitClient{
		RepoURL: repoURL,
		Branch:  branch,
		Token:   token,
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
func (c *GitClient) UpdateFile(filePath string, newContent []byte, commitMsg string) error {
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
	_, err = wt.Commit(commitMsg, &git.CommitOptions{})
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

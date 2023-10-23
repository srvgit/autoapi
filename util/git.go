package util

import (
	"fmt"
	"os"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func CloneRepo(repoURL, directory string) error {
	fmt.Printf("Cloning %s into %s...\n", repoURL, directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	if err != nil {
		return fmt.Errorf("failed to clone the repository: %v", err)
	}

	fmt.Println("Repository cloned successfully.")
	return nil
}

func CleanUp(directory string) error {
	fmt.Printf("Cleaning up: %s...\n", directory)

	err := os.RemoveAll(directory)
	if err != nil {
		return fmt.Errorf("failed to clean up the directory: %v", err)
	}

	fmt.Println("Clean up successful.")
	return nil
}

func CommitAndPush(directory, commitMessage string) error {
	fmt.Printf("Committing and pushing changes from %s...\n", directory)

	// Open the local repository
	r, err := git.PlainOpen(directory)
	if err != nil {
		return fmt.Errorf("failed to open the repository: %v", err)
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get the worktree: %v", err)
	}

	// Add all changes to the staging area
	fmt.Println("Adding changes to the staging area...")
	_, err = w.Add(".")
	if err != nil {
		return fmt.Errorf("failed to add changes to the staging area: %v", err)
	}

	// Commit the changes
	fmt.Println("Committing changes...")
	_, err = w.Commit(commitMessage, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "srujan valisetti",
			Email: "repgit@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		return fmt.Errorf("failed to commit changes: %v", err)
	}

	// Retrieve username and password from environment variables
	username := os.Getenv("GIT_USERNAME")
	password := os.Getenv("GIT_PASSWORD")

	// Set up authentication
	auth := &http.BasicAuth{
		Username: username,
		Password: password,
	}

	// Push changes
	fmt.Println("Pushing changes...")
	err = r.Push(&git.PushOptions{

		Auth:     auth,
		Progress: os.Stdout,
	})
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Println("Repository already up to date.")
			return nil
		}
		return fmt.Errorf("failed to push changes: %v", err)
	}

	fmt.Println("Changes committed and pushed successfully.")
	return nil
}

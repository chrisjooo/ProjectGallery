package validations

import (
	"ProjectGallery/models"
	"errors"
)

func AccountValidation(u *models.Account) error {
	if u.Username == "" {
		return errors.New("Username is required")
	}
	if u.Password == "" {
		return errors.New("Password is required")
	}
	if u.FullName == "" {
		return errors.New("Full name is required")
	}
	if u.Email == "" {
		return errors.New("Email is required")
	}

	return nil
}

func ProjectValidation(u *models.Project) error {
	if u.Name == "" {
		return errors.New("Project name is required")
	}
	if u.Author == "" {
		return errors.New("Undefined author")
	}

	return nil
}

func VoteValidation(u *models.Vote) error {
	if u.Author == "" {
		return errors.New("Undefined user")
	}
	if u.ProjectId == 0 {
		return errors.New("Undefined project ID")
	}
	return nil
}

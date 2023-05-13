package main

import "github.com/HipolitoBautista/internal/models"

type templateData struct {
	Form           []*models.Form
	Flash          string
	UserPermStatus string
	Comments       []*models.Comments
}

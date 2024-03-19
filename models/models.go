package models

import "github.com/jinzhu/gorm"

type Resource struct {
    gorm.Model
    Name        string
    Description string
    // Add other fields as needed
}
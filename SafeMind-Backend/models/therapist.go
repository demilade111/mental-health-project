package models

import "gorm.io/gorm"

type Therapist struct {
    gorm.Model
    UserID      uint   `gorm:"unique" json:"user_id"` 
    Speciality  string `json:"speciality"`
    ProfilePicture  string `json:"profile_pic"`
    Experience  int    `json:"experience"` 
    License     string `json:"license"`  
    SessionRate float64 `json:"session_rate"` 
    User        User   `gorm:"foreignKey:UserID"` 
}

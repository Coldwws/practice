package models

type Room struct {
    ID          int    `json:"id"`
    Number      int    `json:"number"`
    Type        string `json:"type"`
    Description string `json:"description"`
}

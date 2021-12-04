package entry

import "encoding/json"

type BookInput struct{
    Title string `json:"title" binding:"required"`
    Price json.Number `json:"price" binding:"required,number,gt=0"`
    SubTitle string `json:"sub_title" binding:"required"`
    Description string
    Release_date string
}

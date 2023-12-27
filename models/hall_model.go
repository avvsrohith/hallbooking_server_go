package models

import ( 
// "gorm.io/gorm" 
) 


type Hall struct { 
Name            string      `json:"name"` 
Description     string      `json:"description"`
Capacity        int32        `json:"capacity"`
Rent            int32        `json:"rent"`
Equipment       string         `json:"equipment"` 
}
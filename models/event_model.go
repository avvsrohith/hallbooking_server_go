package models

import ( 
// "gorm.io/gorm" 
) 


type Event struct { 
Person_name            	string      `json:"person_name"` 
Person_contact     		string      `json:"person_contact"`
Date        			int32        `json:"date"`
Event_name            	int32        `json:"event_name"`
Start_time       		string         `json:"start_time"` 
End_time	   			string         `json:"end_time"`
Hall_id	   				string         `json:"hall_id"`
}
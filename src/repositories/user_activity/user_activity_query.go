package useractivity

const (
	GetTotalTodayActivity = `
		SELECT 
			COUNT(*)
		FROM 
			user_activities 
	`
)

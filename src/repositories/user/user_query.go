package user

const (
	GetRecomendUser = `
	SELECT %s 
	FROM 
		users u 
	WHERE 
		u.id NOT IN 
			(SELECT ua.passed_user_id FROM user_activities ua 
				WHERE ua.user_id = ? AND created_at > ? AND created_at < ? AND ua.passed_user_id IS NOT NULL)
		AND
		u.id NOT IN 
			(SELECT ua.liked_user_id FROM user_activities ua 
				WHERE ua.user_id = ? AND created_at > ? AND created_at < ? AND ua.liked_user_id IS NOT NULL) 
		AND 
		u.id NOT IN (?)
		AND u.status = 1
	LIMIT 1
	`
)

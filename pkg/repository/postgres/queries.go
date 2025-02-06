package postgres

const (
	querySignUp     = "INSERT INTO users (user_name, email, enc_password, reg_date) VALUES ($1, $2, $3, $4)"
	querySignIn     = "SELECT id, user_name, enc_password FROM users WHERE email = $1"
	queryGetProfile = "SELECT user_name, email, reg_date FROM users WHERE id = $1"

	queryCreateTask  = "INSERT INTO task (user_id, title, description, is_important, due_date, created_date, is_done) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	queryGetAllTask  = "SELECT title, description, is_important, is_done, due_date, created_date FROM task WHERE user_id = $1"
	queryGetTaskByID = "SELECT user_id, title, description, is_important, is_done, due_date, created_date FROM task WHERE id = $1"
	queryUpdateTask  = "UPDATE task SET title = $1, description = $2, is_important = $3, due_date = $4, is_done = $5 WHERE id = $6"
	queryDeleteTask  = "DELETE FROM task WHERE id = $1"

	queryCreateEvent  = "INSERT INTO events (user_id, name, description, appointed_date) VALUES ($1, $2, $3, $4)"
	queryGetAllEvents = "SELECT name, description, appointed_date FROM events WHERE user_id = $1"
	queryGetEventByID = "SELECT name, description, appointed_date FROM events WHERE id = $1"
	queryUpdateEvent  = "UPDATE events SET name = $1, description = $2, appointed_date = $3 WHERE id = $4"
	queryDeleteEvent  = "DELETE FROM events WHERE id = $1"

	queryGetMyDay = "SELECT task.title, task.description, task.is_important, task.is_done, events.name, events.description FROM task INNER JOIN events ON task.user_id = events.user_id AND task.due_date = events.appointed_date WHERE task.user_id = $1 AND task.due_date = $2"
)

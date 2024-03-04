import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    int
	Name  string
	Posts []Post
}

type Post struct {
	ID      int
	Title   string
	Content string
	UserID  int
}

func main() {
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/*_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		);

		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			Title TEXT
			Content TEXT
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users (id)
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		INSERT INTO users(name) VALUES('John');
		INSERT INTO users(name) VALUES('Alice');

		INSERT INTO posts(title, content, user_id) VALUES('Post 1', 'Content 1', 1);
		INSERT INTO posts(title, content, user_id) VALUES('Post 2', 'Content 2', 1);
		INSERT INTO posts(title, content, user_id) VALUES('Post 3', 'Content 3', 2);
	`)
	if  err != nil {
		log.Println(err)
		return
	}
	*/

	rows, err := db.Query("SELECT id, name FROM users;")
	fmt.Println("SELECT id, name FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	posts := getAllPosts(db)
	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}

		for _,  post := range posts {
			if user.ID == post.UserID {
				user.Posts = append(user.Posts, post)
			}
		}

		users = append(users, user)
	}

	for _, user := range users {
		ftm.Printf("User: %s\n", user.Name)
		for _, post := range user.Posts {
			fmt.Printf("	Post: %s\n", post.Title)
		}
	}
}

func getAllPosts(db *sql.DB) []Post {
	rows, err := db.Query("SELECT id, title, content, user_id;")
	fmt.Println("SELECT id, title, content, user_id FROM posts;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var userPosts []Post

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID)
		if err != nil {
			log.Fatal(err)
		}
		userPosts = append(userPosts, post)
	}

	return userPosts	
}

func getPostsForUser(db *sql.DB, userID int) []Post {
	rows, err := db.Query("SELECT id, title, content, user_id FROM posts WHERE user_id = ?;", userID)
	fmt.Printf("SELECT id, title, content, user_id FROM posts WHERE user_id = %s\n;", strconv.Itoa(userID))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var userPosts []Post

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID)
		if err != nil {
			log.Fatal(err)
		}
		userPosts = append(userPosts, post)
	}

	return userPosts	
}
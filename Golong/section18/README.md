## Udemy Todo App

DBの作成 + Userテーブルの設定、作成

ユーザーの作成（Create）
```
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
```
```
sqlite3 webapp.sql
sqlite> SELECT * FROM users;
1|6ebff1fa-0027-11ee-85de-d22101a468e6|test|test@test.com|51abb9636078defbf888d8457a7c76f85c8f114c|2023-06-01 11:53:05.686254+09:00
```

ユーザーの取得（Get）
```
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select * from users where id=?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}
```

```
{1 6ebff1fa-0027-11ee-85de-d22101a468e6 test test@test.com 51abb9636078defbf888d8457a7c76f85c8f114c 2023-06-01 11:53:05.686254 +0900 +0900}
```

ユーザーの更新（Update）
```
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
```
```
{1 6ebff1fa-0027-11ee-85de-d22101a468e6 test test@test.com 51abb9636078defbf888d8457a7c76f85c8f114c 2023-06-01 11:53:05.686254 +0900 +0900}
{1 6ebff1fa-0027-11ee-85de-d22101a468e6 test2 test2@test.com 51abb9636078defbf888d8457a7c76f85c8f114c 2023-06-01 11:53:05.686254 +0900 +0900}
```
ユーザーの削除（Delete）
```
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
```

```
{1 6ebff1fa-0027-11ee-85de-d22101a468e6 test2 test2@test.com 51abb9636078defbf888d8457a7c76f85c8f114c 2023-06-01 11:53:05.686254 +0900 +0900}
{0     0001-01-01 00:00:00 +0000 UTC}
```

Todoテーブルの作成

Todoの作成（Create）
```
func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
```
```
sqlite> SELECT * FROM todos;
1|First Todo|2|2023-06-01 14:22:21.223805+09:00
```
Todoの取得（シングルセレクト）（Read）
```
func GetTodo(id int) (todo Todo, err error) {
	cmd := `select * from todos
	where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}
```

Todoの取得（マルチセレクト）
```
func GetTodos() (todos []Todo, err error) {
	cmd := `select * from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}
```
```
{1 First Todo 2 2023-06-01 14:22:21.223805 +0900 +0900}
{2 Second Todo 2 2023-06-01 14:41:19.099977 +0900 +0900}
```

Todoの取得（マルチセレクト２）
```
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos
	where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}
```
```
{1 First Todo 2 2023-06-01 14:22:21.223805 +0900 +0900}
{2 Second Todo 2 2023-06-01 14:41:19.099977 +0900 +0900}
```

Todoの更新（Update）

```
func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, user_id = ? 
	where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
```
```

```
Todoの削除（Delete）
```
```
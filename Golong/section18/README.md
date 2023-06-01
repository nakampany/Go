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
		created_at
		) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())
```
```
sqlite3 webapp.sql
sqlite> SELECT * FROM users;
1|6ebff1fa-0027-11ee-85de-d22101a468e6|test|test@test.com|51abb9636078defbf888d8457a7c76f85c8f114c|2023-06-01 11:53:05.686254+09:00
```

ユーザーの取得（Read）

ユーザーの更新（Update）

ユーザーの削除（Delete)



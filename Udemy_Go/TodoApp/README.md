## todoApp

## 使用技術
- Docker
- postgres
- go


## docker-compose コマンド
- build
  ```
  docker-compose build
  ```
- コンテナ起動
  ```
  docker-compose up -d
  ```
- app
  ```
  docker-compose exec app sh
  ```
- postgres
  ```
  docker-compose exec postgres bash
  ```

- docker down
  ```
  docker-compose down
  ```


### postgresのデータベース確認
- postgres
  ```
  docker-compose exec postgres bash
  ```
- データベースに接続
  ```sql
  // （psql -d データベース名 -U ユーザ名）
  psql -d postgres -U postgres


  SELECT * FROM users;
  ```
- postgres
  ```
  docker-compose exec postgres bash
  ```


### 参考

- uuid
https://github.com/google/uuid
データベースを扱う場合、テーブルの各行に対して一意の識別子を与えるために、ある種のidフィールドを使用する

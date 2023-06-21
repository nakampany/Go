FROM postgres:latest

# /docker-entrypoint-initdb.d/ 内に sql ファイルを配置すると、それらを自動で実行してくれるらしい
COPY ./go/database/init/*.sql /docker-entrypoint-initdb.d/

RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV LANG ja_JP.utf8


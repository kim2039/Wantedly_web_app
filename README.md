# Wantedlyインターン選考課題
Dockerコンテナ上で，"Hello world!"というメッセージをJSON形式で返すAPIです．

## How To
クローンしたリポジトリ直下の'Dockerfile'からイメージを作成する．
~~~
$ sudo docker build -t hello:1.0 .
~~~
<br>

作成したイメージを元にコンテナを起動．<br>
~~~
$ sudo docker run -p 8080:80 hello:1.0
~~~

<br>

そしてリクエストを送ると<br>
~~~
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
~~~
<br>

~~~
{
  "message": "Hello World!!"
}
~~~
<br>
が返ってくる．

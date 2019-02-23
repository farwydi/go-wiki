# go-wiki

Для того, чтобы роутить необходимо поднять [реверс прокси](https://github.com/jwilder/nginx-proxy):

`docker run -d -p 80:80 -v /var/run/docker.sock:/tmp/docker.sock:ro jwilder/nginx-proxy`

И подключится к сети сека.
Также прописать DNS:
> 127.0.0.1 api.go-wiki.origin
>
> 127.0.0.1 go-wiki.origin
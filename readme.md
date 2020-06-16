Setup webdav for zotero attachments sync


```bash
$ docker build . -t yanyaoer/caddy2-webdav
$ docker run -it --rm -p 8881:8080 -v /etc/caddy/:/data-caddy/ -v /data/webdav:/data yanyaoer/caddy2-webdav:latest
```

refer:

<https://www.digitalocean.com/community/tutorials/how-to-remotely-access-gui-applications-using-docker-and-caddy-on-ubuntu-18-04>

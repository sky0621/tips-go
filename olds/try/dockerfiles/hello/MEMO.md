# dockerhub

<pre>
$ sudo docker login
Authenticating with existing credentials...
WARNING! Your password will be stored unencrypted in /home/sky0621/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
</pre>

<pre>
$ sudo docker build . -t sky0621dhub/golanghello
</pre>

<pre>
$ sudo docker run sky0621dhub/golanghello
Hello
</pre>

<pre>
$ sudo docker push sky0621dhub/golanghello
The push refers to repository [docker.io/sky0621dhub/golanghello]
06685027e9f1: Pushed 
latest: digest: sha256:e88b9b8a8b5a3a082b1c547ddd7148b68bb0faabeec2f9a8bc26d735c92b3003 size: 528
</pre>

<pre>
$ sudo docker pull sky0621dhub/golanghello
Using default tag: latest
latest: Pulling from sky0621dhub/golanghello
Digest: sha256:e88b9b8a8b5a3a082b1c547ddd7148b68bb0faabeec2f9a8bc26d735c92b3003
Status: Image is up to date for sky0621dhub/golanghello:latest
docker.io/sky0621dhub/golanghello:latest
</pre>

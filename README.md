# Tweet Image Viewer

## How to setup

```shell
npm i -g svelte-cli # if you don't install svelte-cli
svelte compile --format iife view/Time.html > ./assets/index.js
```

## How to use

```shell
$ export TWITTER_CKEY=CONSUMER_KEY
$ export TWITTER_CSECRET=CONSUMER_SECRET
$ export TWITTER_AKEY=ACCESS_TOKEN
$ export TWITTER_ASECRET=ACCESS_SECRET
$ go run main.go
```

And then, open localhost:8000 and input tweet id into form. You'll get the id's images. You can also download them if you click them.

# LICENSE
MIT

package main

import (
        "net/http"
        "strconv"
        "os"
        "encoding/json"
        "regexp"
        "html/template"
        "log"

        "github.com/zenazn/goji"
        "github.com/zenazn/goji/web"
        "github.com/ChimeraCoder/anaconda"
)

func media(c web.C, w http.ResponseWriter, r *http.Request) {
        cKey := os.Getenv("TWITTER_CKEY")
        cSecret := os.Getenv("TWITTER_CSECRET")
        aKey := os.Getenv("TWITTER_AKEY")
        aSecret := os.Getenv("TWITTER_ASECRET")
        anaconda.SetConsumerKey(cKey)
        anaconda.SetConsumerSecret(cSecret)
        api := anaconda.NewTwitterApi(aKey, aSecret)

        reg := regexp.MustCompile(`[^0-9]+`)
        id := c.URLParams["id"]
        if reg.MatchString(id) {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("Bad Request"))
            return
        }

        tweetId, _ := strconv.ParseInt(id, 10, 64)
        tweet, err := api.GetTweet(tweetId, nil)

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }

        urls := []string{}
        if tweet.ExtendedEntities.Media != nil {
            for _, media := range tweet.ExtendedEntities.Media {
                urls = append(urls, media.Media_url_https)
            }
        }

        encoder := json.NewEncoder(w)
        encoder.Encode(urls)
}

func index(c web.C, w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("view/index.html"))

    if err := t.ExecuteTemplate(w, "index.html", ""); err != nil {
        log.Fatal(err)
    }
}

func main() {
        staticPattern := regexp.MustCompile("/assets/*")
        goji.Handle(staticPattern, http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

        goji.Get("/twitter/:id", media)
        goji.Get("/", index)
        goji.Serve()
}

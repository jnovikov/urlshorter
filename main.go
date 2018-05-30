package main

import (
	"net/http"
	"fmt"
	"urlshorter/url"
	"github.com/spf13/viper"
	"urlshorter/redis"
	"urlshorter/keystorage"
	"urlshorter/lru"
)

var generator url.Generator

func BuildUrl(relative string) string {
	fqdn := viper.GetString("fqdn")
	port := viper.GetString("port")
	return url.Absolute(fqdn, port, relative)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	short := url.ParsePath(r)
	res := generator.Find(short)
	if res == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}
	http.Redirect(w, r, res, 302)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	sourceUrl := r.FormValue("url")
	if sourceUrl == "" {
		w.Write([]byte("Short url"))
		return
	}
	result := generator.New(sourceUrl)
	w.Write([]byte(BuildUrl(result)))

}

func main() {
	// Read config from file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetDefault("host", "")
	viper.SetDefault("port", "8000")
	viper.SetDefault("url_size", 6)
	viper.SetDefault("fqdn", "localhost")
	viper.SetDefault("redis_host", "localhost")
	viper.SetDefault("redis_port", "6379")
	viper.SetDefault("pool_size", 10)
	viper.SetDefault("cache_size", 5000)
	viper.ReadInConfig()

	//Create redis connection pool

	pool := redis.NewPool(viper.GetString("redis_host"), viper.GetString("redis_port"), viper.GetInt("pool_size"))

	factory := redis.NewFactory(pool)

	// Create storage

	urlSt := factory.HStorage("url_to_short")
	shortSt := factory.HStorage("short_to_url")

	//Create new Generator
	g := url.NewRandomGenerator()
	g.SetUrlSize(viper.GetInt("url_size"))

	g.SetShortStorage(keystorage.NewCacheStorage(shortSt, lru.Storage(viper.GetInt("cache_size"))))
	g.SetUrlStorage(keystorage.NewCacheStorage(urlSt, lru.Storage(viper.GetInt("cache_size"))))
	g.SetUrlSize(viper.GetInt("url_size"))

	generator = g
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", addHandler)
	serveUrl := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))
	fmt.Println("Serving at ", serveUrl)
	http.ListenAndServe(serveUrl, nil)
}

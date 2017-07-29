package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"math/rand"
	"sync"
	"time"
)

// JSON-parsed format of an article.
type Article struct {
	Author			string
	Title			string
	Description		string
	Url				string
	UrlToImage		string
	PublishedAt		string
	// Custom parameters:
	NewsSourceId	string
	Host			string
	Category		string
	Language		string
	Country			string
}

// For rendering the news article information.
type ArticleArg struct {
	Article
	Index			int
}

// A news source to request the news from.
type NewsSource struct {
	Id					string
	Name				string
	Description			string
	Url					string
	Category			string
	Language			string
	Country				string
	SortBysAvailable	[]string
	// Custom parameters:
	Icon				string
}
type NewsSources map[string]NewsSource

var (
	// newsServer populates the articles.
	newsServerRunning = false
	mutex = &sync.RWMutex{}
	articles []Article
	
	// Custom-written data from https://newsapi.org/v1/sources?language=en query
	newsSources NewsSources

	// News source icons no longer part of API, so have to set manually.
	newsSourceIcons map[string]string = map[string]string{
		"abc-news-au": "https://icons.better-idea.org/icon?url=http://www.abc.net.au/news&size=70..120..200",
		"al-jazeera-english": "https://icons.better-idea.org/icon?url=http://www.aljazeera.com&size=70..120..200",
		"ars-technica": "https://icons.better-idea.org/icon?url=http://arstechnica.com&size=70..120..200",
		"associated-press": "https://icons.better-idea.org/icon?url=https://apnews.com/&size=70..120..200",
		"bbc-news": "https://icons.better-idea.org/icon?url=http://www.bbc.co.uk/news&size=70..120..200",
		"bbc-sport": "https://icons.better-idea.org/icon?url=http://www.bbc.co.uk/sport&size=70..120..200",
		"bloomberg": "https://icons.better-idea.org/icon?url=http://www.bloomberg.com&size=70..120..200",
		"breitbart-news": "https://icons.better-idea.org/icon?url=http://www.breitbart.com&size=70..120..200",
		"business-insider": "https://icons.better-idea.org/icon?url=http://www.businessinsider.com&size=70..120..200",
		"business-insider-uk": "https://icons.better-idea.org/icon?url=http://uk.businessinsider.com&size=70..120..200",
		"buzzfeed": "https://icons.better-idea.org/icon?url=https://www.buzzfeed.com&size=70..120..200",
		"cnbc": "https://icons.better-idea.org/icon?url=http://www.cnbc.com&size=70..120..200",
		"cnn": "https://icons.better-idea.org/icon?url=http://us.cnn.com&size=70..120..200",
		"daily-mail": "https://icons.better-idea.org/icon?url=http://www.dailymail.co.uk/home/index.html&size=70..120..200",
		"der-tagesspiegel": "https://icons.better-idea.org/icon?url=http://www.tagesspiegel.de&size=70..120..200",
		"die-zeit": "https://icons.better-idea.org/icon?url=http://www.zeit.de/index&size=70..120..200",
		"engadget": "https://icons.better-idea.org/icon?url=https://www.engadget.com&size=70..120..200",
		"entertainment-weekly": "https://icons.better-idea.org/icon?url=http://www.ew.com&size=70..120..200",
		"espn": "https://icons.better-idea.org/icon?url=http://espn.go.com&size=70..120..200",
		"espn-cric-info": "https://icons.better-idea.org/icon?url=http://www.espncricinfo.com/&size=70..120..200",
		"financial-times": "https://icons.better-idea.org/icon?url=http://www.ft.com/home/uk&size=70..120..200",
		"focus": "https://icons.better-idea.org/icon?url=http://www.focus.de&size=70..120..200",
		"football-italia": "https://icons.better-idea.org/icon?url=http://www.football-italia.net&size=70..120..200",
		"fortune": "https://icons.better-idea.org/icon?url=http://fortune.com&size=70..120..200",
		"four-four-two": "https://icons.better-idea.org/icon?url=http://www.fourfourtwo.com/news&size=70..120..200",
		"fox-sports": "https://icons.better-idea.org/icon?url=http://www.foxsports.com&size=70..120..200",
		"google-news": "https://icons.better-idea.org/icon?url=https://news.google.com&size=70..120..200",
		"gruenderszene": "https://icons.better-idea.org/icon?url=http://www.gruenderszene.de&size=70..120..200",
		"hacker-news": "https://icons.better-idea.org/icon?url=https://news.ycombinator.com&size=70..120..200",
		"handelsblatt": "https://icons.better-idea.org/icon?url=http://www.handelsblatt.com&size=70..120..200",
		"ign": "https://icons.better-idea.org/icon?url=http://www.ign.com&size=70..120..200",
		"independent": "https://icons.better-idea.org/icon?url=http://www.independent.co.uk&size=70..120..200",
		"mashable": "https://icons.better-idea.org/icon?url=http://mashable.com&size=70..120..200",
		"metro": "https://icons.better-idea.org/icon?url=http://metro.co.uk&size=70..120..200",
		"mirror": "https://icons.better-idea.org/icon?url=http://www.mirror.co.uk/&size=70..120..200",
		"mtv-news": "https://icons.better-idea.org/icon?url=http://www.mtv.com/news&size=70..120..200",
		"mtv-news-uk": "https://icons.better-idea.org/icon?url=http://www.mtv.co.uk/news&size=70..120..200",
		"national-geographic": "https://icons.better-idea.org/icon?url=http://news.nationalgeographic.com&size=70..120..200",
		"new-scientist": "https://icons.better-idea.org/icon?url=https://www.newscientist.com/section/news&size=70..120..200",
		"newsweek": "https://icons.better-idea.org/icon?url=http://www.newsweek.com&size=70..120..200",
		"new-york-magazine": "https://icons.better-idea.org/icon?url=http://nymag.com&size=70..120..200",
		"nfl-news": "https://icons.better-idea.org/icon?url=http://www.nfl.com/news&size=70..120..200",
		"polygon": "https://icons.better-idea.org/icon?url=http://www.polygon.com&size=70..120..200",
		"recode": "https://icons.better-idea.org/icon?url=http://www.recode.net&size=70..120..200",
		"reddit-r-all": "https://icons.better-idea.org/icon?url=https://www.reddit.com/r/all&size=70..120..200",
		"reuters": "https://icons.better-idea.org/icon?url=http://www.reuters.com&size=70..120..200",
		"spiegel-online": "https://icons.better-idea.org/icon?url=http://www.spiegel.de&size=70..120..200",
		"t3n": "https://icons.better-idea.org/icon?url=http://t3n.de&size=70..120..200",
		"talksport": "https://icons.better-idea.org/icon?url=http://talksport.com&size=70..120..200",
		"techcrunch": "https://icons.better-idea.org/icon?url=https://techcrunch.com&size=70..120..200",
		"techradar": "https://icons.better-idea.org/icon?url=http://www.techradar.com&size=70..120..200",
		"the-economist": "https://icons.better-idea.org/icon?url=http://www.economist.com&size=70..120..200",
		"the-guardian-au": "https://icons.better-idea.org/icon?url=https://www.theguardian.com/au&size=70..120..200",
		"the-guardian-uk": "https://icons.better-idea.org/icon?url=https://www.theguardian.com/uk&size=70..120..200",
		"the-hindu": "https://icons.better-idea.org/icon?url=http://www.thehindu.com&size=70..120..200",
		"the-huffington-post": "https://icons.better-idea.org/icon?url=http://www.huffingtonpost.com&size=70..120..200",
		"the-lad-bible": "https://icons.better-idea.org/icon?url=http://www.theladbible.com&size=70..120..200",
		"the-new-york-times": "https://icons.better-idea.org/icon?url=http://www.nytimes.com&size=70..120..200",
		"the-next-web": "https://icons.better-idea.org/icon?url=http://thenextweb.com&size=70..120..200",
		"the-sport-bible": "https://icons.better-idea.org/icon?url=http://www.thesportbible.com&size=70..120..200",
		"the-telegraph": "https://icons.better-idea.org/icon?url=http://www.telegraph.co.uk&size=70..120..200",
		"the-times-of-india": "https://icons.better-idea.org/icon?url=http://timesofindia.indiatimes.com&size=70..120..200",
		"the-verge": "https://icons.better-idea.org/icon?url=http://www.theverge.com&size=70..120..200",
		"the-wall-street-journal": "https://icons.better-idea.org/icon?url=http://www.wsj.com&size=70..120..200",
		"the-washington-post": "https://icons.better-idea.org/icon?url=https://www.washingtonpost.com&size=70..120..200",
		"time": "https://icons.better-idea.org/icon?url=http://time.com&size=70..120..200",
		"usa-today": "https://icons.better-idea.org/icon?url=http://www.usatoday.com/news&size=70..120..200",
		"wired-de": "https://icons.better-idea.org/icon?url=https://www.wired.de&size=70..120..200",
		"wirtschafts-woche": "https://icons.better-idea.org/icon?url=http://www.wiwo.de&size=70..120..200",
	}
)

// TODO: implement news debug output functions
func newsPrint(s string) {
}

func newsPrintVal(s string, val interface{}) {
}

func newsPrintf(format string, args... interface{}) {
}

//////////////////////////////////////////////////////////////////////////////
//
// fetches news sources
//
//////////////////////////////////////////////////////////////////////////////
func fetchNewsSources() bool {
	print("fetchNewsSources")
	
	// Note: I should be passing in category, language, and country parameters.
	newsRequestUrl := "https://newsapi.org/v1/sources"
	newsRequestUrl += "?apiKey=" + flags.newsAPIKey
	
	newsPrintVal("newsRequestUrl", newsRequestUrl)
	
	resp, err := httpGet(newsRequestUrl)
	check(err)
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	
	newsPrintVal("body", string(body))
	
	// Parse the News Sources json.
	var newsSourcesResp struct {
		Status	string
		Sources	[]NewsSource
	}
	err = json.Unmarshal(body, &newsSourcesResp)
	check(err)
	
	// News request returned an error.
	if newsSourcesResp.Status != "ok" {
		fmt.Printf("Error fetching news sources: '%s'\n", body)
		return false
	}
	
	// Copy news source data to newsSources.
	newsSources = NewsSources{}
	for _, newsSource := range newsSourcesResp.Sources {
		newsPrintVal("newsSource", newsSource)
		newsSources[newsSource.Id] = newsSource	
	}
	
	// Set news source icons.  (No longer part of API, so have to set manually.)
	for source, icon := range newsSourceIcons {
		tmp := newsSources[source]
		tmp.Icon = icon
		newsSources[source] = tmp
	}
	
	return true
}

//////////////////////////////////////////////////////////////////////////////
//
// fetches news articles from a single source
//
//////////////////////////////////////////////////////////////////////////////
func fetchNews(newsSource string, c chan []Article) {
	// Note: I should be passing in category, language, and country parameters.
	newsRequestUrl := "https://newsapi.org/v1/articles"
	//newsRequestUrl += "?sortBy=latest"
	newsRequestUrl += "?apiKey=" + flags.newsAPIKey
	newsRequestUrl += "&source=" + newsSource
	
	newsPrintVal("newsRequestUrl", newsRequestUrl)
	
	resp, err := httpGet(newsRequestUrl)
	if err != nil {
		fmt.Printf("Error fetching news from '%s': '%s'\n", newsSource, err)
		c <- []Article{}
		return
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error fetching news from '%s': '%s'\n", newsSource, err)
		c <- []Article{}
		return
	}
	
	// Parse the News API json.
	var news struct {
		Status		string
		Source		string
		SortBy		string
		Articles	[]Article
	}
	err = json.Unmarshal(body, &news)
	if err != nil {
		fmt.Printf("Error fetching news from '%s': '%s' '%s'\n", newsSource, err, body)
		c <- []Article{}
		return
	}
	
	// News request returned an error.
	if news.Status != "ok" {
		fmt.Printf("Error fetching news from '%s': '%s'\n", newsSource, body)
		c <- []Article{}
		return
	}

	for i := 0; i < len(news.Articles); i++ {
		// Set the news source
		news.Articles[i].NewsSourceId = newsSource
		
		// Parse the hostname.  TODO: parse away the "www."
		u, err := url.Parse(news.Articles[i].Url)
		if err != nil {
			news.Articles[i].Host = "Error parsing hostname"
		} else {
			news.Articles[i].Host = u.Host
		}
		
		// Set the category, language, and country.
		news.Articles[i].Category = newsSources[newsSource].Category
		news.Articles[i].Language = newsSources[newsSource].Language
		news.Articles[i].Country  = newsSources[newsSource].Country
	}
	
	c <- news.Articles
}

//////////////////////////////////////////////////////////////////////////////
//
// news server - On startup, and every 5 minutes, fetches the latest news.
//
//////////////////////////////////////////////////////////////////////////////
func newsServer() {
	newsServerRunning = true
	defer func(){newsServerRunning = false}()
	
	print("newsServer - fetchNewsSources")
	ok := fetchNewsSources()
	if !ok {
		return
	}
	
	for {
		newsPrint("========================================")
		newsPrint("============ FETCHING NEWS =============")
		newsPrint("========================================\n")
		
		c := make(chan []Article)
		timeout := time.After(5 * time.Second)
		
		newsPrintVal("len(newsSources)", len(newsSources))
		
		for _, newsSource := range newsSources {
			newsPrintVal("Fetching article from", newsSource.Id)
			
			go fetchNews(newsSource.Id, c)
		}
		
		newArticles := []Article{}
		numSourcesFetched := 0
		fetchNewsLoop: for {
			select {
				case newArticlesFetched := <-c:
					newArticles = append(newArticles, newArticlesFetched...)
					numSourcesFetched++
					
					newsPrintVal("New articles fetched", numSourcesFetched)
				case <- timeout:
					newsPrint("Timeout!")
					break fetchNewsLoop
			}
		}
	
		print("Copying new articles")
		mutex.Lock()
		articles = newArticles
		mutex.Unlock()
		print("New articles copied")
	
		print("Sleeping 5 minutes")
		time.Sleep(5 * time.Minute)
	}
}

//////////////////////////////////////////////////////////////////////////////
//
// display news
// TODO: santize (html- and url-escape the arguments).  (Make sure URL's don't point back to votezilla.)
// TODO: use a caching, resizing image proxy for the images.
//
//////////////////////////////////////////////////////////////////////////////
func newsHandler(w http.ResponseWriter, r *http.Request) {
	if (!newsServerRunning) {
		go newsServer()
		time.Sleep(2 * time.Second)
	}

	numArticlesToDisplay := min(50, len(articles))
	
	articleArgs := make([]ArticleArg, numArticlesToDisplay)
	
	perm := rand.Perm(len(articles))
	
	mutex.RLock()
	for i := 0; i < numArticlesToDisplay; i++ {
		article := articles[perm[i]] // shuffle the article order (to mix between sources)

		// Copy the article information.
		articleArgs[i].Article = article

		// Set the index
		articleArgs[i].Index = i + 1
	}
	mutex.RUnlock()

	// Get the username.
	userId := GetSession(r)
	username := getUsername(userId)

	// Render the news articles.
	newsArgs := struct {
		PageArgs
		Username	string
		Articles	[]ArticleArg
	}{
		PageArgs: PageArgs{Title: "votezilla - News"},
		Username: username,
		Articles: articleArgs,
	}
	executeTemplate(w, "news", newsArgs)
}


///////////////////////////////////////////////////////////////////////////////
//
// display news sources - TODO: checkboxes so user can pick 
//                        which news sources they want to see.
//
///////////////////////////////////////////////////////////////////////////////
func newsSourcesHandler(w http.ResponseWriter, r *http.Request) {
	if (!newsServerRunning) {
		go newsServer()
		time.Sleep(2 * time.Second)
	}
	
	newsSourcesArgs := struct {
		PageArgs
		NewsSources NewsSources
	}{
		PageArgs: PageArgs{Title: "News Sources"},
		NewsSources: newsSources,
	}
	fmt.Println("newsSourcesArgs: %#v\n", newsSourcesArgs)
	executeTemplate(w, "newsSources", newsSourcesArgs)	
}

///////////////////////////////////////////////////////////////////////////////
//
// init news - starts the news server
//
///////////////////////////////////////////////////////////////////////////////
func InitNews() {
	if (!newsServerRunning) {
		go newsServer()
	}	
}
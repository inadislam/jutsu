package controllers

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	models "github.com/inadislam/jutsu/app/models/anime"
	"github.com/inadislam/jutsu/pkg/utils"
)

var (
	default_url = "anime/"
)

type Info struct {
	Type, Source, Status, Duration, Rating, Synopsis, Background, Season string
	Episodes, ScoredBy, Rank, Popularity, Members, Favorites, Year       int
	Airing                                                               bool
	Aired                                                                models.Aired
	Score                                                                float64
	Broadcast                                                            []models.Broadcast
	Producers                                                            []models.Producers
	Licensors                                                            []models.Licensors
	Studios                                                              []models.Studios
	Genres                                                               []models.Genres
	ExplicitGenres                                                       []models.Genres
	Themes                                                               []models.Themes
	Demographics                                                         []string
}

func AnimeHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	anime_url := GetUrl(id)
	anime_id := utils.GetMalId(anime_url)
	images := GetImages(anime_url)
	trailer := GetTrailer(anime_url)
	title, title_english, title_japanese, title_synonyms := GetTitle(anime_url)
	info := GetInformation(anime_url)
	var (
		inf        Info
		exgen, dem []interface{}
	)
	for _, k := range info {
		inf = k
	}
	data := models.Anime{
		MalID:          anime_id,
		URL:            anime_url,
		Images:         images,
		Trailer:        trailer,
		Title:          title,
		TitleEnglish:   title_english,
		TitleJapanese:  title_japanese,
		TitleSynonyms:  title_synonyms,
		Type:           inf.Type,
		Source:         inf.Source,
		Episodes:       inf.Episodes,
		Status:         inf.Status,
		Airing:         inf.Airing,
		Aired:          inf.Aired,
		Duration:       inf.Duration,
		Rating:         inf.Rating,
		Score:          inf.Score,
		ScoredBy:       inf.ScoredBy,
		Rank:           inf.Rank,
		Popularity:     inf.Popularity,
		Members:        inf.Members,
		Favorites:      inf.Favorites,
		Synopsis:       inf.Synopsis,
		Background:     inf.Background,
		Season:         inf.Season,
		Year:           inf.Year,
		Broadcast:      inf.Broadcast,
		Producers:      inf.Producers,
		Licensors:      inf.Licensors,
		Studios:        inf.Studios,
		Genres:         inf.Genres,
		ExplicitGenres: exgen,
		Themes:         inf.Themes,
		Demographics:   dem,
	}

	if data.URL == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"type":    http.StatusText(400),
			"message": "Resource does not exist",
			"error":   "404 on " + os.Getenv("BASE_URL") + "anime/" + id,
		})
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

func GetUrl(id string) string {
	var (
		anime_url string
	)
	h := colly.NewCollector()
	h.OnHTML("body", func(e *colly.HTMLElement) {
		es := e.DOM
		anime_url, _ = es.Find(".breadcrumb .di-ib:nth-child(3) a").Attr("href")
	})
	h.Visit(os.Getenv("BASE_URL") + default_url + id)
	return anime_url
}

func GetImages(anime_url string) []models.Images {
	var im []models.Images
	h := colly.NewCollector()
	h.OnHTML(".leftside > div:first-child", func(e *colly.HTMLElement) {
		es := e.DOM
		jpg_url, _ := es.Find("a > img").Attr("data-src")
		sju := strings.Split(jpg_url, ".jpg")
		webp_url := strings.Join(sju, ".webp")
		tmp := models.Images{
			Jpg: models.Jpg{
				ImageURL:      jpg_url,
				SmallImageURL: strings.Join(sju, "t.jpg"),
				LargeImageURL: strings.Join(sju, "l.jpg"),
			},
			Webp: models.Webp{
				ImageURL:      webp_url,
				SmallImageURL: strings.Join(sju, "t.webp"),
				LargeImageURL: strings.Join(sju, "l.webp"),
			},
		}
		im = append(im, tmp)
	})
	h.Visit(anime_url)
	return im
}

func GetTrailer(anime_url string) []models.Trailer {
	var tr []models.Trailer
	h := colly.NewCollector()
	h.OnHTML(".video-promotion", func(e *colly.HTMLElement) {
		es := e.DOM
		embed_url, _ := es.Find("a").Attr("href")
		youtube_id := strings.Split(strings.Split(embed_url, "?")[0], "embed/")[1]
		url := "https://www.youtube.com/watch?v=" + youtube_id
		defaultImageUrl := "https://img.youtube.com/vi/"
		tmp := models.Trailer{
			YoutubeID: youtube_id,
			URL:       url,
			EmbedURL:  embed_url,
			TrailerImages: models.TrailerImages{
				ImageURL:        defaultImageUrl + youtube_id + "/default.jpg",
				SmallImageURL:   defaultImageUrl + youtube_id + "/sddefault.jpg",
				MediumImageURL:  defaultImageUrl + youtube_id + "/mqdefault.jpg",
				LargeImageURL:   defaultImageUrl + youtube_id + "/hqdefault.jpg",
				MaximumImageURL: defaultImageUrl + youtube_id + "/maxresdefault.jpg",
			},
		}
		tr = append(tr, tmp)
	})
	h.Visit(anime_url)
	return tr
}

func GetTitle(anime_url string) (string, string, string, []string) {
	var (
		title, title_english, title_japanese string
		title_synonyms                       []string
	)
	h := colly.NewCollector()
	h.OnHTML("body", func(e *colly.HTMLElement) {
		es := e.DOM
		title = es.Find("h1.title-name > strong").Text()
		title_english = title
		title_japanese = strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Japanese:")`).Parent().Text(), "Japanese: "), ""))
		title_synonyms = strings.Split(strings.Join(strings.Split(es.Find(`span:contains("Synonyms:")`).Parent().Text(), "Synonyms:"), ""), ",")
	})
	h.Visit(anime_url)
	return title, title_english, title_japanese, title_synonyms
}

func GetInformation(anime_url string) []Info {
	var inf []Info
	h := colly.NewCollector()
	h.OnHTML("body", func(e *colly.HTMLElement) {
		es := e.DOM
		Type := strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Type:")`).Parent().Text(), "Type:"), ""))
		source := strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Source:")`).Parent().Text(), "Source:"), ""))
		episodes, _ := strconv.Atoi(strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Episodes:")`).Parent().Text(), "Episodes:"), "")))
		status := strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Status:")`).Parent().Text(), "Status:"), ""))
		var airing bool
		if status == "Finished Airing" {
			airing = false
		} else {
			airing = true
		}
		air_string := strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Aired:")`).Parent().Text(), "Aired:"), ""))
		var air_from, air_to string
		if strings.Contains(air_string, "to") {
			for i, k := range strings.Split(air_string, "to") {
				if i == 0 {
					air_from = string(k)
				}
				if i == 1 {
					air_to = string(k)
				}
			}
		} else {
			air_from = air_string
		}
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		pf := strings.Split(air_from, ",")
		var pfd, pfm, pfy int
		for i, k := range pf {
			if i == 0 {
				pfd, _ = strconv.Atoi(strings.Join(re.FindAllString(string(k), -1), ""))
				for j, jm := range strings.Split(string(k), " ") {
					if j == 0 {
						pfm = utils.GetMonthNumber(jm)
					}
				}
			}
			if i == 1 {
				pfy, _ = strconv.Atoi(strings.TrimSpace(string(k)))
			}
		}
		pt := strings.Split(air_to, ",")
		var ptd, ptm, pty int
		for i, k := range pt {
			if i == 0 {
				ptd, _ = strconv.Atoi(strings.Join(re.FindAllString(string(k), -1), ""))
				for j, jm := range strings.Split(string(k), " ") {
					if j == 1 {
						ptm = utils.GetMonthNumber(jm)
					}
				}
			}
			if i == 1 {
				pty, _ = strconv.Atoi(strings.TrimSpace(string(k)))
			}
		}
		duration := strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Duration:")`).Parent().Text(), "Duration:"), ""))
		rating := strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Rating:")`).Parent().Text(), "Rating:"), ""))
		year := pfy
		score, _ := strconv.ParseFloat(strings.TrimSpace(es.Find("span.score-label").Text()), 64)
		scoredBy, _ := strconv.Atoi(es.Find(`[itemprop="ratingCount"]`).Text())
		h, _ := es.Find(`span:contains("Ranked:")`).Parent().Html()
		var rank int
		for i, k := range strings.Split(h, "#") {
			if i == 1 {
				for j, jm := range strings.Split(k, "<sup") {
					if j == 0 {
						rank, _ = strconv.Atoi(jm)
					}
				}
			}
		}
		p, _ := es.Find(`span:contains("Popularity:")`).Parent().Html()
		var popularity int
		for i, k := range strings.Split(p, "#") {
			if i == 1 {
				popularity, _ = strconv.Atoi(strings.TrimSpace(k))
			}
		}
		members, _ := strconv.Atoi(strings.Replace(strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Members:")`).Parent().Text(), "Members:"), "")), ",", "", -1))
		favorites, _ := strconv.Atoi(strings.Replace(strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Favorites:")`).Parent().Text(), "Favorites:"), "")), ",", "", -1))
		broadcast := GetBroadcast(es)
		producers := GetProducers(es)
		licensors := GetLicensor(es)
		studios := GetStudios(es)
		genres := GetGenre(es)
		var exgen []models.Genres
		themes := GetThemes(es)
		season := es.Find(".season a").Text()
		desc := es.Find(`[itemprop="description"]`).Text()
		bg := strings.TrimSpace(es.Find(".rightside").Find(`[valign="top"]`).Clone().Children().Remove().End().Text())
		tmp := Info{
			Type:     Type,
			Source:   source,
			Episodes: episodes,
			Status:   status,
			Airing:   airing,
			Aired: models.Aired{
				From: strings.TrimSpace(air_from),
				To:   strings.TrimSpace(air_to),
				Prop: models.Prop{
					From: models.From{
						Day:   pfd,
						Month: pfm,
						Year:  pfy,
					},
					To: models.To{
						Day:   ptd,
						Month: ptm,
						Year:  pty,
					},
				},
				String: air_string,
			},
			Duration:       duration,
			Rating:         rating,
			Score:          score,
			ScoredBy:       scoredBy,
			Rank:           rank,
			Popularity:     popularity,
			Members:        members,
			Favorites:      favorites,
			Synopsis:       desc,
			Background:     bg,
			Season:         season,
			Broadcast:      broadcast,
			Producers:      producers,
			Licensors:      licensors,
			Studios:        studios,
			Genres:         genres,
			ExplicitGenres: exgen,
			Themes:         themes,
			Demographics:   []string{},
			Year:           year,
		}
		inf = append(inf, tmp)
	})
	h.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	h.Visit(anime_url)
	return inf
}

func GetBroadcast(es *goquery.Selection) []models.Broadcast {
	var broadcast []models.Broadcast
	full_string := strings.TrimSpace(strings.Join(strings.Split(es.Find(`span:contains("Broadcast:")`).Parent().Text(), "Broadcast:"), ""))
	var bday, btime, btz string
	for i, k := range strings.Split(full_string, " ") {
		if i == 0 {
			bday = k
		}
		if i == 2 {
			btime = k
		}
		if i == 3 {
			btz = "Asia/Tokyo"
		}
	}
	tmp := models.Broadcast{
		Day:      bday,
		Time:     btime,
		Timezone: btz,
		String:   full_string,
	}
	broadcast = append(broadcast, tmp)
	return broadcast
}

func GetProducers(es *goquery.Selection) []models.Producers {
	var pr []models.Producers
	j := es.Find(`span:contains("Producers:")`).Parent()
	j.Find("a").Each(func(_ int, s *goquery.Selection) {
		pnam, _ := s.Attr("title")
		plink, _ := s.Attr("href")
		var Typ string
		var MalId int
		h := strings.Split(plink, "/")
		for i, k := range h {
			if i == 1 {
				Typ = k
			}
			if i == 3 {
				MalId, _ = strconv.Atoi(strings.TrimSpace(k))
			}
		}
		tmp := models.Producers{
			MalID: MalId,
			Type:  Typ,
			Name:  pnam,
			URL:   os.Getenv("BASE_URL") + plink,
		}
		if MalId != 0 {
			pr = append(pr, tmp)
		}
	})
	return pr
}

func GetLicensor(es *goquery.Selection) []models.Licensors {
	var li []models.Licensors
	j := es.Find(`span:contains("Licensors:")`).Parent()
	j.Find("a").Each(func(_ int, s *goquery.Selection) {
		lnam, _ := s.Attr("title")
		llink, _ := s.Attr("href")
		var Typ string
		var MalId int
		h := strings.Split(llink, "/")
		for i, k := range h {
			if i == 1 {
				Typ = k
			}
			if i == 3 {
				MalId, _ = strconv.Atoi(strings.TrimSpace(k))
			}
		}
		tmp := models.Licensors{
			MalID: MalId,
			Type:  Typ,
			Name:  lnam,
			URL:   os.Getenv("BASE_URL") + llink,
		}
		if MalId != 0 {
			li = append(li, tmp)
		}
	})
	return li
}

func GetStudios(es *goquery.Selection) []models.Studios {
	var li []models.Studios
	j := es.Find(`span:contains("Studios:")`).Parent()
	j.Find("a").Each(func(_ int, s *goquery.Selection) {
		lnam, _ := s.Attr("title")
		llink, _ := s.Attr("href")
		var Typ string
		var MalId int
		h := strings.Split(llink, "/")
		for i, k := range h {
			if i == 1 {
				Typ = k
			}
			if i == 3 {
				MalId, _ = strconv.Atoi(strings.TrimSpace(k))
			}
		}
		tmp := models.Studios{
			MalID: MalId,
			Type:  Typ,
			Name:  lnam,
			URL:   os.Getenv("BASE_URL") + llink,
		}
		if MalId != 0 {
			li = append(li, tmp)
		}
	})
	return li
}

func GetGenre(es *goquery.Selection) []models.Genres {
	var li []models.Genres
	j := es.Find(`span:contains("Genres:")`).Parent()
	j.Find("a").Each(func(_ int, s *goquery.Selection) {
		lnam, _ := s.Attr("title")
		llink, _ := s.Attr("href")
		var Typ string
		var MalId int
		h := strings.Split(llink, "/")
		for i, k := range h {
			if i == 1 {
				Typ = k
			}
			if i == 3 {
				MalId, _ = strconv.Atoi(strings.TrimSpace(k))
			}
		}
		tmp := models.Genres{
			MalID: MalId,
			Type:  Typ,
			Name:  lnam,
			URL:   os.Getenv("BASE_URL") + llink,
		}
		if MalId != 0 {
			li = append(li, tmp)
		}
	})
	return li
}

func GetThemes(es *goquery.Selection) []models.Themes {
	var li []models.Themes
	j := es.Find(`span:contains("Themes:")`).Parent()
	j.Find("a").Each(func(_ int, s *goquery.Selection) {
		lnam, _ := s.Attr("title")
		llink, _ := s.Attr("href")
		var Typ string
		var MalId int
		h := strings.Split(llink, "/")
		for i, k := range h {
			if i == 1 {
				Typ = k
			}
			if i == 3 {
				MalId, _ = strconv.Atoi(strings.TrimSpace(k))
			}
		}
		tmp := models.Themes{
			MalID: MalId,
			Type:  Typ,
			Name:  lnam,
			URL:   os.Getenv("BASE_URL") + llink,
		}
		if MalId != 0 {
			li = append(li, tmp)
		}
	})
	return li
}

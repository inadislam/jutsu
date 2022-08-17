package controllers

import (
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	models "github.com/inadislam/jutsu/app/models/anime"
	"github.com/inadislam/jutsu/pkg/utils"
)

func GetCharacter(c *fiber.Ctx) error {
	id := c.Params("id")
	c_url := GetURL(id)
	var char []models.Characters
	var vac []models.VoiceActors
	h := colly.NewCollector()
	h.OnHTML("body", func(e *colly.HTMLElement) {
		es := e.DOM
		es.Find(".anime-character-container table").Each(func(_ int, s *goquery.Selection) {
			charac_url, _ := s.Find("td:nth-child(2) > div:nth-child(3) > a").Attr("href")
			malid := utils.GetMalId(charac_url)
			name := s.Find("h3.h3_character_name").Text()
			im, _ := s.Find("img:nth-child(1)").Attr("data-src")
			ima := strings.Split(im, "/images")
			var image string
			for i, k := range ima {
				if i == 1 {
					image = "https://cdn.myanimelist.net/images"+k
				}
			}
			imgwebp := strings.Join(strings.Split(image, ".jpg"), ".webp")
			space := regexp.MustCompile(`\s+`)
			role := strings.TrimSpace(space.ReplaceAllString(s.Find("td:nth-child(2) > div:nth-child(4)").Text(), " "))
			s.Find("td:nth-child(3) table tr").Each(func(_ int, ss *goquery.Selection) {
				vurl, _ := ss.Find("td:first-child a").Attr("href")
				vmalid := utils.GetMalId(vurl)
				im, _ := ss.Find("td:nth-child(2) img").Attr("data-src")
				ima := strings.Split(im, "/images")
				var image string
				for i, k := range ima {
					if i == 1 {
						image = "https://cdn.myanimelist.net/images"+k
					}
				}
				vname, _ := ss.Find("td:nth-child(2) img").Attr("alt")
				language := strings.TrimSpace(space.ReplaceAllString(ss.Find(".js-anime-character-language").Text(), " "))
				tmp := models.VoiceActors{
					Person: models.Person{
						MalID: vmalid,
						URL:   vurl,
						Images: models.VAImages{
							Jpg: models.CharJpg{
								ImageURL: image,
							},
						},
						Name: vname,
					},
					Language: language,
				}
				vac = append(vac, tmp)
			})
			tmp := models.Characters{
				Character: models.Character{
					MalID: malid,
					URL:   charac_url,
					Images: models.CharImages{
						Jpg: models.CharJpg{
							ImageURL: image,
						},
						Webp: models.CharWebp{
							ImageURL:      imgwebp,
							SmallImageURL: strings.Join(strings.Split(imgwebp, ".webp"), "t.webp"),
						},
					},
					Name: name,
				},
				Role:        role,
				VoiceActors: vac,
			}
			char = append(char, tmp)
		})
	})
	h.Visit(c_url)
	return c.Status(200).JSON(fiber.Map{
		"data": char,
	})
}

func GetURL(id string) string {
	char_url := os.Getenv("BASE_URL") + "anime/" + id + "/_/characters"
	return char_url
}

func GetStaff(c *fiber.Ctx) error {
	return c.JSON("Get Staff")
}

package internals

import (
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func Visitor(visitURL string, maxDepth int) []string {
	socialDomains := []string{"twitter.com", "instagram.com", "facebook.com", "twitch.tv", "tiktok.com"}
	var socialLinks []string
	var visitedLinks []string
	denyList := []string{".js", ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".mp4", ".webm", ".mp3", ".csv", ".ogg", ".wav", ".flac", ".aac", ".wma", ".wmv", ".avi", ".mpg", ".mpeg", ".mov", ".mkv", ".zip", ".rar", ".7z", ".tar", ".iso", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".rtf", ".odt", ".ods", ".odp", ".odg", ".odf", ".odb", ".odc", ".odm", ".avi", ".mpg", ".mpeg", ".mov", ".mkv", ".zip", ".rar", ".7z", ".tar", ".iso", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".rtf", ".odt", ".ods", ".odp", ".odg", ".odf", ".odb", ".odc", ".odm", ".mp4", ".webm", ".mp3", ".ogg", ".wav", ".flac", ".aac", ".wma", ".wmv", ".avi", ".mpg", ".mpeg", ".mov", ".mkv", ".zip", ".rar", ".7z", ".tar", ".iso", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".rtf", ".odt", ".ods", ".odp", ".odg", ".odf", ".odb", ".odc", ".odm", ".mp4", ".webm", ".mp3", ".ogg", ".wav", ".flac", ".aac", ".wma", ".wmv", ".avi", ".mpg", ".mpeg", ".mov", ".mkv", ".zip", ".rar", ".7z", ".tar", ".iso", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".rtf", ".odt"}

	c := colly.NewCollector()
	c.UserAgent = UserAgent
	c.SetRequestTimeout(5 * time.Second)
	c.MaxDepth = maxDepth
	c.AllowURLRevisit = false //there is a bug in colly that prevents this from working. We have to check it manually
	u, err := url.Parse(visitURL)
	if err != nil {
		panic(err)
	}
	domain := u.Host
	path := u.Path
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		u2, err := url.Parse(link)
		if err != nil {
			panic(err)
		}
		linkDomain := u2.Host
		for _, domain := range socialDomains {
			if strings.Contains(linkDomain, domain) {
				socialLinks = append(socialLinks, e.Request.URL.String()+"|"+link)
			}
		}
		if strings.Contains(linkDomain, domain) {
			visitFlag := true
			for _, extension := range denyList {
				if strings.Contains(strings.ToLower(link), extension) {
					visitFlag = false
				}
			}
			for _, value := range visitedLinks {
				if strings.ToLower(link) == value {
					visitFlag = false
				}
			}

			if !strings.HasPrefix(u2.Path, path) {
				visitFlag = false
			}
			// if it's True it will append
			if visitFlag {
				visitedLinks = append(visitedLinks, link)
				err := e.Request.Visit(link)
				if err != nil {
					panic(err)
				}

			}
		}

	})

	visError := c.Visit(visitURL)
	if visError != nil {
		panic(visError)
	}
	return socialLinks
}

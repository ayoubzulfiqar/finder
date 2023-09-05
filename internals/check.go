package internals

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/fatih/color"
)

func CheckTakeOver(socialLinks []string) {
	var alreadyChecked []string
	for _, value := range socialLinks {
		foundLink := strings.Split(value, "|")[0]
		socialLink := strings.Split(value, "|")[1]
		if StringInSlice(socialLink, &alreadyChecked) {
			continue
		}
		alreadyChecked = append(alreadyChecked, socialLink)
		if len(socialLink) > 60 || strings.Contains(socialLink, "intent/tweet") || strings.Contains(socialLink, "twitter.com/share") || strings.Contains(socialLink, "twitter.com/privacy") || strings.Contains(socialLink, "facebook.com/home") || strings.Contains(socialLink, "instagram.com/p/") {
			continue
		}
		u, err := url.Parse(socialLink)
		if err != nil {
			continue
		}
		domain := u.Host
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		if strings.Contains(domain, "facebook.com") {
			if strings.Count(socialLink, ".") > 1 {
				socialLink = "https://" + strings.Split(socialLink, ".")[1] + "." + strings.Split(socialLink, ".")[2]
			}
			socialLink = strings.Replace(socialLink, "www.", "", -1)
			tempLink := strings.Replace(socialLink, "facebook.com", "tr-tr.facebook.com", -1)
			resp, err := http.Get(tempLink)
			if err != nil {
				continue
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			if strings.Contains(string(body), "404 Not Found") {
				color.Green("Possible Takeover: " + socialLink + " at " + foundLink)

			}

		}
		if strings.Contains(domain, "tiktok.com") {
			if strings.Count(strings.Replace(socialLink, "www.", "", -1), ".") > 1 {
				continue
			}
			client := &http.Client{Transport: tr}

			req, err := http.NewRequest("GET", socialLink, nil)
			if err != nil {
				continue
			}

			req.Header.Set("User-Agent", UserAgent)

			resp, err := client.Do(req)
			if err != nil {
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode == 404 {
				color.Green("Possible Takeover: " + socialLink + " at " + foundLink)
			}
		}
		if strings.Contains(domain, "instagram.com") {

			if strings.Count(strings.Replace(socialLink, "www.", "", -1), ".") > 1 {
				continue
			}
			if !strings.Contains(socialLink, "instagram.com/") {
				continue
			}
			tempLink := "https://www.picuki.com/profile/" + strings.Split(socialLink, "instagram.com/")[1]
			client := &http.Client{Transport: tr}
			req, err := http.NewRequest("GET", tempLink, nil)
			if err != nil {
				continue
			}

			req.Header.Set("User-Agent", UserAgent)

			resp, err := client.Do(req)
			if err != nil {
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode == 404 {
				color.Green("Possible Takeover: " + socialLink + " at " + foundLink)
			}
		}
		if strings.Contains(domain, "twitter.com") {
			if strings.Count(strings.Replace(socialLink, "www.", "", -1), ".") > 1 {
				continue
			}
			u, err := url.Parse(socialLink)
			if err != nil {
				panic(err)
			}
			userName := u.Path
			tempLink := "https://nitter.net" + userName
			client := &http.Client{}
			req, err := http.NewRequest("GET", tempLink, nil)
			if err != nil {
				continue
			}

			req.Header.Set("User-Agent", UserAgent)

			resp, err := client.Do(req)
			if err != nil {
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode == 404 {
				color.Green("Possible Takeover: " + socialLink + " at " + foundLink)
			}
		}
	}
	// return
}

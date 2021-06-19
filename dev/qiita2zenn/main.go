package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	QiitaAccessToken = "47e42408b47cb4994c142efc4577fceaf405ee92"
	QiitaUser        = "kaakaa_hoe"
	QiitaPerPage     = 74
	QiitaMaxPage     = 1
)

type QiitaItem struct {
	CreatedAt string `json:"created_at"`
	Id        string `json:"id"`
	Title     string `json:"title"`
	Private   bool   `json:"private"`
}

type QiitaPatch struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

func main() {
	client := &http.Client{}

	var all []QiitaItem
	page := 1
	for {
		log.Printf("Get items. page %d", page)
		items, err := getItems(client, page)
		if err != nil {
			log.Fatal(err)
		}
		all = append(all, items...)
		if len(items) < 100 {
			break
		}
		log.Println(len(items))
		page += 1
		if QiitaMaxPage < page {
			break
		}
	}

	nums := len(all)
	for i, v := range all {
		if v.Private {
			// skip private article
			continue
		}
		log.Printf("Patching article(%d/%d): %s", i+1, nums, v.Title)
		if err := patchItem(client, v); err != nil {
			log.Printf(" [ERROR] %s", err.Error())
		}
	}
}

func getItems(client *http.Client, page int) ([]QiitaItem, error) {
	req, err := http.NewRequest(http.MethodGet, "https://qiita.com/api/v2/authenticated_user/items", nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	params.Add("page", strconv.Itoa(page))
	params.Add("per_page", fmt.Sprintf("%d", QiitaPerPage))
	params.Add("query", fmt.Sprintf("user:%s", QiitaUser))
	req.URL.RawQuery = params.Encode()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", QiitaAccessToken))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get items %s", resp.Status)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var items []QiitaItem
	if err := json.Unmarshal(b, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func patchItem(client *http.Client, item QiitaItem) error {
	t, err := time.Parse(time.RFC3339, item.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("https://zenn.dev/kaakaa/articles/qiita-%s-%s", t.Format("20060102"), item.Id)

	b, err := json.Marshal(QiitaPatch{
		Body:  fmt.Sprintf("この記事は Zenn に移行しました。\n%s", url),
		Title: item.Title,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("https://qiita.com/api/v2/items/%s", item.Id), bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", QiitaAccessToken))

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to request patch: %s", resp.Status)
	}
	return nil
}

package sonolusgo

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetEmptyList[ItemType SonolusItem](page int, queryMap map[string]string) (pageCount int, items []ItemType) {
	return 0, []ItemType{}
}

func GetEmptySearch() (search []ServerForm) {
	return []ServerForm{}
}

func GetEmptyItem[ItemType SonolusItem](name string) (item ItemType, description string, err error) {
	return item, description, nil
}

func GetEmptyRecommend[ItemType SonolusItem](name string) (items []ItemType) {
	return []ItemType{}
}

var fullListStr map[SonolusCategory][]byte

func loadFullListStr(filename string) []byte {
	if filename != "" {
		str, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		return str
	} else {
		return []byte{}
	}
}

func initFullListStr() {
	if InfoFilePath == nil {
		panic("not Initialized infoFilePath")
	}
	fullListStr = map[SonolusCategory][]byte{
		CategoryPostItem:    loadFullListStr(InfoFilePath.Posts),
		CategoryPlaylist:    loadFullListStr(InfoFilePath.Playlist),
		CategoryLevels:      loadFullListStr(InfoFilePath.Levels),
		CategorySkins:       loadFullListStr(InfoFilePath.Skins),
		CategoryBackgrounds: loadFullListStr(InfoFilePath.Backgrounds),
		CategoryEffect:      loadFullListStr(InfoFilePath.Effects),
		CategoryParticle:    loadFullListStr(InfoFilePath.Particles),
		CategoryEngine:      loadFullListStr(InfoFilePath.Engines),
		CategoryReplayItem:  loadFullListStr(InfoFilePath.Replays),
	}
}

func GetList[ItemType SonolusItem](page int, queryMap map[string]string) (pageCount int, items []ItemType) {
	var temp ItemType
	var fullItems []ItemType
	category := temp.GetCategory()
	if fullListStr == nil {
		initFullListStr()
	}
	err := json.Unmarshal(fullListStr[category], &fullItems)
	if err != nil {
		panic(err)
	}
	pageCount = (len(fullItems)-1)/20 + 1
	if page < pageCount {
		indexStart := page * 20
		indexEnd := (page + 1) * 20
		if indexEnd >= len(fullItems) {
			indexEnd = len(fullItems)
		}
		items = fullItems[indexStart:indexEnd]
	}
	return pageCount, items
}

func GetItem[ItemType SonolusItem](name string) (item ItemType, description string, err error) {
	var fullItems []ItemType
	category := item.GetCategory()
	if fullListStr == nil {
		initFullListStr()
	}
	err = json.Unmarshal(fullListStr[category], &fullItems)
	if err != nil {
		panic(err)
	}
	for _, listItem := range fullItems {
		if listItem.GetName() == name {
			return listItem, "", nil
		}
	}
	return item, "", fmt.Errorf("item not found")
}

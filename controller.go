package sonolusgo

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func parseSearchQueryStr(ctx *gin.Context, sto ServerOption) (key string, value string) {
	query, exist := ctx.GetQuery(sto.GetQuery())
	if exist {
		return sto.GetQuery(), sto.GetValueStr(query)
	} else {
		return sto.GetQuery(), ""
	}
}

func listParseQuery(ctx *gin.Context, search ServerOptionSection) (localization string, page int, queryMap map[string]string) {
	queryMap = make(map[string]string)
	localization = ctx.Query("localization")
	pageStr := ctx.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}
	for _, searchOption := range search.Options {
		key, val := parseSearchQueryStr(ctx, searchOption)
		queryMap[key] = val
	}
	// default add "keywords" into querymap
	queryMap["keywords"], _ = ctx.GetQuery("keywords")
	return localization, page, queryMap
}

func ListHandler[Item SonolusItem](handler SonolusService[Item]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		search := handler.Search()
		_, page, queryMap := listParseQuery(ctx, search)
		pageCount, items := handler.List(page, queryMap)
		ctx.JSON(http.StatusOK, ItemList[Item]{
			PageCount: pageCount,
			Items:     items,
			Searches:  search,
		})
	}
}

func DetailsHandler[Item SonolusItem](handler SonolusService[Item]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		itemName := ctx.Param("name")
		item, description, err := handler.Item(itemName)
		if err != nil {
			ctx.Status(http.StatusNotFound)
			log.Println(err)
			return
		}
		// recommend := handler.Recommend(itemName)
		ctx.JSON(http.StatusOK, ItemDetail[Item]{
			Item:         item,
			Description:  description,
			HasCommunity: false,
			Leaderboards: []interface{}{},
			Actions:      []interface{}{},
			Sections:     []ItemSection[Item]{
				//	{
				//	Title: "#RECOMMENDED",
				//	Items: recommend,
				//	Icon:  "star",
				//}
			},
		})
	}
}

func getDefaultSearch(ctx *gin.Context, search ServerOptionSection) (queryMap map[string]string) {
	// 这里的ctx仅借用，不会在这里传入search的query
	queryMap = make(map[string]string)
	for _, searchOption := range search.Options {
		key, val := parseSearchQueryStr(ctx, searchOption)
		queryMap[key] = val
	}
	return queryMap
}

func getFirst5Item[ItemType SonolusItem](items []ItemType) (result []ItemType) {
	if len(items) < 5 {
		return items
	} else {
		return items[0:5]
	}
}

func InfoHandler[Item SonolusItem](handler SonolusService[Item]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var temp Item
		_, items := handler.List(0, getDefaultSearch(ctx, handler.Search()))
		//if len(items) == 0 {
		//	ctx.Status(http.StatusNotFound)
		//	return
		//}
		ctx.JSON(http.StatusOK, ItemInfo[Item]{
			Banner: handler.Banner,
			Sections: []ItemSection[Item]{
				{
					Title:    temp.GetName(),
					Items:    getFirst5Item[Item](items),
					ItemType: "level", // TODO temporary fix for sonolus-test-server
				},
			},
			// Searches: []ServerOptionSection{handler.Search()},
			Searches: []ServerOptionSection{},
		})
	}
}

// TODO: Localization
func ServerInfoHandler(server *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ServerInfo{
			Title:             server.ServerName,
			Description:       "SonolusGo 服务器",
			HasAuthentication: false,
			HasMultiplayer:    false,
			Buttons: []ServerInfoButtonType{
				{"level"},
			},
			Banner:        server.ServerBanner,
			Configuration: Configuration{Options: []any{}},
		})
	}
}

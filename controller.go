package sonolusgo

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func parseSearchQueryStr(ctx *gin.Context, sto SearchOption) (key string, value string) {
	query, exist := ctx.GetQuery(sto.GetQuery())
	if exist {
		return sto.GetQuery(), sto.GetValueStr(query)
	} else {
		return sto.GetQuery(), ""
	}
}

func listParseQuery(ctx *gin.Context, search Search) (localization string, page int, queryMap map[string]string) {
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
	return localization, page, queryMap
}

func ListHandler[Item SonolusItem](handler SonolusHandlers[Item]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		search := handler.Search()
		_, page, queryMap := listParseQuery(ctx, search)
		pageCount, items := handler.List(page, queryMap)
		ctx.JSON(http.StatusOK, ItemList[Item]{
			PageCount: pageCount,
			Items:     items,
			Search:    search,
		})
	}
}

func DetailsHandler[Item SonolusItem](handler SonolusHandlers[Item]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		itemName := ctx.Param("name")
		item, description, err := handler.Item(itemName)
		if err != nil {
			ctx.Status(http.StatusNotFound)
			log.Println(err)
			return
		}
		recommend := handler.Recommend(itemName)
		ctx.JSON(http.StatusOK, ItemDetail[Item]{
			Item:        item,
			Description: description,
			Recommended: recommend,
		})
	}
}

func getDefaultSearch(ctx *gin.Context, search Search) (queryMap map[string]string) {
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

// TODO: Localization
func InfoHandler(server *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Levels Section
		levelsSearch := server.Handlers.Levels.Search()
		_, levelsItems := server.Handlers.Levels.List(0, getDefaultSearch(ctx, levelsSearch))
		// Skins Section
		skinsSearch := server.Handlers.Skins.Search()
		_, skinsItems := server.Handlers.Skins.List(0, getDefaultSearch(ctx, skinsSearch))
		// Background Section
		backgroundsSearch := server.Handlers.Backgrounds.Search()
		_, backgroundsItems := server.Handlers.Backgrounds.List(0, getDefaultSearch(ctx, backgroundsSearch))
		// Effect Section
		effectsSearch := server.Handlers.Effects.Search()
		_, effectsItems := server.Handlers.Effects.List(0, getDefaultSearch(ctx, effectsSearch))
		// Particle Section
		particlesSearch := server.Handlers.Particles.Search()
		_, particlesItems := server.Handlers.Particles.List(0, getDefaultSearch(ctx, particlesSearch))
		// Engine Section
		enginesSearch := server.Handlers.Engines.Search()
		_, enginesItems := server.Handlers.Engines.List(0, getDefaultSearch(ctx, enginesSearch))

		ctx.JSON(http.StatusOK, ServerInfo{
			Title:  server.ServerName,
			Banner: server.ServerBanner,
			Levels: Section[Level]{
				Items:  getFirst5Item[Level](levelsItems),
				Search: levelsSearch,
			},
			Skins: Section[Skin]{
				Items:  getFirst5Item[Skin](skinsItems),
				Search: skinsSearch,
			},
			Backgrounds: Section[Background]{
				Items:  getFirst5Item[Background](backgroundsItems),
				Search: backgroundsSearch,
			},
			Effects: Section[Effect]{
				Items:  getFirst5Item[Effect](effectsItems),
				Search: effectsSearch,
			},
			Particles: Section[Particle]{
				Items:  getFirst5Item[Particle](particlesItems),
				Search: particlesSearch,
			},
			Engines: Section[Engine]{
				Items:  getFirst5Item[Engine](enginesItems),
				Search: enginesSearch,
			},
		})
	}
}

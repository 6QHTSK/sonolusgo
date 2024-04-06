package sonolusgo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SonolusService[ItemType SonolusItem] struct {
	List      func(page int, queryMap map[string]string) (pageCount int, items []ItemType)
	Search    func() (search ServerOptionSection)
	Item      func(name string) (item ItemType, description string, err error)
	Recommend func(name string) (items []ItemType)
	Banner    SRL
}

type Handlers struct {
	Posts       SonolusService[Post]
	Playlist    SonolusService[Playlist]
	Levels      SonolusService[Level]
	Skins       SonolusService[Skin]
	Backgrounds SonolusService[Background]
	Effects     SonolusService[Effect]
	Particles   SonolusService[Particle]
	Engines     SonolusService[Engine]
	Replay      SonolusService[Replay]
}

type RouterGroups struct {
	Sonolus     *gin.RouterGroup
	Posts       *gin.RouterGroup
	Playlist    *gin.RouterGroup
	Levels      *gin.RouterGroup
	Skins       *gin.RouterGroup
	Backgrounds *gin.RouterGroup
	Effects     *gin.RouterGroup
	Particles   *gin.RouterGroup
	Engines     *gin.RouterGroup
	Replays     *gin.RouterGroup
}

type ServerInfoFilePath struct {
	Posts       string
	Playlist    string
	Levels      string
	Skins       string
	Backgrounds string
	Effects     string
	Particles   string
	Engines     string
	Replays     string
}

var InfoFilePath *ServerInfoFilePath

type Server struct {
	RepoDir      string
	ServerName   string
	ServerBanner SRL
	ServerInfo   ServerInfoFilePath
	Handlers     Handlers
	RouterGroups RouterGroups
}

func DefaultConfig() *Server {
	return &Server{
		RepoDir:      "./sonolus/repository",
		ServerName:   "Sonolus Go Framework Server",
		ServerBanner: SRL{},
		ServerInfo: ServerInfoFilePath{
			Levels:      "",
			Skins:       "./sonolus/skins.json",
			Backgrounds: "./sonolus/backgrounds.json",
			Effects:     "./sonolus/effects.json",
			Particles:   "./sonolus/particles.json",
			Engines:     "./sonolus/engines.json",
		},
		Handlers: Handlers{
			Posts: SonolusService[Post]{
				List:      GetEmptyList[Post],
				Search:    GetEmptySearch,
				Item:      GetItem[Post],
				Recommend: GetEmptyRecommend[Post],
				Banner:    SRL{},
			},
			Playlist: SonolusService[Playlist]{
				List:      GetEmptyList[Playlist],
				Search:    GetEmptySearch,
				Item:      GetItem[Playlist],
				Recommend: GetEmptyRecommend[Playlist],
				Banner:    SRL{},
			},
			Levels: SonolusService[Level]{
				List:      GetEmptyList[Level],
				Search:    GetEmptySearch,
				Item:      GetEmptyItem[Level],
				Recommend: GetEmptyRecommend[Level],
				Banner:    SRL{},
			},
			Skins: SonolusService[Skin]{
				List:      GetList[Skin],
				Search:    GetEmptySearch,
				Item:      GetItem[Skin],
				Recommend: GetEmptyRecommend[Skin],
				Banner:    SRL{},
			},
			Backgrounds: SonolusService[Background]{
				List:      GetList[Background],
				Search:    GetEmptySearch,
				Item:      GetItem[Background],
				Recommend: GetEmptyRecommend[Background],
				Banner:    SRL{},
			},
			Effects: SonolusService[Effect]{
				List:      GetList[Effect],
				Search:    GetEmptySearch,
				Item:      GetItem[Effect],
				Recommend: GetEmptyRecommend[Effect],
				Banner:    SRL{},
			},
			Particles: SonolusService[Particle]{
				List:      GetList[Particle],
				Search:    GetEmptySearch,
				Item:      GetItem[Particle],
				Recommend: GetEmptyRecommend[Particle],
				Banner:    SRL{},
			},
			Engines: SonolusService[Engine]{
				List:      GetList[Engine],
				Search:    GetEmptySearch,
				Item:      GetItem[Engine],
				Recommend: GetEmptyRecommend[Engine],
				Banner:    SRL{},
			},
			Replay: SonolusService[Replay]{
				List:      GetEmptyList[Replay],
				Search:    GetEmptySearch,
				Item:      GetItem[Replay],
				Recommend: GetEmptyRecommend[Replay],
				Banner:    SRL{},
			},
		},
	}
}

func loadItemHandlers[Item SonolusItem](itemGroup *gin.RouterGroup, itemService SonolusService[Item]) {
	itemGroup.GET("/info", InfoHandler[Item](itemService))
	itemGroup.GET("/list", ListHandler[Item](itemService))
	itemGroup.GET("/:name", DetailsHandler[Item](itemService))
}

func (server *Server) LoadHandlers(parentHandler *gin.Engine) {
	server.RouterGroups.Sonolus = parentHandler.Group("/sonolus", SonolusVersionHandler)
	{
		server.RouterGroups.Sonolus.GET("/info", ServerInfoHandler(server))
		server.RouterGroups.Posts = server.RouterGroups.Sonolus.Group("/posts")
		{
			loadItemHandlers[Post](server.RouterGroups.Posts, server.Handlers.Posts)
		}
		server.RouterGroups.Playlist = server.RouterGroups.Sonolus.Group("/playlists")
		{
			loadItemHandlers[Playlist](server.RouterGroups.Playlist, server.Handlers.Playlist)
		}
		server.RouterGroups.Levels = server.RouterGroups.Sonolus.Group("/levels")
		{
			loadItemHandlers[Level](server.RouterGroups.Levels, server.Handlers.Levels)
		}
		server.RouterGroups.Skins = server.RouterGroups.Sonolus.Group("/skins")
		{
			loadItemHandlers[Skin](server.RouterGroups.Skins, server.Handlers.Skins)
		}
		server.RouterGroups.Backgrounds = server.RouterGroups.Sonolus.Group("/backgrounds")
		{
			loadItemHandlers[Background](server.RouterGroups.Backgrounds, server.Handlers.Backgrounds)
		}
		server.RouterGroups.Effects = server.RouterGroups.Sonolus.Group("/effects")
		{
			loadItemHandlers[Effect](server.RouterGroups.Effects, server.Handlers.Effects)
		}
		server.RouterGroups.Particles = server.RouterGroups.Sonolus.Group("/particles")
		{
			loadItemHandlers[Particle](server.RouterGroups.Particles, server.Handlers.Particles)
		}
		server.RouterGroups.Engines = server.RouterGroups.Sonolus.Group("/engines")
		{
			loadItemHandlers[Engine](server.RouterGroups.Engines, server.Handlers.Engines)
		}
		server.RouterGroups.Replays = server.RouterGroups.Sonolus.Group("/replays")
		{
			loadItemHandlers[Replay](server.RouterGroups.Replays, server.Handlers.Replay)
		}
		server.RouterGroups.Sonolus.StaticFS("/repository", http.Dir(server.RepoDir))
	}
	InfoFilePath = &server.ServerInfo
}

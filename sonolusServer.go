package sonolusgo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SonolusHandlers[ItemType SonolusItem] struct {
	List      func(page int, queryMap map[string]string) (pageCount int, items []ItemType)
	Search    func() (search Search)
	Item      func(name string) (item ItemType, description string, err error)
	Recommend func(name string) (items []ItemType)
}

type Handlers struct {
	Levels      SonolusHandlers[Level]
	Skins       SonolusHandlers[Skin]
	Backgrounds SonolusHandlers[Background]
	Effects     SonolusHandlers[Effect]
	Particles   SonolusHandlers[Particle]
	Engines     SonolusHandlers[Engine]
}

type RouterGroups struct {
	Sonolus     *gin.RouterGroup
	Levels      *gin.RouterGroup
	Skins       *gin.RouterGroup
	Backgrounds *gin.RouterGroup
	Effects     *gin.RouterGroup
	Particles   *gin.RouterGroup
	Engines     *gin.RouterGroup
}

type ServerInfoFilePath struct {
	Levels      string
	Skins       string
	Backgrounds string
	Effects     string
	Particles   string
	Engines     string
}

var InfoFilePath *ServerInfoFilePath

type Server struct {
	RepoDir      string
	ServerName   string
	ServerBanner SRLServerBanner
	ServerInfo   ServerInfoFilePath
	Handlers     Handlers
	RouterGroups RouterGroups
}

func DefaultConfig() *Server {
	return &Server{
		RepoDir:      "./sonolus/repository",
		ServerName:   "Sonolus Go Framework Server",
		ServerBanner: SRLServerBanner{},
		ServerInfo: ServerInfoFilePath{
			Levels:      "",
			Skins:       "./sonolus/skins.json",
			Backgrounds: "./sonolus/backgrounds.json",
			Effects:     "./sonolus/effects.json",
			Particles:   "./sonolus/particles.json",
			Engines:     "./sonolus/engines.json",
		},
		Handlers: Handlers{
			Levels: SonolusHandlers[Level]{
				List:      GetEmptyList[Level],
				Search:    GetEmptySearch,
				Item:      GetEmptyItem[Level],
				Recommend: GetEmptyRecommend[Level],
			},
			Skins: SonolusHandlers[Skin]{
				List:      GetList[Skin],
				Search:    GetEmptySearch,
				Item:      GetItem[Skin],
				Recommend: GetEmptyRecommend[Skin],
			},
			Backgrounds: SonolusHandlers[Background]{
				List:      GetList[Background],
				Search:    GetEmptySearch,
				Item:      GetItem[Background],
				Recommend: GetEmptyRecommend[Background],
			},
			Effects: SonolusHandlers[Effect]{
				List:      GetList[Effect],
				Search:    GetEmptySearch,
				Item:      GetItem[Effect],
				Recommend: GetEmptyRecommend[Effect],
			},
			Particles: SonolusHandlers[Particle]{
				List:      GetList[Particle],
				Search:    GetEmptySearch,
				Item:      GetItem[Particle],
				Recommend: GetEmptyRecommend[Particle],
			},
			Engines: SonolusHandlers[Engine]{
				List:      GetList[Engine],
				Search:    GetEmptySearch,
				Item:      GetItem[Engine],
				Recommend: GetEmptyRecommend[Engine],
			},
		},
	}
}

func (server *Server) LoadHandlers(parentHandler *gin.Engine) {
	server.RouterGroups.Sonolus = parentHandler.Group("/sonolus", SonolusVersionHandler)
	{
		server.RouterGroups.Sonolus.GET("/info", InfoHandler(server))
		server.RouterGroups.Levels = server.RouterGroups.Sonolus.Group("/levels")
		{
			server.RouterGroups.Levels.GET("/list", ListHandler[Level](server.Handlers.Levels))
			server.RouterGroups.Levels.GET("/:name", DetailsHandler[Level](server.Handlers.Levels))
		}
		server.RouterGroups.Skins = server.RouterGroups.Sonolus.Group("/skins")
		{
			server.RouterGroups.Skins.GET("/list", ListHandler[Skin](server.Handlers.Skins))
			server.RouterGroups.Skins.GET("/:name", DetailsHandler[Skin](server.Handlers.Skins))
		}
		server.RouterGroups.Backgrounds = server.RouterGroups.Sonolus.Group("/backgrounds")
		{
			server.RouterGroups.Backgrounds.GET("/list", ListHandler[Background](server.Handlers.Backgrounds))
			server.RouterGroups.Backgrounds.GET("/:name", DetailsHandler[Background](server.Handlers.Backgrounds))
		}
		server.RouterGroups.Effects = server.RouterGroups.Sonolus.Group("/effects")
		{
			server.RouterGroups.Effects.GET("/list", ListHandler[Effect](server.Handlers.Effects))
			server.RouterGroups.Effects.GET("/:name", DetailsHandler[Effect](server.Handlers.Effects))
		}
		server.RouterGroups.Particles = server.RouterGroups.Sonolus.Group("/particles")
		{
			server.RouterGroups.Particles.GET("/list", ListHandler[Particle](server.Handlers.Particles))
			server.RouterGroups.Particles.GET("/:name", DetailsHandler[Particle](server.Handlers.Particles))
		}
		server.RouterGroups.Engines = server.RouterGroups.Sonolus.Group("/engines")
		{
			server.RouterGroups.Engines.GET("/list", ListHandler[Engine](server.Handlers.Engines))
			server.RouterGroups.Engines.GET("/:name", DetailsHandler[Engine](server.Handlers.Engines))
		}
		server.RouterGroups.Sonolus.StaticFS("/repository", http.Dir(server.RepoDir))
	}
	InfoFilePath = &server.ServerInfo
}

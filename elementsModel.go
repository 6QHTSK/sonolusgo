package sonolusgo

type SonolusCategory int

const (
	CategoryLevels SonolusCategory = iota
	CategorySkins
	CategoryBackgrounds
	CategoryEffect
	CategoryParticle
	CategoryEngine
	TotalCategoryCnt
)

type SonolusItem interface {
	Level | Skin | Background | Effect | Particle | Engine
	GetName() string
	GetCategory() SonolusCategory
}

type Skin struct {
	Name      string           `json:"name"`
	Version   int              `json:"version"`
	Title     string           `json:"title"`
	Subtitle  string           `json:"subtitle"`
	Author    string           `json:"author"`
	Thumbnail SRLSkinThumbnail `json:"thumbnail"`
	Data      SRLSkinData      `json:"data"`
	Texture   SRLSkinTexture   `json:"texture"`
}

func (item Skin) GetName() string {
	return item.Name
}

func (item Skin) GetCategory() SonolusCategory {
	return CategorySkins
}

type Background struct {
	Name          string                     `json:"name"`
	Version       int                        `json:"version"`
	Title         string                     `json:"title"`
	Subtitle      string                     `json:"subtitle"`
	Author        string                     `json:"author"`
	Thumbnail     SRLBackgroundThumbnail     `json:"thumbnail"`
	Data          SRLBackgroundData          `json:"data"`
	Image         SRLBackgroundImage         `json:"image"`
	Configuration SRLBackgroundConfiguration `json:"configuration"`
}

func (item Background) GetName() string {
	return item.Name
}

func (item Background) GetCategory() SonolusCategory {
	return CategoryBackgrounds
}

type Effect struct {
	Name      string             `json:"name"`
	Version   int                `json:"version"`
	Title     string             `json:"title"`
	Subtitle  string             `json:"subtitle"`
	Author    string             `json:"author"`
	Thumbnail SRLEffectThumbnail `json:"thumbnail"`
	Data      SRLEffectData      `json:"data"`
	Audio     SRLEffectAudio     `json:"audio"`
}

func (item Effect) GetName() string {
	return item.Name
}

func (item Effect) GetCategory() SonolusCategory {
	return CategoryEffect
}

type Particle struct {
	Name      string               `json:"name"`
	Version   int                  `json:"version"`
	Title     string               `json:"title"`
	Subtitle  string               `json:"subtitle"`
	Author    string               `json:"author"`
	Thumbnail SRLParticleThumbnail `json:"thumbnail"`
	Data      SRLParticleData      `json:"data"`
	Texture   SRLParticleTexture   `json:"texture"`
}

func (item Particle) GetName() string {
	return item.Name
}

func (item Particle) GetCategory() SonolusCategory {
	return CategoryParticle
}

type Engine struct {
	Name          string                 `json:"name"`
	Version       int                    `json:"version"`
	Title         string                 `json:"title"`
	Subtitle      string                 `json:"subtitle"`
	Author        string                 `json:"author"`
	Skin          Skin                   `json:"skin"`
	Background    Background             `json:"background"`
	Effect        Effect                 `json:"effect"`
	Particle      Particle               `json:"particle"`
	Thumbnail     SRLEngineThumbnail     `json:"thumbnail"`
	PlayData      SRLEnginePlayData      `json:"playData"`
	PreviewData   SRLEnginePreviewData   `json:"previewData"`
	TutorialData  SRLEngineTutorialData  `json:"tutorialData"`
	Rom           *SRLEngineRom          `json:"rom"`
	Configuration SRLEngineConfiguration `json:"configuration"`
}

func (item Engine) GetName() string {
	return item.Name
}

func (item Engine) GetCategory() SonolusCategory {
	return CategoryEngine
}

type UseItem[ItemType Skin | Background | Effect | Particle] struct {
	UseDefault bool      `json:"useDefault"`
	Item       *ItemType `json:"item"`
}

type Level struct {
	Name          string              `json:"name"`
	Version       int                 `json:"version"`
	Rating        int                 `json:"rating"`
	Title         string              `json:"title"`
	Artists       string              `json:"artists"`
	Author        string              `json:"author"`
	Engine        Engine              `json:"engine"`
	UseSkin       UseItem[Skin]       `json:"useSkin"`
	UseBackground UseItem[Background] `json:"useBackground"`
	UseEffect     UseItem[Effect]     `json:"useEffect"`
	UseParticle   UseItem[Particle]   `json:"useParticle"`
	Cover         SRLLevelCover       `json:"cover"`
	Bgm           SRLLevelBgm         `json:"bgm"`
	Preview       *SRLLevelPreview    `json:"preview"`
	Data          SRLLevelData        `json:"data"`
}

func (item Level) GetName() string {
	return item.Name
}

func (item Level) GetCategory() SonolusCategory {
	return CategoryLevels
}

type ItemList[ItemType SonolusItem] struct {
	PageCount int        `json:"pageCount"`
	Items     []ItemType `json:"items"`
	Search    Search     `json:"search"`
}

type ItemDetail[ItemType SonolusItem] struct {
	Item        ItemType   `json:"item"`
	Description string     `json:"description"`
	Recommended []ItemType `json:"recommended"`
}

type Section[ItemType SonolusItem] struct {
	Items  []ItemType `json:"items"`
	Search Search     `json:"search"`
}

type ServerInfo struct {
	Title       string              `json:"title"`
	Banner      SRLServerBanner     `json:"banner"`
	Levels      Section[Level]      `json:"levels"`
	Skins       Section[Skin]       `json:"skins"`
	Backgrounds Section[Background] `json:"backgrounds"`
	Effects     Section[Effect]     `json:"effects"`
	Particles   Section[Particle]   `json:"particles"`
	Engines     Section[Engine]     `json:"engines"`
}

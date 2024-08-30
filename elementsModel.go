package sonolusgo

type SonolusCategory int

const (
	CategoryPostItem SonolusCategory = iota
	CategoryPlaylist
	CategoryLevels
	CategorySkins
	CategoryBackgrounds
	CategoryEffect
	CategoryParticle
	CategoryEngine
	CategoryReplayItem
	TotalCategoryCnt
)

type SonolusItem interface {
	Post | Playlist | Level | Skin | Background | Effect | Particle | Engine | Replay
	GetName() string
	GetCategory() SonolusCategory
}

type SRL struct {
	Hash string `json:"hash"`
	Url  string `json:"url"`
}

type Tag struct {
	Title string `json:"title"`
	Icon  string `json:"icon,omitempty"`
}

type SonolusItemBase struct {
	Name     string `json:"name"`
	Source   string `json:"source,omitempty"`
	Version  int    `json:"version"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Author   string `json:"author"`
	Tags     []Tag  `json:"tags"`
}

type Skin struct {
	SonolusItemBase
	Thumbnail SRL `json:"thumbnail"`
	Data      SRL `json:"data"`
	Texture   SRL `json:"texture"`
}

func (item Skin) GetName() string {
	return item.Name
}

func (item Skin) GetCategory() SonolusCategory {
	return CategorySkins
}

type Background struct {
	SonolusItemBase
	Thumbnail     SRL `json:"thumbnail"`
	Data          SRL `json:"data"`
	Image         SRL `json:"image"`
	Configuration SRL `json:"configuration"`
}

func (item Background) GetName() string {
	return item.Name
}

func (item Background) GetCategory() SonolusCategory {
	return CategoryBackgrounds
}

type Effect struct {
	SonolusItemBase
	Thumbnail SRL `json:"thumbnail"`
	Data      SRL `json:"data"`
	Audio     SRL `json:"audio"`
}

func (item Effect) GetName() string {
	return item.Name
}

func (item Effect) GetCategory() SonolusCategory {
	return CategoryEffect
}

type Particle struct {
	SonolusItemBase
	Thumbnail SRL `json:"thumbnail"`
	Data      SRL `json:"data"`
	Texture   SRL `json:"texture"`
}

func (item Particle) GetName() string {
	return item.Name
}

func (item Particle) GetCategory() SonolusCategory {
	return CategoryParticle
}

type Engine struct {
	SonolusItemBase
	Skin          Skin       `json:"skin"`
	Background    Background `json:"background"`
	Effect        Effect     `json:"effect"`
	Particle      Particle   `json:"particle"`
	Thumbnail     SRL        `json:"thumbnail"`
	PlayData      SRL        `json:"playData"`
	WatchData     SRL        `json:"watchData"`
	PreviewData   SRL        `json:"previewData"`
	TutorialData  SRL        `json:"tutorialData"`
	Rom           *SRL       `json:"rom,omitempty"`
	Configuration SRL        `json:"configuration"`
}

func (item Engine) GetName() string {
	return item.Name
}

func (item Engine) GetCategory() SonolusCategory {
	return CategoryEngine
}

type UseItem[ItemType Skin | Background | Effect | Particle] struct {
	UseDefault bool      `json:"useDefault"`
	Item       *ItemType `json:"item,omitempty"`
}

type Level struct {
	Name          string              `json:"name"`
	Source        string              `json:"source"`
	Version       int                 `json:"version"`
	Rating        int                 `json:"rating"`
	Title         string              `json:"title"`
	Artists       string              `json:"artists"`
	Author        string              `json:"author"`
	Tags          []Tag               `json:"tags"`
	Engine        Engine              `json:"engine"`
	UseSkin       UseItem[Skin]       `json:"useSkin"`
	UseBackground UseItem[Background] `json:"useBackground"`
	UseEffect     UseItem[Effect]     `json:"useEffect"`
	UseParticle   UseItem[Particle]   `json:"useParticle"`
	Cover         SRL                 `json:"cover"`
	Bgm           SRL                 `json:"bgm"`
	Preview       *SRL                `json:"preview,omitempty"`
	Data          SRL                 `json:"data"`
}

func (item Level) GetName() string {
	return item.Name
}

func (item Level) GetCategory() SonolusCategory {
	return CategoryLevels
}

type Playlist struct {
	Name      string  `json:"name"`
	Source    string  `json:"source,omitempty"`
	Version   int     `json:"version"`
	Title     string  `json:"title"`
	Subtitle  string  `json:"subtitle"`
	Author    string  `json:"author"`
	Tags      []Tag   `json:"tags"`
	Levels    []Level `json:"levels"`
	Thumbnail SRL     `json:"thumbnail"`
}

func (item Playlist) GetName() string {
	return item.Name
}

func (item Playlist) GetCategory() SonolusCategory {
	return CategoryPlaylist
}

type Post struct {
	Name      string `json:"name"`
	Source    string `json:"source,omitempty"`
	Version   int    `json:"version"`
	Title     string `json:"title"`
	Time      string `json:"time"`
	Author    string `json:"author"`
	Tags      []Tag  `json:"tags"`
	Thumbnail SRL    `json:"thumbnail"`
}

func (item Post) GetName() string {
	return item.Name
}

func (item Post) GetCategory() SonolusCategory {
	return CategoryPostItem
}

type Replay struct {
	Name      string  `json:"name"`
	Source    string  `json:"source,omitempty"`
	Version   int     `json:"version"`
	Title     string  `json:"title"`
	Subtitle  string  `json:"subtitle"`
	Author    string  `json:"author"`
	Tags      []Tag   `json:"tags"`
	Levels    []Level `json:"levels"`
	Data      SRL     `json:"data"`
	Thumbnail SRL     `json:"thumbnail"`
}

func (item Replay) GetName() string {
	return item.Name
}

func (item Replay) GetCategory() SonolusCategory {
	return CategoryReplayItem
}

type ItemList[ItemType SonolusItem] struct {
	PageCount int          `json:"pageCount"`
	Items     []ItemType   `json:"items"`
	Searches  []ServerForm `json:"searches"`
}

type ItemSection[ItemType SonolusItem] struct {
	Title    string     `json:"title"`
	Items    []ItemType `json:"items"`
	ItemType string     `json:"itemType"`
	Icon     string     `json:"icon,omitempty"`
}

type ItemDetail[ItemType SonolusItem] struct {
	Item         ItemType                `json:"item"`
	Description  string                  `json:"description"`
	HasCommunity bool                    `json:"hasCommunity"`
	Leaderboards []interface{}           `json:"leaderboards"`
	Actions      []interface{}           `json:"actions"`
	Sections     []ItemSection[ItemType] `json:"sections"`
}

type ItemInfo[ItemType SonolusItem] struct {
	Banner   SRL                     `json:"banner"`
	Sections []ItemSection[ItemType] `json:"sections"`
	Searches []ServerForm            `json:"searches"`
}

type ServerInfoButtonType struct {
	Type string `json:"type"`
}

type Configuration struct {
	Options []any `json:"options"` // Not Certain here
}

type ServerInfo struct {
	Title             string                 `json:"title"`
	Description       string                 `json:"description,omitempty"`
	HasAuthentication bool                   `json:"hasAuthentication"`
	HasMultiplayer    bool                   `json:"hasMultiplayer"`
	Buttons           []ServerInfoButtonType `json:"buttons"`
	Banner            SRL                    `json:"banner"`
	Configuration     Configuration          `json:"configuration"`
}

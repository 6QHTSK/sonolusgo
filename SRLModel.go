package sonolusgo

type SRLItem struct {
	Type string `json:"type"`
	Hash string `json:"hash"`
	Url  string `json:"url"`
}

type SRLLevelCover SRLItem

func NewSRLLevelCover(hash string, url string) SRLLevelCover {
	return SRLLevelCover{
		Type: "LevelCover",
		Hash: hash,
		Url:  url,
	}
}

type SRLLevelBgm SRLItem

func NewSRLLevelBgm(hash string, url string) SRLLevelBgm {
	return SRLLevelBgm{
		Type: "LevelBgm",
		Hash: hash,
		Url:  url,
	}
}

type SRLLevelPreview SRLItem

func NewSRLLevelPreview(hash string, url string) SRLLevelPreview {
	return SRLLevelPreview{
		Type: "LevelPreview",
		Hash: hash,
		Url:  url,
	}
}

type SRLLevelData SRLItem

func NewSRLLevelData(hash string, url string) SRLLevelData {
	return SRLLevelData{
		Type: "LevelData",
		Hash: hash,
		Url:  url,
	}
}

type SRLSkinThumbnail SRLItem

func NewSRLSkinThumbnail(hash string, url string) SRLSkinThumbnail {
	return SRLSkinThumbnail{
		Type: "SkinThumbnail",
		Hash: hash,
		Url:  url,
	}
}

type SRLSkinData SRLItem

func NewSRLSkinData(hash string, url string) SRLSkinData {
	return SRLSkinData{
		Type: "SkinData",
		Hash: hash,
		Url:  url,
	}
}

type SRLSkinTexture SRLItem

func NewSRLSkinTexture(hash string, url string) SRLSkinTexture {
	return SRLSkinTexture{
		Type: "SkinTexture",
		Hash: hash,
		Url:  url,
	}
}

type SRLBackgroundThumbnail SRLItem

func NewSRLBackgroundThumbnail(hash string, url string) SRLBackgroundThumbnail {
	return SRLBackgroundThumbnail{
		Type: "BackgroundThumbnail",
		Hash: hash,
		Url:  url,
	}
}

type SRLBackgroundData SRLItem

func NewSRLBackgroundData(hash string, url string) SRLBackgroundData {
	return SRLBackgroundData{
		Type: "BackgroundData",
		Hash: hash,
		Url:  url,
	}
}

type SRLBackgroundImage SRLItem

func NewSRLBackgroundImage(hash string, url string) SRLBackgroundImage {
	return SRLBackgroundImage{
		Type: "BackgroundImage",
		Hash: hash,
		Url:  url,
	}
}

type SRLBackgroundConfiguration SRLItem

func NewSRLBackgroundConfiguration(hash string, url string) SRLBackgroundConfiguration {
	return SRLBackgroundConfiguration{
		Type: "BackgroundConfiguration",
		Hash: hash,
		Url:  url,
	}
}

type SRLEffectThumbnail SRLItem

func NewSRLEffectThumbnail(hash string, url string) SRLEffectThumbnail {
	return SRLEffectThumbnail{
		Type: "EffectThumbnail",
		Hash: hash,
		Url:  url,
	}
}

type SRLEffectData SRLItem

func NewSRLEffectData(hash string, url string) SRLEffectData {
	return SRLEffectData{
		Type: "EffectData",
		Hash: hash,
		Url:  url,
	}
}

type SRLEffectAudio SRLItem

func NewSRLEffectAudio(hash string, url string) SRLEffectAudio {
	return SRLEffectAudio{
		Type: "EffectAudio",
		Hash: hash,
		Url:  url,
	}
}

type SRLParticleThumbnail SRLItem

func NewSRLParticleThumbnail(hash string, url string) SRLParticleThumbnail {
	return SRLParticleThumbnail{
		Type: "ParticleThumbnail",
		Hash: hash,
		Url:  url,
	}
}

type SRLParticleData SRLItem

func NewSRLParticleData(hash string, url string) SRLParticleData {
	return SRLParticleData{
		Type: "ParticleData",
		Hash: hash,
		Url:  url,
	}
}

type SRLParticleTexture SRLItem

func NewSRLParticleTexture(hash string, url string) SRLParticleTexture {
	return SRLParticleTexture{
		Type: "ParticleTexture",
		Hash: hash,
		Url:  url,
	}
}

type SRLEngineThumbnail SRLItem

func NewSRLEngineThumbnail(hash string, url string) SRLEngineThumbnail {
	return SRLEngineThumbnail{
		Type: "EngineThumbnail",
		Hash: hash,
		Url:  url,
	}
}

type SRLEnginePlayData SRLItem

func NewSRLEnginePlayData(hash string, url string) SRLEnginePlayData {
	return SRLEnginePlayData{
		Type: "EnginePlayData",
		Hash: hash,
		Url:  url,
	}
}

type SRLEngineRom SRLItem

func NewSRLEngineRom(hash string, url string) SRLEngineRom {
	return SRLEngineRom{
		Type: "EngineRom",
		Hash: hash,
		Url:  url,
	}
}

type SRLEnginePreviewData SRLItem

func NewSRLEnginePreviewData(hash string, url string) SRLEnginePreviewData {
	return SRLEnginePreviewData{
		Type: "EnginePreviewData",
		Hash: hash,
		Url:  url,
	}
}

type SRLEngineTutorialData SRLItem

func NewSRLEngineTutorialData(hash string, url string) SRLEngineTutorialData {
	return SRLEngineTutorialData{
		Type: "EngineTutorialData",
		Hash: hash,
		Url:  url,
	}
}

type SRLEngineConfiguration SRLItem

func NewSRLEngineConfiguration(hash string, url string) SRLEngineConfiguration {
	return SRLEngineConfiguration{
		Type: "EngineConfiguration",
		Hash: hash,
		Url:  url,
	}
}

type SRLServerBanner SRLItem

func NewSRLServerBanner(hash string, url string) SRLServerBanner {
	return SRLServerBanner{
		Type: "ServerBanner",
		Hash: hash,
		Url:  url,
	}
}

type SRLEngineWatchData SRLItem

func NewSRLEngineWatchData(hash string, url string) SRLEngineWatchData {
	return SRLEngineWatchData{
		Type: "EngineWatchData",
		Hash: hash,
		Url:  url,
	}
}

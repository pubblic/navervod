package navervod

import (
	"encoding/xml"
	"fmt"
)

type SubjectTag struct {
	XMLName xml.Name `xml:"subject"`
	Link    string   `xml:"link,attr"`
	Text    string   `xml:",chardata"`
}

type MetaTag struct {
	XMLName    xml.Name `xml:"meta"`
	SubjectTag SubjectTag
}

type EncodingOptionTag struct {
	XMLName            xml.Name `xml:"encodingOption"`
	Id                 string   `xml:"id,attr"`
	Codec              string   `xml:"codec,attr"`
	Name               string   `xml:"name,attr"`
	Profile            string   `xml:"profile,attr"`
	IsEncodingComplete bool     `xml:"isEncodingComplete,attr"`
	Width              int      `xml:"width,attr"`
	Height             int      `xml:"height,attr"`
	Bitrate            int      `xml:"videoBitrate,attr"`
	CompletePercentage int      `xml:"completePercentage,attr"`
}

type BitrateTag struct {
	XMLName xml.Name `xml:"bitrate"`
	Video   float64  `xml:"video,attr"`
	Audio   float64  `xml:"audio,attr"`
}

type VideoTag struct {
	XMLName           xml.Name          `xml:"video"`
	EncodingOptionTag EncodingOptionTag ``
	BitrateTag        BitrateTag        ``
	Id                string            `xml:"id,attr"`
	UseP2P            bool              `xml:"useP2P,attr"`
	Duration          float64           `xml:"duration,attr"`
	Width             int               `xml:"width,attr"`
	Height            int               `xml:"height,attr"`
	Size              int64             `xml:"size,attr"`
	IsDefault         bool              `xml:"isDefault,attr"`
	Source            string            `xml:"source"`
}

type VideosTag struct {
	XMLName     xml.Name   `xml:"videos"`
	Type        string     `xml:"type,attr"`
	CanAutoPlay bool       `xml:"canAutoPlay,attr"`
	HasPreview  bool       `xml:"hasPreview,attr"`
	IsPreview   bool       `xml:"isPreview,attr"`
	VideoTags   []VideoTag `xml:"video"`
}

type KeyTag struct {
	XMLName xml.Name `xml:"key"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}

type StreamTag struct {
	XMLName   xml.Name `xml:"stream"`
	KeyTags   []KeyTag `xml:"keys>key"`
	Source    string   `xml:"source"`
	VideosTag VideosTag
}

type StreamsTag struct {
	XMLName    xml.Name    `xml:"streams"`
	StreamTags []StreamTag `xml:"stream"`
}

type CaptionTag struct {
	XMLName xml.Name `xml:"caption"`
	Lang    string   `xml:"language,attr"`
	Country string   `xml:"country,attr"`
	Locale  string   `xml:"locale,attr"`
	Label   string   `xml:"label,attr"`
	Type    string   `xml:"type,attr"`
	Source  string   `xml:",chardata"`
}

type CaptionsTag struct {
	XMLName     xml.Name     `xml:"captions"`
	CaptionLang string       `xml:"captionLang,attr"`
	CaptionTags []CaptionTag `xml:"caption"`
}

type VideoSourceTag struct {
	XMLName     xml.Name
	MetaTag     MetaTag
	VideosTag   VideosTag
	StreamsTag  StreamsTag
	CaptionsTag CaptionsTag
}

type ErrorTag struct {
	XMLName xml.Name `xml:"error"`
	Code    string   `xml:"code,attr"`
	Message string   `xml:"message"`
}

func (err ErrorTag) Error() string {
	return fmt.Sprintf("(%s) %s", err.Code, err.Message)
}

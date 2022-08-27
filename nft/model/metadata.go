package model

import "fmt"

type MetadataAttributeTraitType string

func (matt MetadataAttributeTraitType) String() string {
	return string(matt)
}

const (
	NumberTrait          MetadataAttributeTraitType = "number" // is default
	StringTrait          MetadataAttributeTraitType = "string"
	BoostPercentageTrait MetadataAttributeTraitType = "boost_percentage"
)

// CollectionMetadata represents ERC1155 token standard collection metadata
type CollectionMetadata struct {
	Name                 string                    `json:"name"`
	Symbol               string                    `json:"symbol"`
	SellerFeeBasisPoints int64                     `json:"seller_fee_basis_points"`
	Description          string                    `json:"description"`
	ExternalLink         string                    `json:"external_link"`
	Url                  string                    `json:"url"`
	Image                string                    `json:"image"`
	Background           string                    `json:"background"`
	Creators             []MetadataCreatorProperty `json:"creators"`
	IsMutable            bool                      `json:"is_mutable"`
	PrimarySaleHappened  bool                      `json:"primary_sale_happened"`
}

// TokenMetadata represents ERC1155/Metaplex token metadata standard with a bit of enhancements
type TokenMetadata struct {
	Symbol       string                     `json:"symbol"`
	Name         string                     `json:"name"`
	Description  string                     `json:"description"`
	Image        string                     `json:"image"`
	ImageData    string                     `json:"image_data"`
	Background   string                     `json:"background_color"`
	ExternalUrl  string                     `json:"external_url"`
	AnimationUrl string                     `json:"animation_url"`
	AudioUrl     string                     `json:"audio_url"`
	VideoUrl     string                     `json:"video_url"`
	YoutubeUrl   string                     `json:"youtube_url"`
	Attributes   []TokenMetadataAttribute   `json:"attributes"`
	Properties   []any                      `json:"properties"`
	Collection   TokenMetadataCollectionRef `json:"collection"`
}

func (t TokenMetadata) Validate() error {
	if len(t.Name) == 0 {
		return fmt.Errorf("metadata name can't be empty")
	}

	if len(t.Symbol) == 0 {
		return fmt.Errorf("metadata symbol can't be empty")
	}

	for i := range t.Attributes {
		if t.Attributes[i].TraitType == "" {
			return fmt.Errorf("empty trait type (index: %d)", i)
		}
	}

	for i, property := range t.Properties {
		switch property.(type) {
		case MetadataDefaultProperty:
			prop := property.(MetadataDefaultProperty)
			if prop.Name == "" || prop.Value == "" {
				return fmt.Errorf("metadata property can't be empty: %d", i)
			}
		case MetadataCreatorProperty:
			prop := property.(MetadataCreatorProperty)
			if len(prop.Address) < 42 {
				return fmt.Errorf("invalid creator address length: %d", i)
			}
		case MetadataFileProperty:
			prop := property.(MetadataFileProperty)
			if len(prop.Uri) == 0 {
				return fmt.Errorf("invalid metadata uri: %d", i)
			}
			if len(prop.Type) == 0 {
				return fmt.Errorf("invalid file type: %d", i)
			}
		}
	}

	return nil
}

// TokenMetadataAttribute represents NFT Metadata rich attribute
type TokenMetadataAttribute struct {
	Value        string                     `json:"value"`
	MaxValue     string                     `json:"max_value"`
	DisplayValue string                     `json:"display_value"`
	TraitType    MetadataAttributeTraitType `json:"trait_type"`
	DisplayType  string                     `json:"display_type"`
}

// TokenMetadataCollectionRef represents collection metadata reference
type TokenMetadataCollectionRef struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}

// TokenMetadataProperty defines generic type for metadata properties
type TokenMetadataProperty interface {
	MetadataDefaultProperty | MetadataCreatorProperty | MetadataFileProperty
}

type TokenMetadataProperties[T TokenMetadataProperty] []any

type MetadataDefaultProperty struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type MetadataCreatorProperty struct {
	Address string `json:"address"`
	Share   string `json:"share"`
}

type MetadataFileProperty struct {
	Uri  string `json:"uri"`
	Type string `json:"type"`
}

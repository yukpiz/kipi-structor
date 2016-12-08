package model

import (
	"regexp"
	"strconv"
)

type Item struct {
	Id                  int    `xml:"ID,attr"`
	DbStoreType         int    `xml:"DB_StoreType,attr"`
	Category            string `xml:"Category,attr"`
	Locale              string `xml:"Locale,attr"`
	Feature             string `xml:"feature,attr"`
	Xml                 string `xml:"XML,attr"`
	EnName              string `xml:"Text_Name0,attr"`
	JpNameTag           string `xml:"Text_Name1,attr"`
	JpDescTag           string `xml:"Text_Desc1,attr"`
	BundleType          int    `xml:"Bundle_Type,attr"`
	BundleMax           int    `xml:"Bundle_Max,attr"`
	PriceBuy            int    `xml:"Price_Buy,attr"`
	PriceSell           int    `xml:"Price_Sell,attr"`
	AttrActionFlag      int    `xml:"Attr_ActionFlag,attr"`
	AttrType            int    `xml:"Attr_Type,attr"`
	AttrGrade           int    `xml:"Attr_Grade,attr"`
	AttrRaceFilter      string `xml:"Attr_RaceFilter,attr"`
	FileMaleMesh        string `xml:"File_MaleMesh,attr"`
	FileFemaleMesh      string `xml:"File_FemaleMesh,attr"`
	FileGiantMesh       string `xml:"File_GiantMesh,attr"`
	FileFemaleGiantMesh string `xml:"File_FemaleGiantMesh,attr"`
	FileFieldMesh       string `xml:"File_FieldMesh,attr"`
	FileInvImage        string `xml:"File_InvImage,attr"`
	InvXSize            int    `xml:"Inv_XSize,attr"`
	InvYSize            int    `xml:"Inv_YSize,attr"`
	AppWeaponActionType int    `xml:"App_WeaponActionType,attr"`
	AppWearType         string `xml:"App_WearType,attr"`
	AppUseC4Layer       bool   `xml:"App_UseC4Layer,attr"`
	AppColor1           int    `xml:"App_Color1,attr"`
	AppColor2           int    `xml:"App_Color2,attr"`
	AppColor3           int    `xml:"App_Color3,attr"`
	AppColorOrder       int    `xml:"App_ColorOrder,attr"`
	AppAnimationType    string `xml:"App_AnimationType,attr"`
	AppSittingType      int    `xml:"App_SittingType,attr"`
	TasteBeauty         int    `xml:"Taste_Beauty,attr"`
	TasteIndivisuality  int    `xml:"Taste_Indivisuality,attr"`
	TasteLuxury         int    `xml:"Taste_Luxury,attr"`
	TasteToughness      int    `xml:"Taste_Toughness,attr"`
	TasteUtility        int    `xml:"Taste_Utility,attr"`
	TasteRarity         int    `xml:"Taste_Rarity,attr"`
	TasteMeaning        int    `xml:"Taste_Meaning,attr"`
	TasteAdult          int    `xml:"Taste_Adult,attr"`
	TasteManiac         int    `xml:"Taste_Maniac,attr"`
	TasteAnime          int    `xml:"Taste_Anime,attr"`
	TasteSexy           int    `xml:"Taste_Sexy,attr"`
	MetalwareUIToolTip  bool   `xml:"Metalware_UItooltip,attr"`
	EnchantUIToolTip    bool   `xml:"Enchant_UItooltip,attr"`
	UpgradeUIToolTip    bool   `xml:"Upgrade_UItooltip,attr"`
	ParBlockUseFlag     bool   `xml:"Par_BlockUseFlag,attr"`
	ParUpgradeMax       int    `xml:"Par_UpgradeMax,attr"`
	ParDurabilityMax    int    `xml:"Par_DurabillityMax,attr"`
	ParDefense          int    `xml:"Par_Defense,attr"`
	ParProtectRate      int    `xml:"Par_ProtectRate,attr"`
	ParAttackMin        int    `xml:"Par_AttackMin,attr"`
	ParAttackMax        int    `xml:"Par_AttackMax,attr"`
	ParWAttackMin       int    `xml:"Par_WAttackMin,attr"`
	ParWAttackMax       int    `xml:"Par_WAttackMax,attr"`
	ParCriticalRate     int    `xml:"Par_CriticalRate,attr"`
	ParAttackBalance    int    `xml:"Par_AttackBalance,attr"`
	ParEffectiveRange   int    `xml:"Par_EffectiveRange,attr"`
	ParAttackSpeed      int    `xml:"Par_AttackSpeed,attr"`
	ParDownHitCount     int    `xml:"Par_DownHitCount,attr"`
	SmartSearchFlag     string `xml:"SmartSearchFlag,attr"`
	JpNameId            int
	JpName              string
	JpDescId            int
	JpDesc              string
}

func (i *Item) Convert() {
	reg := regexp.MustCompile(`xml\.itemdb\.(\d*)`)
	m0 := reg.FindAllStringSubmatch(i.JpNameTag, -1)
	if len(m0) != 0 {
		i.JpNameId, _ = strconv.Atoi(m0[0][1])
	}

	m1 := reg.FindAllStringSubmatch(i.JpDescTag, -1)
	if len(m1) != 0 {
		i.JpDescId, _ = strconv.Atoi(m1[0][1])
	}
}

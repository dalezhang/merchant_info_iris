package model

import "gopkg.in/mgo.v2/bson"

type Merchant struct {
	ID                   bson.ObjectId `bson:"_id" json:"id"`
	MerchantID           string        `bson:"merchant_id" json:"merchant_id"`
	Name                 string        `bson:"name" json:"name"`
	Status               int64         `bson:"status" json:"status"`
	PartnerID            string        `bson:"partner_id" json:"partner_id"`
	PartnerMchID         string        `bson:"partner_mch_id" json:"partner_mch_id"`
	Password             string        `bson:"password" json:"password"`
	Memo                 string        `bson:"memo" json:"memo"`
	Address              string        `bson:"address" json:"address"`
	AlipayChannelTypeLv1 string        `bson:"alipay_channel_type_lv1 json:"alipay_channel_type_lv1"`
	AlipayChannelTypeLv2 string        `bson:"alipay_channel_type_lv2 json:"alipay_channel_type_lv2"`
	Appid                string        `bson:"appid" json:"appid"`
	PrivateKey           string        `bson:"private_key" json:"private_key"`
	ProductDesc          string        `bson:"product_desc" json:"product_desc"`
	ProductWebsite       string        `bson:"product_website" json:"product_website"`
	Province             string        `bson:"province" json:"province"`
	RouteID              string        `bson:"route_id" json:"route_id"`
	ServiceTel           string        `bson:"service_tel" json:"service_tel"`
	ShareKey             string        `bson:"share_key" json:"share_key"`
	StatusDesc           string        `bson:"status_desc" json:"status_desc"`
	T1Rate               string        `bson:"t1_rate" json:"t1_rate"`
	Urbn                 string        `bson:"urbn" json:"urbn"`
	WechatChannelTypeLv2 string        `bson:"wechat_channel_type_ json:"wechat_channel_type_lv2"`
	Zone                 string        `bson:"zone" json:"zone"`
	ZxAccount            string        `bson:"zx_account" json:"zx_account"`
	ContactEmail         string        `bson:"contact_email" json:"contact_email"`
	ContactName          string        `bson:"contact_name" json:"contact_name"`
	ContactTel           string        `bson:"contact_tel" json:"contact_tel"`
	D0Rate               string        `bson:"d0_rate" json:"d0_rate"`
	ErrorDesc            string        `bson:"error_desc" json:"error_desc"`
	FixedFee             int64         `bson:"fixed_fee" json:"fixed_fee"`
	FullName             string        `bson:"full_name" json:"full_name"`
	H5Path               string        `bson:"h5_path" json:"h5_path"`
	Industry             string        `bson:"industry" json:"industry"`
	InfcStatus           string        `bson:"infc_status" json:"infc_status"`
	MchDealType          string        `bson:"mch_deal_type" json:"mch_deal_type"`
	MchType              string        `bson:"mch_type" json:"mch_type"`
	NanYueMccCode        string        `bson:"nan_yue_mcc_code" json:"nan_yue_mcc_code"`
	BankInfo             struct {
		AccountNum       string `bson:"account_num"  json:"account_num"`
		AccountType      string `bson:"account_type"  json:"account_type"`
		BankFullName     string `bson:"bank_full_name"  json:"bank_full_name"`
		BankSubCode      string `bson:"bank_sub_code"  json:"bank_sub_code"`
		BankSubCodeBank  string `bson:"bank_sub_code_bank"  json:"bank_sub_code_bank"`
		OwnerIDcard      string `bson:"owner_idcard"  json:"owner_idcard"`
		OwnerName        string `bson:"owner_name"  json:"owner_name"`
		Province         string `bson:"province"  json:"province"`
		RightBankCardKey string `bson:"right_bank_card_key"  json:"right_bank_card_key"`
		Urbn             string `bson:"urbn"  json:"urbn"`
		Zone             string `bson:"zone"  json:"zone"`
	} `bson:"bank_info" json:"bank_info"`
	Company struct {
		AccountLicenceKey string `bson:"account_licence_key"  json:"account_licence_key"`
		LicenseCode       string `bson:"license_code"  json:"license_code"`
		LicenseKey        string `bson:"license_key"  json:"license_key"`
		OrgPhotoKey       string `bson:"org_photo_key"  json:"org_photo_key"`
		ShopPictureKey    string `bson:"shop_picture_key"  json:"shop_picture_key"`
	} `bson:"company" json:"company"`
	Instalment struct {
		InstalTypes []interface{} `bson:"instal_types"  json:"instal_types"`
		MchID       string        `bson:"mch_id"  json:"mch_id"`
		MchKey      string        `bson:"mch_key"  json:"mch_key"`
		Rate12      string        `bson:"rate_12"  json:"rate_12"`
		Rate24      string        `bson:"rate_24"  json:"rate_24"`
		Rate6       string        `bson:"rate_6"  json:"rate_6"`
		RatePayer   int64         `bson:"rate_payer"  json:"rate_payer"`
		Status      int64         `bson:"status"  json:"status"`
	} `bson:"instalment" json:"instalment"`
	LegalPerson struct {
		Email                string `bson:"email"  json:"email"`
		IDWithHandKey        string `bson:"id_with_hand_key"  json:"id_with_hand_key"`
		IDentityCardBackKey  string `bson:"identity_card_back_key"  json:"identity_card_back_key"`
		IDentityCardFrontKey string `bson:"identity_card_front_key"  json:"identity_card_front_key"`
		IDentityCardNum      string `bson:"identity_card_num"  json:"identity_card_num"`
		Name                 string `bson:"name"  json:"name"`
		Tel                  string `bson:"tel"  json:"tel"`
	} `bson:"legal_person" json:"legal_person"`
}

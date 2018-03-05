package model

import "gopkg.in/mgo.v2/bson"

type SubMerchant struct {
	ID            bson.ObjectId `bson:"_id"  json:"id"`
	MerchantID    bson.ObjectId `bson:"merchant_id"`
	SubMerchantID string        `bson:"sub_merchant_id"  json:"sub_merchant_id"`
	PartnerMchID  string        `bson:"partner_mch_id"  json:"partner_mch_id"`
	AccountNum    string        `bson:"account_num"  json:"account_num"`
	AccountType   string        `bson:"account_type"  json:"account_type"`
	BankFullName  string        `bson:"bank_full_name"  json:"bank_full_name"`
	BankSubCode   string        `bson:"bank_sub_code"  json:"bank_sub_code"`
	ContactTel    string        `bson:"contact_tel"  json:"contact_tel"`
	OwnerIDcard   string        `bson:"owner_idcard"  json:"owner_idcard"`
	OwnerName     string        `bson:"owner_name"  json:"owner_name"`
	Province      string        `bson:"province"  json:"province"`
	Urbn          string        `bson:"urbn"  json:"urbn"`
	Zone          string        `bson:"zone"  json:"zone"`
}

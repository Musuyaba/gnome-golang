package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Piece struct {
	Data_Print_Id int32  `bson:"Data_Print_Id"`
	Serialisasi   string `bson:"Serialisasi"`
}
type Inner struct {
	Data_Print_Id int32   `bson:"Data_Print_Id"`
	Serialisasi   string  `bson:"Serialisasi"`
	Pieces        []Piece `bson:"Data_Print_Id"`
}
type Box struct {
	Data_Print_Id int32   `bson:"Data_Print_Id"`
	Serialisasi   string  `bson:"Serialisasi"`
	Wrap          []Inner `bson:"Wrap"`
}

type Status string

const (
	produksi  Status = "produksi"
	karantina Status = "karantina"
	reject    Status = "reject"
	lulus     Status = "lulus"
)

type Pallete struct {
	Data_Print_Id *int32        `bson:"Data_Print_Id"`
	Parent_ID     *int32        `bson:"Parent_ID"`
	Serial_Level  *int32        `bson:"Serial_Level"`
	Barcode       *string       `bson:"Barcode"`
	Serialisasi   *string       `bson:"Serialisasi"`
	Batch_No      *string       `bson:"Batch_No"`
	Process_Order *string       `bson:"Process_Order"`
	Scanned       time.Time     `bson:"Scanned"`
	Product       *string       `bson:"product"`
	Nie           *string       `bson:"nie"`
	Sku           *string       `bson:"sku"`
	Counter       *int32        `bson:"Counter"`
	Berat         *float32      `bson:"berat"`
	Md            time.Time     `bson:"md"`
	Ed            time.Time     `bson:"ed"`
	Username      *string       `bson:"Username"`
	Station_name  *string       `bson:"station_name"`
	Grup          *string       `bson:"grup"`
	Ipc           *string       `bson:"ipc"`
	Sample        time.Time     `bson:"Sample"`
	Packer        *string       `bson:"packer"`
	Upload_line   time.Time     `bson:"upload_line"`
	Status        Status        `bson:"status"`
	Sjp           *string       `bson:"sjp"`
	Done          bson.RawValue `bson:"done,omitempty"`
	Synced        bson.RawValue `bson:"synced,omitempty"`
	Koli          []Box         `bson:"Koli"`
}

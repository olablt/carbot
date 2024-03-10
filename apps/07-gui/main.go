package main

import (
	"fmt"
	"log"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olablt/carbot"
	"github.com/olablt/carbot/ui"
)

type Command struct {
	Name string
	Func func()
	Key  key.Filter
	Icon *widget.Icon
}

var confDir = "./apps/07-gui/.conf/"

func main() {
	setupPath(confDir)
	adsDB, err := carbot.OpenCarbotDB(confDir)
	if err != nil {
		log.Fatal("[DEBUG] error opening database", err)
	}
	ads, err := adsDB.GetAds()
	if err != nil {
		log.Fatal("[DEBUG] error getting ads", err)
	}

	// // icons.ContentSave
	ag := NewAdsGUI()
	// commands := []Command{
	// 	{Name: "File: New", Func: nil, Key: key.Filter{Name: "N", Required: key.ModCtrl}},
	// 	{Name: "File: Open", Func: nil, Key: key.Filter{}},
	// 	{Name: "File: Save", Func: nil, Key: key.Filter{Name: "S", Required: key.ModCtrl}, Icon: icons.ContentSave},
	// 	{Name: "File: Save As", Func: nil, Key: key.Filter{}},
	// 	{Name: "Edit: Undo", Func: nil, Key: key.Filter{}},
	// 	{Name: "Edit: Redo", Func: nil, Key: key.Filter{}},
	// 	{Name: "Edit: Cut", Func: nil, Key: key.Filter{}},
	// 	{Name: "Format: Indent", Func: nil, Key: key.Filter{}},
	// 	{Name: "Format: Outdent", Func: nil, Key: key.Filter{}},
	// }
	for i, ad := range ads {
		log.Printf("Ad: %v %v %v %v", ad.Title, ad.Price, ad.Subtitle, ad.Other)
		ag.RegisterCommand(fmt.Sprintf("#%v %v %v \n%v\n%v", i, ad.Subtitle, ad.Price, ad.Other, ad.Title), nil, key.Filter{})
	}

	ui.Loop(func(win *app.Window, gtx layout.Context, th *material.Theme) {
		gtx.Metric = unit.Metric{
			PxPerDp: 1.5,
			PxPerSp: 1.5,
			// PxPerDp: 1.8,
			// PxPerSp: 1.8,
			// PxPerDp: 4,
			// PxPerSp: 4,
		}
		// layout
		ag.Layout(gtx, th)
	})
}

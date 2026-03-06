// Copyright 2024-2026 ProntoGUI, LLC
// ProntoGUI™ is a trademark of ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"

	pg "github.com/prontogui/golib"
)

func imageDemo(pgui pg.ProntoGUI) {

	img, _ := pg.ImageWith{ID: "gopher", FromFile: "go-logo.png"}.Make()

	// Create a hidden group to hold images
	assetsGroup := pg.GroupWith{GroupItems: []pg.Primitive{img}, Status: 2}.Make()
	imgRef1, _ := pg.ImageWith{Ref: "gopher", Embodiment: "width:200, height:200"}.Make()
	imgRef2, _ := pg.ImageWith{Ref: "gopher", Embodiment: "width:300, height:300"}.Make()

	mainFrame := pg.FrameWith{FrameItems: []pg.Primitive{imgRef1, imgRef2}}.Make()

	pgui.SetGUI(mainFrame, assetsGroup)

	for {
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}
}

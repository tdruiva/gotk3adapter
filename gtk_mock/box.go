package gtk_mock

import "github.com/twstrike/gotk3adapter/gtki"

type MockBox struct {
	MockContainer
	gtki.Orientable
}

func (*MockBox) PackEnd(v1 gtki.Widget, v2, v3 bool, v4 uint) {
}

func (*MockBox) PackStart(v1 gtki.Widget, v2, v3 bool, v4 uint) {
}

func (*MockBox) SetChildPacking(v1 gtki.Widget, v2, v3 bool, v4 uint, v5 gtki.PackType) {
}

func (*MockBox) GetOrientation() Orientation {
	return gtki.HorizontalOrientation
}

func (*MockBox) SetOrientation(o Orientation) {
}

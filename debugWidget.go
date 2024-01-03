package debugger

import (
	"strconv"
	"errors"

	"go.wit.com/log"
	"go.wit.com/gui/gui"
)


/*
func setActiveWidget(w *gui.Node) {
	if (w == nil) {
		log.Log(debugError, "setActiveWidget() was sent nil !!!")
		return
	}
	activeWidget = w
	log.Log(true, "The Widget is set to", w.id, w.Name)
	if (activeLabel == nil) {
		// the debug window doesn't exist yet so you can't display the change
		// TODO: make a fake binary tree for this(?)
		return
	}
	title := "ID =" + strconv.Itoa(w.id) + " " + w.Name
	activeLabel.SetText(title)
	activeLabelType.SetText("widget.Type = " + w.WidgetType.String())
	return
}
*/

func DebugWidgetWindow(w *gui.Node) {
	var newW, newB *gui.Node
	if (bugWidget != nil) {
		// this window was already created. Just change the widget we are working against
		setActiveWidget(w)
		return
	}

	newW = w.NewWindow("Widgets")
	newW.Custom = newW.StandardClose
	bugWidget = newW
	newB = newW.NewBox("hBox", true)

	g := newB.NewGroup("widget:")

	g2 := g.NewGroup("widget:")
	activeLabel = g2.NewLabel("undef")
	g2 = g.NewGroup("type:")
	activeLabelType = g2.NewLabel("undef")
	g2 = g.NewGroup("New name:")
	activeLabelNewName = g2.NewCombobox("newthing")
	activeLabelNewName.AddText("penguin")
	activeLabelNewName.AddText("snow")
	activeLabelNewName.AddText("GO")
	activeLabelNewName.AddText("debian")
	activeLabelNewName.AddText("RiscV")

	g2 = g.NewGroup("At X:")
	activeLabelNewX = g2.NewSpinner("tmp spinner", -1, 100)

	g2 = g.NewGroup("At Y:")
	activeLabelNewY = g2.NewSpinner("tmp spinner", -1, 100)

	g2 = g.NewGroup("bool B:")
	activeLabelNewB = g2.NewCheckbox("tmp bool")


	// common things that should work against each widget
	g = newB.NewGroup("common things")
	g.NewButton("Enable()", func () {
		activeWidget.Enable()
	})
	g.NewButton("Disable()", func () {
		activeWidget.Disable()
	})
	g.NewButton("Show()", func () {
		activeWidget.Show()
	})
	g.NewButton("Hide()", func () {
		activeWidget.Hide()
	})
	g.NewButton("Dump()", func () {
		activeWidget.Dump()
	})

	g = newB.NewGroup("add things")
	debugAddWidgetButton(g)
	g.NewLabel("experiments:")
	debugAddWidgetButtons(g)

	g = newB.NewGroup("change things")
	g.NewButton("AddText()", func () {
		activeWidget.AddText(activeLabelNewName.S)
		/*
		activeWidget.S = activeLabelNewName.S
		a := newAction(activeWidget, toolkit.AddText)
		sendAction(a)
		*/
	})
	g.NewButton("SetText()", func () {
		activeWidget.SetText(activeLabelNewName.S)
		/*
		activeWidget.S = activeLabelNewName.S
		a := newAction(activeWidget, toolkit.SetText)
		sendAction(a)
		*/
	})
	g.NewButton("Margin()", func () {
		activeWidget.Margin()
		/*
		a := newAction(activeWidget, toolkit.Margin)
		sendAction(a)
		*/
	})
	g.NewButton("Unmargin()", func () {
		activeWidget.Unmargin()
	})
	g.NewButton("Pad()", func () {
		activeWidget.Pad()
	})
	g.NewButton("Unpad()", func () {
		activeWidget.Unpad()
	})
	g.NewButton("Move(junk)", func () {
		log.Warn("gui.Node Move() not implemented yet")
	})
	g.NewButton("Delete()", func () {
		activeWidget.Delete(activeWidget)
	})

	g = newB.NewGroup("not working?")
	activeJunk = newB.NewGroup("junk:")
	activeJunk.NewLabel("test junk")

	if (activeWidget == nil) {
		setActiveWidget(myGui)
	}
}

func debugAddWidgetButtons(n *gui.Node) {
	n.NewButton("Dropdown", func () {
		a := activeWidget.NewDropdown("tmp dropdown")
		a.AddText("this is better than tcl/tk")
		a.AddText("make something for tim for qflow")
		a.AddText("and for riscv")
		a.Custom = func () {
			log.Log(true, "custom dropdown() a =", a.Name, a.S)
		}
	})
	n.NewButton("Combobox", func () {
		a := activeWidget.NewCombobox("tmp combobox")
		a.AddText("mirrors.wit.com")
		a.AddText("go.wit.com")
		a.Custom = func () {
			log.Log(true, "custom combobox() a =", a.Name, a.S)
		}
	})
	n.NewButton("Grid", func () {
		// Grid numbering by (X,Y)
		// -----------------------------
		// -- (1,1) -- (2,1) -- (3,1) --
		// -- (1,2) -- (2,1) -- (3,1) --
		// -----------------------------

		// SetDebug(true)
		debugGrid = activeWidget.NewGrid("tmp grid", 2, 3)
		debugGridLabel = debugGrid.NewLabel("mirrors.wit.com")
		/*
		debugGrid.SetNext(0,1)
		debugGrid.NewLabel("foo (0,1)")
		debugGrid.SetNext(1,1)
		debugGrid.NewLabel("foo (1,1)")
		debugGrid.SetNext(2,1)
		debugGrid.NewLabel("foo (2,1)")
		*/
		// SetDebug(false)
		DebugWidgetWindow(debugGrid)
	})
	n.NewButton("Image", func () {
		activeWidget.NewImage("image")
	})
	n.NewButton("Box(horizontal)", func () {
		a := activeWidget.NewBox("hBox", true)
		a.NewLabel("hBox")
		a.NewLabel("hBox 2")
	})
	n.NewButton("Box(vertical)", func () {
		a := activeWidget.NewBox("vBox", false)
		a.NewLabel("vBox")
		a.NewLabel("vBox 2")
	})
}

func debugAddWidgetButton(n *gui.Node) {
	activeLabelNewType = n.NewDropdown("tmp dropdown")
	activeLabelNewType.AddText("Window")
	activeLabelNewType.AddText("Tab")
	activeLabelNewType.AddText("Frame")
	activeLabelNewType.AddText("Grid")
	activeLabelNewType.AddText("Group")
	activeLabelNewType.AddText("Box")
	activeLabelNewType.AddText("Button")
	activeLabelNewType.AddText("Checkbox")
	activeLabelNewType.AddText("Dropdown")
	activeLabelNewType.AddText("Combobox")
	activeLabelNewType.AddText("Label")
	activeLabelNewType.AddText("Textbox")
	activeLabelNewType.AddText("Slider")
	activeLabelNewType.AddText("Spinner")
	activeLabelNewType.AddText("Image")
	activeLabelNewType.AddText("Area")
	activeLabelNewType.AddText("Form")
	activeLabelNewType.AddText("Font")
	activeLabelNewType.AddText("Color")
	activeLabelNewType.AddText("Dialog")

	n.NewButton("Add", func () {
		name :=  activeLabelNewName.S
		newX :=  activeLabelNewX.I
		newY :=  activeLabelNewY.I
		newB :=  activeLabelNewB.B

		if (newY == -1) {
			name = name + " (" + strconv.Itoa(activeWidget.NextW) + "," + strconv.Itoa(activeWidget.NextH) + ")"
		} else {
			activeWidget.SetNext(newX, newY)
			name = name + " (" + strconv.Itoa(newX) + "," + strconv.Itoa(newY) + ")"
		}
		log.Log(true, "New Name =", name)
		log.Log(true, "New Type =", activeLabelNewType.S)
		log.Log(true, "New X    =", newX)
		log.Log(true, "New Y    =", newY)
		log.Log(true, "activeWidget.NextW    =", activeWidget.NextW)
		log.Log(true, "activeWidget.NextH    =", activeWidget.NextH)
		log.Log(true, "Add() size (X,Y)", activeWidget.X, activeWidget.Y, "put next thing at (W,H) =", activeWidget.NextW, activeWidget.NextH)
		activeWidget.Dump()

		// activeWidget.X = newX
		// activeWidget.Y = newY

		switch activeLabelNewType.S {
		case "Grid":
			activeWidget.NewGrid(name, newX, newY)
		case "Group":
			activeWidget.NewGroup(name)
		case "Box":
			activeWidget.NewBox(name, newB)
		case "Button":
			activeWidget.NewButton(name, func () {
				log.Log(true, "got to button", name)
			})
		case "Checkbox":
			a := activeWidget.NewCheckbox(name)
			a.Custom = func () {
				log.Log(true, "custom checkox func a=", a.B)
			}
		case "Dropdown":
			a := activeWidget.NewDropdown(name)
			a.AddText(name + " yay")
			a.AddText(name + " haha")
			a.Custom = func () {
				log.Log(true, "WTF a=", a.B)
			}
		case "Combobox":
			a := activeWidget.NewCombobox(name)
			a.AddText(name + " foo")
			a.AddText(name + " bar")
		case "Label":
			activeWidget.NewLabel(name)
		case "Textbox":
			activeWidget.NewTextbox(name)
		case "Slider":
			activeWidget.NewSlider(name, newX, newY)
		case "Spinner":
			activeWidget.NewSpinner(name, newX, newY)
		default:
			log.Error(errors.New("make what type?"), activeLabelNewType.S)
		}
	})
}

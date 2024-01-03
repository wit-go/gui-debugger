package debugger

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

type LogSettings struct {
	ready	bool
	hidden	bool
	err	error
	name	string

	parent	*gui.Node // should be the root of the 'gui' package binary tree
	window	*gui.Node // our window for displaying the log package settings
	group	*gui.Node //
	grid	*gui.Node //
	checkbox *gui.Node
	label *gui.Node

}

func (ls *LogSettings) Set(b bool) {
	log.Set(ls.name, b)
	ls.checkbox.Set(b)
}

func NewLogFlag(p *gui.Node, name string) *LogSettings {
	ls := new(LogSettings)
	ls.parent = p
	ls.ready = false
	ls.name = name

	ls.checkbox = p.NewCheckbox(name)
	ls.label = p.NewLabel("Enable log." + name)
	ls.checkbox.Set(log.Get(name))
	ls.checkbox.Custom = func() {
		log.Set(name, ls.checkbox.B)
	}
	return ls
}

// Let's you toggle on and off the various types of debugging output
// These checkboxes should be in the same order as the are printed
func DebugFlags(n *gui.Node) {
	var newW, newB, g *gui.Node

	logGadgets := make(map[string]*LogSettings)

	newW = myGui.NewWindow("Debug Flags")
	newW.Custom = myGui.StandardClose

	newB = newW.NewBox("hBox", true)
	g = newB.NewGroup("Show").Pad()

	g.NewButton("log.SetTmp()", func () {
		log.SetTmp()
	})

	g.NewButton("log.All(true)", func () {
		for _, lf := range logGadgets {
			lf.Set(true)
		}
		log.All(true)
	})

	g.NewButton("log.All(false)", func () {
		for _, lf := range logGadgets {
			lf.Set(false)
		}
		log.All(false)
	})

	g.NewButton("Dump Flags", func () {
		// ShowDebugValues()
		log.ListFlags()
	})

	/*
	g.NewButton("All On", func () {
		SetDebug(true)
	})

	g.NewButton("All Off", func () {
		SetDebug(false)
	})
	*/

	g = newB.NewGroup("List")
	g = g.NewGrid("flags grid", 2, 2)

	logGadgets["INFO"] = NewLogFlag(g, "INFO")
	logGadgets["WARN"] = NewLogFlag(g, "WARN")
	logGadgets["SPEW"] = NewLogFlag(g, "SPEW")
	logGadgets["ERROR"] = NewLogFlag(g, "ERROR")
}

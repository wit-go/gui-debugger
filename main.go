package debugger

import 	(
	"os"

	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
	"go.wit.com/gui/gadgets/logsettings"
)

/*
	Creates a window helpful for debugging this package
*/

func DebugWindow(p *gui.Node) {
	if (me != nil) {
		log.Warn("Draw then Toggle() debuging window here")
		me.bugWin.Toggle()
		return
	}
	me = new(debuggerSettings)
	me.myGui = p

	me.bugWin = gadgets.NewBasicWindow(p,"go.wit.com/gui debug window")
	me.bugWin.Draw()
	DebugWindow2(me.bugWin.Box(), "Debug Tab")

	// initialize the log settings window (does not display it)
	me.myLS = logsettings.New(me.myGui)

	if ArgDebug() {
		log.SetTmp()
	}
}

func DebugWindow2(newB *gui.Node, title string) *gui.Node {
	var gr *gui.Node

//////////////////////// main debug things //////////////////////////////////
	gr = newB.NewGroup("Debugging Windows:")

	gr.NewButton("logging", func () {
		me.myLS.Toggle()
	})
	gr.NewButton("Widgets Window", func () {
		if me.widgets == nil {
			me.widgets = DebugWidgetWindow(me.myGui)
			return
		}
		me.widgets.Toggle()
	})

	gr.NewLabel("Force Quit:")

	gr.NewButton("os.Exit()", func () {
		os.Exit(0)
	})

//////////////////////// window debugging things //////////////////////////////////
	gr = newB.NewGroup("list things")

	gr.NewButton("List toolkits", func () {
		dropdownWindow(gr)
		bugWin.ListToolkits()
	})
	gr.NewButton("List Windows", func () {
		dropdownWindow(gr)
	})
	gr.NewButton("List Window Widgets", func () {
		dropdownWindowWidgets(gr)
	})

	gr = newB.NewGroup("more things")

	gr.NewButton("Node.ListChildren(true)", func () {
		if (activeWidget == nil) {
			activeWidget = bugWin
		}
		activeWidget.ListChildren(true)
	})

	gr.NewButton("test conc", func () {
		log.Log(WARN, "TODO: fix me")
		// makeConc()
	})

	gr.NewButton("List Plugins", func () {
		log.Log(WARN, "TODO: fix me")
		/*
		for _, aplug := range allPlugins {
			log.Log(true, "Loaded plugin:", aplug.name, aplug.filename)
		}
		*/
	})

	gr.NewButton("load toolkit 'gocui'", func () {
		bugWin.LoadToolkit("gocui")
	})

	gr.NewButton("load toolkit 'andlabs'", func () {
		bugWin.LoadToolkit("andlabs")
	})

	gr = newB.NewGroup("Learn GO")

	gr.NewButton("GO Language Internals", func () {
		if me.golang == nil {
			me.golang = DebugGolangWindow(me.myGui)
			return
		}
		log.Warn("going to toggle golang window")
		if me.golang.Ready() {
			me.golang.Toggle()
		}
	})
	gr.NewButton("GO Channels debug", func () {
		if me.gochan == nil {
			me.gochan = DebugGoChannels(me.myGui)
			return
		}
		log.Warn("going to toggle go channels window")
		if me.gochan.Ready() {
			me.gochan.Toggle()
		}
	})

	return newB
}

func dropdownWindow(p *gui.Node) {
	var mapWindows map[string]*gui.Node
	mapWindows = make(map[string]*gui.Node)

	dd := p.NewDropdown("Window Dropdown")
	dd.Custom = func() {
		name := dd.S
		activeWidget = mapWindows[name]
		setActiveWidget(activeWidget)
		log.Log(INFO, "The Window was set to", name)
	}
	log.Log(INFO, "dd =", dd)
	if (activeWidget == nil) {
		// the debug window doesn't exist yet so you can't display the change
		// TODO: make a fake binary tree for this(?)
		return
	}

	// var last = ""
	for _, child := range p.Children() {
		log.Log(INFO, "\t\t", child.GetName())
		dd.AddDropdownName(child.GetName())
		// last = child.Name
		mapWindows[child.GetName()] = child
		if (activeWidget == nil) {
			activeWidget = child
		}
	}
	// dd.SetDropdownName(last)
}

func dropdownWindowWidgets(p *gui.Node) {
	var mapWindows map[string]*gui.Node
	mapWindows = make(map[string]*gui.Node)

	dd := p.NewDropdown("Window Widgets Dropdown")
	dd.Custom = func() {
		name := dd.S
		activeWidget = mapWindows[name]
		setActiveWidget(activeWidget)
	}
	log.Log(INFO, "dd =", dd)

	// log.Log("dumpWidget() ", b, listChildrenDepth, defaultPadding, n.id, info)

	var addDropdowns func (*gui.Node)
	addDropdowns = func (n *gui.Node) {
		// s := n.dumpWidget(true)
		s := n.GetName()
		dd.AddDropdownName(s)
		mapWindows[s] = n

		for _, child := range n.Children() {
			addDropdowns(child)
		}
	}

	// list everything in the binary tree
	addDropdowns(bugWin)
}

func setActiveWidget(w *gui.Node) {
	log.Warn("TODO: setActiveWidget()")
}

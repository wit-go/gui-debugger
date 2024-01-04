package debugger

import 	(
	"os"

	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

/*
	Creates a window helpful for debugging this package
*/

func DebugWindow(p *gui.Node) {
	myGui = p
	bugWin = myGui.NewWindow("go.wit.com/gui debug window")
	bugWin.StandardClose()
	bugTab = DebugWindow2(bugWin, "Debug Tab")
	bugTab.StandardClose()
	if ArgDebug() {
		log.SetTmp()
	}
}

func DebugWindow2(n *gui.Node, title string) *gui.Node {
	var newW, newB, gr *gui.Node
	// var logSettings *gadgets.LogSettings

	// time.Sleep(1 * time.Second)
	newW = n.NewWindow(title)

	newB = newW.NewBox("hBox", true)

//////////////////////// main debug things //////////////////////////////////
	gr = newB.NewGroup("Debugging Windows:")

	gr.NewButton("logging", func () {
		DebugFlags(myGui)
	})
	gr.NewButton("Debug Widgets", func () {
		DebugWidgetWindow(myGui)
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
		log.Log(true, "TODO: fix me")
		// makeConc()
	})

	gr.NewButton("List Plugins", func () {
		log.Log(true, "TODO: fix me")
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
		DebugGolangWindow(myGui)
	})
	gr.NewButton("GO Channels debug", func () {
		DebugGoChannels(myGui)
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
		log.Log(true, "The Window was set to", name)
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
}

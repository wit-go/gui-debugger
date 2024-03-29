package debugger

import (
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
	"go.wit.com/gui/gadgets/logsettings"
)

var me *debuggerSettings

type debuggerSettings struct {
	ready	bool
	hidden	bool
	err	error

	myGui *gui.Node

	bugWin *gadgets.BasicWindow
	widgets *gadgets.BasicWindow
	golang *gadgets.BasicWindow
	gochan *gadgets.BasicWindow

	myLS *logsettings.LogSettings

	mapWindows map[string]*gui.Node // tracks all windows that exist
}

var bugWin *gui.Node
/*
// main debugging window
var bugTab *gui.Node
var myGui *gui.Node
*/

// global var for checking to see if this
// window/tab for debugging a widget exists
// check the binary tree instead (?) for a window called "Widgets" (bad idea)
var bugWidget *gui.Node

// the widget all these actions are run against
var activeWidget *gui.Node

// for testing move, this is the node things are put on
var activeJunk *gui.Node

// the label where the user can see which widget is active
var activeLabel *gui.Node
var activeLabelType *gui.Node
var activeLabelNewName *gui.Node
var activeLabelNewType *gui.Node
var activeLabelNewX *gui.Node
var activeLabelNewY *gui.Node
var activeLabelNewB *gui.Node

// tmp junk
var debugGrid *gui.Node
var debugGridLabel *gui.Node
var debugWidgetBut1, debugWidgetBut2 *gui.Node

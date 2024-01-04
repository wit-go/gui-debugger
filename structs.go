package debugger

import (
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets/logsettings"
)

// main debugging window
var myGui *gui.Node
var bugWin *gui.Node
var bugTab *gui.Node

var myLS *logsettings.LogSettings

var mapWindows map[string]*gui.Node // tracks all windows that exist

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

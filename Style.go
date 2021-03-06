package imgui

// #include "StyleWrapper.h"
import "C"

// StyleVarID identifies a style variable in the UI style.
type StyleVarID int

const (
	// StyleVarAlpha is a float
	StyleVarAlpha StyleVarID = iota
	// StyleVarWindowPadding is a Vec2
	StyleVarWindowPadding
	// StyleVarWindowRounding is a float
	StyleVarWindowRounding
	// StyleVarWindowBorderSize is a float
	StyleVarWindowBorderSize
	// StyleVarWindowMinSize is a Vec2
	StyleVarWindowMinSize
	// StyleVarWindowTitleAlign is a Vec2
	StyleVarWindowTitleAlign
	// StyleVarChildRounding is a float
	StyleVarChildRounding
	// StyleVarChildBorderSize is a float
	StyleVarChildBorderSize
	// StyleVarPopupRounding is a float
	StyleVarPopupRounding
	// StyleVarPopupBorderSize is a float
	StyleVarPopupBorderSize
	// StyleVarFramePadding is a Vec2
	StyleVarFramePadding
	// StyleVarFrameRounding is a float
	StyleVarFrameRounding
	// StyleVarFrameBorderSize is a float
	StyleVarFrameBorderSize
	// StyleVarItemSpacing is a Vec2
	StyleVarItemSpacing
	// StyleVarItemInnerSpacing is a Vec2
	StyleVarItemInnerSpacing
	// StyleVarIndentSpacing is a float
	StyleVarIndentSpacing
	// StyleVarScrollbarSize is a float
	StyleVarScrollbarSize
	// StyleVarScrollbarRounding is a float
	StyleVarScrollbarRounding
	// StyleVarGrabMinSize is a float
	StyleVarGrabMinSize
	// StyleVarGrabRounding is a float
	StyleVarGrabRounding
	// StyleVarTabRounding is a float
	StyleVarTabRounding
	// StyleVarButtonTextAlign is a Vec2
	StyleVarButtonTextAlign
)

// StyleColorID identifies a color in the UI style.
type StyleColorID int

// StyleColor identifier
const (
	StyleColorText StyleColorID = iota
	StyleColorTextDisabled
	StyleColorWindowBg
	StyleColorChildBg
	StyleColorPopupBg
	StyleColorBorder
	StyleColorBorderShadow
	StyleColorFrameBg
	StyleColorFrameBgHovered
	StyleColorFrameBgActive
	StyleColorTitleBg
	StyleColorTitleBgActive
	StyleColorTitleBgCollapsed
	StyleColorMenuBarBg
	StyleColorScrollbarBg
	StyleColorScrollbarGrab
	StyleColorScrollbarGrabHovered
	StyleColorScrollbarGrabActive
	StyleColorCheckMark
	StyleColorSliderGrab
	StyleColorSliderGrabActive
	StyleColorButton
	StyleColorButtonHovered
	StyleColorButtonActive
	StyleColorHeader
	StyleColorHeaderHovered
	StyleColorHeaderActive
	StyleColorSeparator
	StyleColorSeparatorHovered
	StyleColorSeparatorActive
	StyleColorResizeGrip
	StyleColorResizeGripHovered
	StyleColorResizeGripActive
	StyleColorTab
	StyleColorTabHovered
	StyleColorTabActive
	StyleColorTabUnfocused
	StyleColorTabUnfocusedActive
	StyleColorPlotLines
	StyleColorPlotLinesHovered
	StyleColorPlotHistogram
	StyleColorPlotHistogramHovered
	StyleColorTextSelectedBg
	StyleColorDragDropTarget
	StyleColorNavHighlight          // Gamepad/keyboard: current highlighted item
	StyleColorNavWindowingHighlight // Highlight window when using CTRL+TAB
	StyleColorNavWindowingDarkening // Darken/colorize entire screen behind the CTRL+TAB window list, when active
	StyleColorModalWindowDarkening  // Darken/colorize entire screen behind a modal window, when one is active
)

// Style describes the overall graphical representation of the user interface.
type Style uintptr

func (style Style) handle() C.IggGuiStyle {
	return C.IggGuiStyle(style)
}

// WindowPadding is the padding within a window.
func (style Style) WindowPadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetWindowPadding(style.handle(), valueArg)
	valueFin()
	return value
}

// WindowMinSize is the minimum window size. This is a global setting.
func (style Style) WindowMinSize() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetWindowMinSize(style.handle(), valueArg)
	valueFin()
	return value
}

// WindowTitleAlign is the alignment for title bar text.
// Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func (style Style) WindowTitleAlign() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetWindowTitleAlign(style.handle(), valueArg)
	valueFin()
	return value
}

// FramePadding is the padding within a framed rectangle (used by most widgets).
func (style Style) FramePadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetFramePadding(style.handle(), valueArg)
	valueFin()
	return value
}

// ItemSpacing is the horizontal and vertical spacing between widgets/lines.
func (style Style) ItemSpacing() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetItemSpacing(style.handle(), valueArg)
	valueFin()
	return value
}

// ItemInnerSpacing is the horizontal and vertical spacing between elements of
// a composed widget (e.g. a slider and its label).
func (style Style) ItemInnerSpacing() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetItemInnerSpacing(style.handle(), valueArg)
	valueFin()
	return value
}

// TouchExtraPadding allows you to expand reactive bounding box for touch-based
// system where touch position is not accurate enough.
// Unfortunately we don't sort widgets so priority on overlap will always be
// given to the first widget. So don't grow this too much!
func (style Style) TouchExtraPadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetTouchExtraPadding(style.handle(), valueArg)
	valueFin()
	return value
}

// ButtonTextAlign is the alignment of button text when button is larger than
// text. Defaults to (0.5f,0.5f) for horizontally+vertically centered.
func (style Style) ButtonTextAlign() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetButtonTextAlign(style.handle(), valueArg)
	valueFin()
	return value
}

// DisplayWindowPadding - Window position are clamped to be visible within the
// display area by at least this amount. Only applies to regular windows.
func (style Style) DisplayWindowPadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetDisplayWindowPadding(style.handle(), valueArg)
	valueFin()
	return value
}

// DisplaySafeAreaPadding - If you cannot see the
// edges of your screen (e.g. on a TV) increase the safe area padding.
// Apply to popups/tooltips as well regular windows.
// NB: Prefer configuring your TV sets correctly!
func (style Style) DisplaySafeAreaPadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetDisplaySafeAreaPadding(style.handle(), valueArg)
	valueFin()
	return value
}

// SetColor sets a color value of the UI style.
func (style Style) SetColor(id StyleColorID, value Vec4) {
	valueArg, _ := value.wrapped()
	C.iggStyleSetColor(style.handle(), C.int(id), valueArg)
}

// ScaleAllSizes applies a scaling factor to all sizes.
// To scale your entire UI (e.g. if you want your app to use High DPI or generally be DPI aware) you may use this helper function.
// Scaling the fonts is done separately and is up to you.
//
// Important: This operation is lossy because all sizes are rounded to integer.
// If you need to change your scale multiples, call this over a freshly initialized style rather than scaling multiple times.
func (style Style) ScaleAllSizes(scale float32) {
	C.iggStyleScaleAllSizes(style.handle(), C.float(scale))
}

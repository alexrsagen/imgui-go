#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern void iggStyleGetWindowPadding(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetWindowMinSize(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetWindowTitleAlign(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetFramePadding(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetItemSpacing(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetItemInnerSpacing(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetTouchExtraPadding(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetButtonTextAlign(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetDisplayWindowPadding(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetDisplaySafeAreaPadding(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleSetColor(IggGuiStyle handle, int colorID, IggVec4 const *value);
extern void iggStyleScaleAllSizes(IggGuiStyle handle, float scale);

#ifdef __cplusplus
}
#endif

#include "imguiWrappedHeader.h"
#include "StyleWrapper.h"
#include "WrapperConverter.h"

void iggStyleGetWindowPadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->WindowPadding);
}

void iggStyleGetWindowMinSize(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->WindowMinSize);
}

void iggStyleGetWindowTitleAlign(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->WindowTitleAlign);
}

void iggStyleGetFramePadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->FramePadding);
}

void iggStyleGetItemSpacing(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->ItemSpacing);
}

void iggStyleGetItemInnerSpacing(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->ItemInnerSpacing);
}

void iggStyleGetTouchExtraPadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->TouchExtraPadding);
}

void iggStyleGetButtonTextAlign(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->ButtonTextAlign);
}

void iggStyleGetDisplayWindowPadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->DisplayWindowPadding);
}

void iggStyleGetDisplaySafeAreaPadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->DisplaySafeAreaPadding);
}

void iggStyleSetColor(IggGuiStyle handle, int colorID, IggVec4 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   if ((colorID >= 0) && (colorID < ImGuiCol_COUNT))
   {
      importValue(style->Colors[colorID], *value);
   }
}

void iggStyleScaleAllSizes(IggGuiStyle handle, float scale)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->ScaleAllSizes(scale);
}

// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/axzilla/goilerplate/pkg/icons"
)

// DropdownMenuItem represents an item in the dropdown menu
type DropdownMenuItem struct {
	Label     string             // The text to display for the menu item
	Value     string             // The value associated with the menu item (for non-link items)
	Href      string             // The URL to navigate to if the item is a link
	Target    string             // The target attribute for link items (e.g., "_blank" for new tab)
	IconLeft  templ.Component    // An optional icon component to display on the left side of the item
	IconRight templ.Component    // An optional icon component to display on the right side of the item
	SubItems  []DropdownMenuItem // A slice of DropdownMenuItems for nested submenus
	Disabled  bool               // Whether the item is disabled or not
	// Attributes allows passing additional HTML attributes to the accordion container element.
	// Default: nil
	Attributes templ.Attributes
}

// DropdownMenuProps defines the properties for the DropdownMenu component
type DropdownMenuProps struct {
	Items    []DropdownMenuItem // The list of items to display in the dropdown menu
	Trigger  templ.Component    // An optional custom trigger component (default is a button if not provided)
	Class    string             // Additional CSS classes to apply to the root element
	Position string             // The preferred position of the dropdown ("left", "right", "top", or "bottom")
}

// renderMenuItem is a helper function to render a single menu item
// It handles rendering of regular items, links, and submenus recursively
func renderMenuItem(item DropdownMenuItem, index int, depth int) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if len(item.SubItems) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"relative group\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 = []any{
				"w-full text-left flex items-center justify-between px-4 py-2 text-sm",
				templ.KV("text-foreground hover:bg-accent hover:text-accent-foreground", !item.Disabled),
				templ.KV("text-muted-foreground cursor-not-allowed", item.Disabled),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" role=\"menuitem\" tabindex=\"-1\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("menu-item-%d", index))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 44, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" aria-haspopup=\"true\" aria-expanded=\"false\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.Disabled {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" disabled")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, item.Attributes)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><span class=\"flex items-center gap-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconLeft != nil {
				templ_7745c5c3_Err = item.IconLeft.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(item.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 54, Col: 17}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconRight != nil {
				templ_7745c5c3_Err = item.IconRight.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = icons.ChevronRight(icons.IconProps{Size: "16"}).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if depth < 3 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"absolute left-full top-0 hidden group-hover:block\"><div class=\"py-1 bg-popover rounded-md shadow-lg border border-border\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for subIndex, subItem := range item.SubItems {
					templ_7745c5c3_Err = renderMenuItem(subItem, subIndex, depth+1).Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else if item.Href != "" {
			var templ_7745c5c3_Var6 = []any{
				"block px-4 py-2 text-sm flex items-center",
				templ.KV("text-foreground hover:bg-accent hover:text-accent-foreground", !item.Disabled),
				templ.KV("text-muted-foreground cursor-not-allowed", item.Disabled),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var6...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 templ.SafeURL = templ.SafeURL(item.Href)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var7)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" target=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(item.Target)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 75, Col: 23}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var6).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" role=\"menuitem\" tabindex=\"-1\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("menu-item-%d", index))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 83, Col: 42}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, item.Attributes)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><span class=\"flex items-center gap-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconLeft != nil {
				templ_7745c5c3_Err = item.IconLeft.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(item.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 90, Col: 16}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconRight != nil {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"ml-auto\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = item.IconRight.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			var templ_7745c5c3_Var12 = []any{
				"w-full text-left flex items-center justify-between px-4 py-2 text-sm",
				templ.KV("text-foreground hover:bg-accent hover:text-accent-foreground", !item.Disabled),
				templ.KV("text-muted-foreground cursor-not-allowed", item.Disabled),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var12...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var12).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" role=\"menuitem\" tabindex=\"-1\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("menu-item-%d", index))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 107, Col: 42}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.Disabled {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" disabled")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, item.Attributes)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><span class=\"flex items-center gap-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconLeft != nil {
				templ_7745c5c3_Err = item.IconLeft.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(item.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 115, Col: 16}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if item.IconRight != nil {
				templ_7745c5c3_Err = item.IconRight.Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

// DropdownMenu renders an enhanced dropdown menu component
// It can be customized with various styles, positions, and behaviors.
//
// Usage:
//
//	@components.DropdownMenu(components.DropdownMenuProps{
//	  Trigger: components.Button(components.ButtonProps{
//	    Text: "Open Menu",
//	    IconRight: components.Icon(components.IconProps{Name: "chevron-down"}),
//	  }),
//	  Position: "bottom",
//	  Items: []components.DropdownMenuItem{
//	    {Label: "Option 1", Value: "opt1"},
//	    {Label: "Option 2", Value: "opt2", IconLeft: components.Icon(components.IconProps{Name: "settings"})},
//	    {Label: "Submenu", SubItems: []components.DropdownMenuItem{
//	      {Label: "Sub Option 1", Value: "sub1"},
//	      {Label: "Sub Option 2", Value: "sub2"},
//	    }},
//	  },
//	})
//
// Props:
//   - Items: Slice of DropdownMenuItem, defining the content of the dropdown
//   - Trigger: Custom component to trigger the dropdown (optional, default is a basic button)
//   - Class: Additional CSS classes to apply to the root element
//   - Position: Preferred position of the dropdown ("left", "right", "top", or "bottom")
//   - Attributes: Additional HTML attributes to apply to the accordion container element. Default: nil
//
// Features:
//   - Supports nested submenus up to 3 levels deep
//   - Automatically adjusts position based on available space
//   - Supports custom icons for menu items
//   - Handles disabled states for menu items
//   - Uses Tailwind CSS for styling, including dark mode support
//   - Implements keyboard navigation and ARIA attributes for accessibility
func DropdownMenu(props DropdownMenuProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var16 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var16 == nil {
			templ_7745c5c3_Var16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var17 = []any{"relative inline-block text-left", props.Class}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var17...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{ \n\t\t\topen: false,\n\t\t\tposition: $el.dataset.position,\n\t\t\tverticalPosition: &#39;bottom&#39;,\n\t\t\tupdatePosition() {\n\t\t\t\tconst menu = this.$refs.menu;\n\t\t\t\tconst rect = menu.getBoundingClientRect();\n\t\t\t\tconst viewportWidth = window.innerWidth || document.documentElement.clientWidth;\n\t\t\t\tconst viewportHeight = window.innerHeight || document.documentElement.clientHeight;\n\t\t\t\t\n\t\t\t\tif (this.position === &#39;left&#39; &amp;&amp; rect.left &lt; 0) {\n\t\t\t\t\tthis.position = &#39;right&#39;;\n\t\t\t\t} else if (this.position !== &#39;left&#39; &amp;&amp; rect.right &gt; viewportWidth) {\n\t\t\t\t\tthis.position = &#39;left&#39;;\n\t\t\t\t}\n\n\t\t\t\tif (this.verticalPosition === &#39;bottom&#39; &amp;&amp; rect.bottom &gt; viewportHeight) {\n\t\t\t\t\tthis.verticalPosition = &#39;top&#39;;\n\t\t\t\t} else if (this.verticalPosition === &#39;top&#39; &amp;&amp; rect.top &lt; 0) {\n\t\t\t\t\tthis.verticalPosition = &#39;bottom&#39;;\n\t\t\t\t}\n\t\t\t}\n\t\t}\" @resize.window=\"if (open) updatePosition()\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var18 string
		templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var17).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-position=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var19 string
		templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(props.Position)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 186, Col: 32}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div @click=\"open = !open; if(open) $nextTick(() =&gt; updatePosition())\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Trigger != nil {
			templ_7745c5c3_Err = props.Trigger.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = Button(ButtonProps{
				Text:      "Options",
				Variant:   "outline",
				IconRight: icons.ChevronDown(icons.IconProps{Size: "16"}),
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var20 = []any{
			"absolute z-50 mt-2 w-56 rounded-md shadow-lg bg-popover ring-1 ring-black ring-opacity-5 focus:outline-none",
			"border border-border",
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var20...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-ref=\"menu\" x-show=\"open\" @click.away=\"open = false\" @keydown.escape.window=\"open = false\" x-transition:enter=\"transition ease-out duration-100\" x-transition:enter-start=\"transform opacity-0 scale-95\" x-transition:enter-end=\"transform opacity-100 scale-100\" x-transition:leave=\"transition ease-in duration-75\" x-transition:leave-start=\"transform opacity-100 scale-100\" x-transition:leave-end=\"transform opacity-0 scale-95\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var21 string
		templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var20).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/dropdown_menu.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" :class=\"{\n\t\t\t\t&#39;right-0&#39;: position === &#39;left&#39;,\n\t\t\t\t&#39;left-0&#39;: position !== &#39;left&#39;,\n\t\t\t\t&#39;bottom-full mb-2&#39;: verticalPosition === &#39;top&#39;,\n\t\t\t\t&#39;top-full mt-2&#39;: verticalPosition === &#39;bottom&#39;\n\t\t\t}\" role=\"menu\" aria-orientation=\"vertical\" aria-labelledby=\"dropdown-menu-button\" tabindex=\"-1\"><div class=\"py-1\" role=\"none\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for index, item := range props.Items {
			templ_7745c5c3_Err = renderMenuItem(item, index, 0).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate

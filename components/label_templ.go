// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.856
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/axzilla/templui/utils"

type LabelProps struct {
	ID            string
	Class         string
	Attributes    templ.Attributes
	For           string
	Error         string
	DisabledClass string
}

func Label(props ...LabelProps) templ.Component {
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
		var p LabelProps
		if len(props) > 0 {
			p = props[0]
		}
		var templ_7745c5c3_Var2 = []any{
			utils.TwMerge(
				"text-sm font-medium leading-none inline-block",
				utils.If(len(p.Error) > 0, "text-destructive"),
				p.Class,
			),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<label")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.ID != "" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, " id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/label.templ`, Line: 21, Col: 12}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if p.For != "" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, " for=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(p.For)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/label.templ`, Line: 24, Col: 14}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, " class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/label.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.DisabledClass != "" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, " data-disabled-style=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(p.DisabledClass)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/label.templ`, Line: 34, Col: 40}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, " data-disabled-style=\"opacity-50 cursor-not-allowed\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, p.Attributes)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, ">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "</label>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func LabelScript() templ.Component {
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
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		handle := templ.NewOnceHandle()
		templ_7745c5c3_Var8 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "<script defer nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/label.templ`, Line: 47, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "\">\n\t\t\t// Inside LabelScript\n\t\t\tdocument.addEventListener('DOMContentLoaded', function() {\n\t\t\t  // Label click handler - automatically works for standard inputs\n\t\t\t  // For buttons, we only need to add a click handler\n\t\t\t  document.querySelectorAll('label[for]').forEach(function(label) {\n\t\t\t\tlabel.addEventListener('click', function(e) {\n\t\t\t\t  // Get the for-ID at the time of the click, not during initialization\n\t\t\t\t  const forId = this.getAttribute('for');\n\t\t\t\t  \n\t\t\t\t  // Find the target element with the corresponding ID\n\t\t\t\t  const targetElement = document.getElementById(forId);\n\t\t\t\t  \n\t\t\t\t  // If the target element is a button, process the click\n\t\t\t\t  if (targetElement && targetElement.tagName === 'BUTTON') {\n\t\t\t\t\t// Prevent the default label behavior\n\t\t\t\t\te.preventDefault();\n\t\t\t\t\t\n\t\t\t\t\t// Stop event propagation to avoid duplicate clicks\n\t\t\t\t\te.stopPropagation();\n\t\t\t\t\t\n\t\t\t\t\tif (!targetElement.disabled) {\n\t\t\t\t\t  // Ensure the button receives focus\n\t\t\t\t\t  targetElement.focus();\n\t\t\t\t\t  \n\t\t\t\t\t  // Slightly delay the click to avoid event conflicts\n\t\t\t\t\t  setTimeout(function() {\n\t\t\t\t\t\t// Create a new click event for the button\n\t\t\t\t\t\tconst clickEvent = new MouseEvent('click', {\n\t\t\t\t\t\t  bubbles: true,\n\t\t\t\t\t\t  cancelable: true,\n\t\t\t\t\t\t  view: window\n\t\t\t\t\t\t});\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Dispatch the click event on the button\n\t\t\t\t\t\ttargetElement.dispatchEvent(clickEvent);\n\t\t\t\t\t  }, 10);\n\t\t\t\t\t}\n\t\t\t\t  }\n\t\t\t\t});\n\t\t\t\t\n\t\t\t\t// Disabled style handler\n\t\t\t\tfunction updateLabelStyle(label) {\n\t\t\t\t  const forId = label.getAttribute('for');\n\t\t\t\t  const targetElement = document.getElementById(forId);\n\t\t\t\t  \n\t\t\t\t  if (targetElement && targetElement.disabled) {\n\t\t\t\t\t// Retrieve the custom disabled style or use the default\n\t\t\t\t\tconst disabledStyle = label.getAttribute('data-disabled-style') || 'opacity-50 cursor-not-allowed';\n\t\t\t\t\tdisabledStyle.split(' ').forEach(function(className) {\n\t\t\t\t\t  if (className) {\n\t\t\t\t\t\tlabel.classList.add(className);\n\t\t\t\t\t  }\n\t\t\t\t\t});\n\t\t\t\t  } else {\n\t\t\t\t\t// Remove the disabled style\n\t\t\t\t\tconst disabledStyle = label.getAttribute('data-disabled-style') || 'opacity-50 cursor-not-allowed';\n\t\t\t\t\tdisabledStyle.split(' ').forEach(function(className) {\n\t\t\t\t\t  if (className) {\n\t\t\t\t\t\tlabel.classList.remove(className);\n\t\t\t\t\t  }\n\t\t\t\t\t});\n\t\t\t\t  }\n\t\t\t\t}\n\t\t\t\t\n\t\t\t\t// Update the initial style\n\t\t\t\tupdateLabelStyle(label);\n\t\t\t\t\n\t\t\t\t// Create a MutationObserver for this specific label\n\t\t\t\tconst forId = label.getAttribute('for');\n\t\t\t\tconst targetElement = document.getElementById(forId);\n\t\t\t\t\n\t\t\t\tif (targetElement) {\n\t\t\t\t  const observer = new MutationObserver(function() {\n\t\t\t\t\tupdateLabelStyle(label);\n\t\t\t\t  });\n\t\t\t\t  \n\t\t\t\t  observer.observe(targetElement, { \n\t\t\t\t\tattributes: true, \n\t\t\t\t\tattributeFilter: ['disabled'] \n\t\t\t\t  });\n\t\t\t\t}\n\t\t\t  });\n\t\t\t});\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = handle.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate

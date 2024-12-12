// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/axzilla/templui/pkg/icons"
)

type ToastProps struct {
	Message     string // Message to display
	Type        string // Type of the toast (default, success, error, warning, info)
	Position    string // Position of the toast (top-right, top-left, top-center, bottom-right, bottom-left, bottom-center)
	Duration    int    // Duration in milliseconds
	Dismissible bool   // Show dismiss button
	Size        string // Size of the toast (sm, md, lg)
	Icon        bool   // Show icon
}

// Flexible toast component for notifications and feedback.
func Toast(props ToastProps) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(`{
            show: true,
            message: "%s",
            type: "%s",
            position: "%s",
            duration: %d,
            dismissible: %t,
            size: "%s",
            timer: null,
            startTimer() {
                if (this.duration <= 0) return;
                this.timer = setTimeout(() => this.show = false, this.duration);
            },
            pauseTimer() {
                if (this.timer) clearTimeout(this.timer);
                const progress = this.$refs.progress;
                if (progress) {
                    const width = progress.getBoundingClientRect().width;
                    const total = progress.parentElement.getBoundingClientRect().width;
                    this.duration = (width / total) * this.duration;
                    progress.style.transition = "none";
                    progress.style.width = width + "px";
                }
            },
            resumeTimer() {
                const progress = this.$refs.progress;
                if (progress) {
                    progress.style.transition = "width " + this.duration + "ms linear";
                    progress.style.width = "0";
                    this.startTimer();
                }
            }
        }`, props.Message, props.Type, props.Position, props.Duration, props.Dismissible, props.Size))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toast.templ`, Line: 53, Col: 101}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"startTimer()\" @mouseenter=\"pauseTimer()\" @mouseleave=\"resumeTimer()\" x-show=\"show\" x-transition:enter=\"transition ease-out duration-300\" x-transition:enter-start=\"opacity-0 translate-y-4\" x-transition:enter-end=\"opacity-100 translate-y-0\" x-transition:leave=\"transition ease-in duration-200\" x-transition:leave-start=\"opacity-100 translate-y-0\" x-transition:leave-end=\"opacity-0 translate-y-4\" @click=\"if(dismissible) show = false\" class=\"z-50 fixed pointer-events-auto\" :class=\"{\n            &#39;top-4 right-4&#39;: position === &#39;top-right&#39;,\n            &#39;top-4 left-4&#39;: position === &#39;top-left&#39;,\n            &#39;top-4 left-1/2 -translate-x-1/2&#39;: position === &#39;top-center&#39;,\n            &#39;bottom-4 right-4&#39;: position === &#39;bottom-right&#39;,\n            &#39;bottom-4 left-4&#39;: position === &#39;bottom-left&#39;,\n            &#39;bottom-4 left-1/2 -translate-x-1/2&#39;: position === &#39;bottom-center&#39;,\n            &#39;w-72&#39;: size === &#39;sm&#39;,\n            &#39;w-96&#39;: size === &#39;md&#39;,\n            &#39;w-[30rem]&#39;: size === &#39;lg&#39;\n        }\"><div class=\"bg-primary-foreground rounded-lg shadow-sm border pt-5 pb-4 px-4 flex items-center justify-center relative overflow-hidden\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Duration > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"absolute top-0 left-0 right-0 h-1\"><div x-ref=\"progress\" class=\"absolute inset-0\" :class=\"{\n                            &#39;bg-green-500&#39;: type === &#39;success&#39;,\n                            &#39;bg-red-500&#39;: type === &#39;error&#39;,\n                            &#39;bg-yellow-500&#39;: type === &#39;warning&#39;,\n                            &#39;bg-blue-500&#39;: type === &#39;info&#39;,\n                            &#39;bg-gray-500&#39;: type === &#39;default&#39;\n                        }\" x-init=\"\n                            $el.style.transition = &#39;width &#39; + duration + &#39;ms linear&#39;;\n                            $el.style.width = &#39;100%&#39;;\n                            setTimeout(() =&gt; {\n                                $el.style.width = &#39;0%&#39;;\n                            }, 10)\n                        \"></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if props.Icon {
			if props.Type == "success" {
				templ_7745c5c3_Err = icons.CircleCheck(icons.IconProps{Size: "18", Class: "text-green-500 mr-3"}).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if props.Type == "error" {
				templ_7745c5c3_Err = icons.CircleX(icons.IconProps{Size: "18", Class: "text-red-500 mr-3"}).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if props.Type == "warning" {
				templ_7745c5c3_Err = icons.TriangleAlert(icons.IconProps{Size: "18", Class: "text-yellow-500 mr-3"}).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if props.Type == "info" {
				templ_7745c5c3_Err = icons.Info(icons.IconProps{Size: "18", Class: "text-blue-500 mr-3"}).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex-1\" x-text=\"message\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Dismissible {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button @click.stop=\"show = false\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = icons.X(icons.IconProps{
				Size:  "18",
				Class: "opacity-75 hover:opacity-100",
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate

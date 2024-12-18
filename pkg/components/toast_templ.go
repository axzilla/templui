// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/axzilla/templui/pkg/icons"
	"strconv"
)

type ToastProps struct {
	Message     string // Message to display
	Type        string // Type of the toast
	Position    string // Position of the toast
	Duration    int    // Duration in milliseconds
	Dismissible bool   // Show dismiss button
	Size        string // Size of the toast
	Icon        bool   // Show icon
}

func (p ToastProps) withDefaults() ToastProps {
	if p.Message == "" {
		p.Message = "Notification"
	}
	if p.Type == "" {
		p.Type = "default"
	}
	if p.Position == "" {
		p.Position = "bottom-right"
	}
	if p.Duration == 0 {
		p.Duration = 3000
	}
	if p.Size == "" {
		p.Size = "md"
	}
	return p
}

// Flexible toast component for notifications and feedback.
func ToastScript() templ.Component {
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
		handler := templ.NewOnceHandle()
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script defer nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toast.templ`, Line: 41, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">\t\n            document.addEventListener('alpine:init', () => {\n                Alpine.data('toast', () => ({\n                    show: true,\n                    duration: 3000,\n                    timer: null,\n\n                    init() {\n                        this.duration = parseInt(this.$el.dataset.duration || 0);\n                        this.startTimer();\n\n                        if (this.duration > 0) {\n                            const progress = this.$refs.progress;\n                            if (progress) {\n                                progress.style.transition = `width ${this.duration}ms linear`;\n                                progress.style.width = '100%';\n                                setTimeout(() => {\n                                    progress.style.width = '0%';\n                                }, 10);\n                            }\n                        }\n                    },\n\n                    startTimer() {\n                        if (this.duration <= 0) return;\n                        this.timer = setTimeout(() => {\n                            this.show = false;\n                        }, this.duration);\n                    },\n\n                    pauseTimer() {\n                        if (this.timer) clearTimeout(this.timer);\n                        const progress = this.$refs.progress;\n                        if (progress) {\n                            const width = progress.getBoundingClientRect().width;\n                            const total = progress.parentElement.getBoundingClientRect().width;\n                            this.duration = (width / total) * this.duration;\n                            progress.style.transition = \"none\";\n                            progress.style.width = width + \"px\";\n                        }\n                    },\n\n                    resumeTimer() {\n                        const progress = this.$refs.progress;\n                        if (progress) {\n                            progress.style.transition = \"width \" + this.duration + \"ms linear\";\n                            progress.style.width = \"0\";\n                            this.startTimer();\n                        }\n                    },\n\n                    dismissToast() {\n                        this.show = false;\n                    }\n                }))\n            })\n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = handler.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

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
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		props = props.withDefaults()
		var templ_7745c5c3_Var5 = []any{
			"z-50 fixed pointer-events-auto",
			templ.KV("top-4 right-4", props.Position == "top-right"),
			templ.KV("top-4 left-4", props.Position == "top-left"),
			templ.KV("top-4 left-1/2 -translate-x-1/2", props.Position == "top-center"),
			templ.KV("bottom-4 right-4", props.Position == "bottom-right"),
			templ.KV("bottom-4 left-4", props.Position == "bottom-left"),
			templ.KV("bottom-4 left-1/2 -translate-x-1/2", props.Position == "bottom-center"),
			templ.KV("w-72", props.Size == "sm"),
			templ.KV("w-96", props.Size == "md"),
			templ.KV("w-[30rem]", props.Size == "lg"),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var5...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div data-duration=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(props.Duration))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toast.templ`, Line: 105, Col: 46}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-data=\"toast\" @mouseenter=\"pauseTimer\" @mouseleave=\"resumeTimer\" x-show=\"show\" x-transition:enter=\"transition ease-out duration-300\" x-transition:enter-start=\"opacity-0 translate-y-4\" x-transition:enter-end=\"opacity-100 translate-y-0\" x-transition:leave=\"transition ease-in duration-200\" x-transition:leave-start=\"opacity-100 translate-y-0\" x-transition:leave-end=\"opacity-0 translate-y-4\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var5).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toast.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"bg-primary-foreground rounded-lg shadow-sm border pt-5 pb-4 px-4 flex items-center justify-center relative overflow-hidden\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Duration > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"absolute top-0 left-0 right-0 h-1\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 = []any{
				"absolute inset-0",
				templ.KV("bg-gray-500", props.Type == "default"),
				templ.KV("bg-green-500", props.Type == "success"),
				templ.KV("bg-red-500", props.Type == "error"),
				templ.KV("bg-yellow-500", props.Type == "warning"),
				templ.KV("bg-blue-500", props.Type == "info"),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var8...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-ref=\"progress\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var8).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toast.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div>")
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(props.Message)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toast.templ`, Line: 159, Col: 38}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Dismissible {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button @click.stop=\"dismissToast\">")
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

// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package chart

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/axzilla/templui/util"

type Variant string

const (
	VariantBar      Variant = "bar"
	VariantLine     Variant = "line"
	VariantPie      Variant = "pie"
	VariantDoughnut Variant = "doughnut"
	VariantRadar    Variant = "radar"
)

type Dataset struct {
	Label           string    `json:"label"`
	Data            []float64 `json:"data"`
	BorderWidth     int       `json:"borderWidth,omitempty"`
	BorderColor     string    `json:"borderColor,omitempty"`
	BackgroundColor string    `json:"backgroundColor,omitempty"`
	Tension         float64   `json:"tension,omitempty"`
	Fill            bool      `json:"fill,omitempty"`
	Stepped         bool      `json:"stepped,omitempty"`
}

type Options struct {
	Responsive bool `json:"responsive,omitempty"`
	Legend     bool `json:"legend,omitempty"`
}

type Data struct {
	Labels   []string  `json:"labels"`
	Datasets []Dataset `json:"datasets"`
}

type Config struct {
	Type        Variant `json:"type"`
	Data        Data    `json:"data"`
	Options     Options `json:"options,omitempty"`
	ShowLegend  bool    `json:"showLegend,omitempty"`
	ShowXAxis   bool    `json:"showXAxis"`
	ShowYAxis   bool    `json:"showYAxis"`
	ShowXLabels bool    `json:"showXLabels"`
	ShowYLabels bool    `json:"showYLabels"`
	ShowXGrid   bool    `json:"showXGrid"`
	ShowYGrid   bool    `json:"showYGrid"`
	Horizontal  bool    `json:"horizontal"`
	Stacked     bool    `json:"stacked"`
}

// Erweiterung des Props um ID und Attributes
type Props struct {
	ID          string
	Variant     Variant
	Data        Data
	Options     Options
	ShowLegend  bool
	ShowXAxis   bool
	ShowYAxis   bool
	ShowXLabels bool
	ShowYLabels bool
	ShowXGrid   bool
	ShowYGrid   bool
	Horizontal  bool
	Stacked     bool
	Class       string
	Attributes  templ.Attributes
}

func Chart(props ...Props) templ.Component {
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
		templ_7745c5c3_Err = Script().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var p Props
		if len(props) > 0 {
			p = props[0]
		}
		if p.ID == "" {
			p.ID = "chart-" + util.RandomID()
		}
		canvasId := p.ID + "-canvas"
		dataId := p.ID + "-data"
		var templ_7745c5c3_Var2 = []any{
			util.TwMerge(
				"chart-container relative",
				p.Class),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/chart/chart.templ`, Line: 83, Col: 11}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/chart/chart.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, p.Attributes)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "><canvas id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(canvasId)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/chart/chart.templ`, Line: 91, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" data-chart-id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(dataId)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/chart/chart.templ`, Line: 91, Col: 48}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\"></canvas></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}

		chartConfig := Config{
			Type:        p.Variant,
			Data:        p.Data,
			Options:     p.Options,
			ShowLegend:  p.ShowLegend,
			ShowXAxis:   p.ShowXAxis,
			ShowYAxis:   p.ShowYAxis,
			ShowXLabels: p.ShowXLabels,
			ShowYLabels: p.ShowYLabels,
			ShowXGrid:   p.ShowXGrid,
			ShowYGrid:   p.ShowYGrid,
			Horizontal:  p.Horizontal,
			Stacked:     p.Stacked,
		}
		templ_7745c5c3_Err = templ.JSONScript(dataId, chartConfig).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var handle = templ.NewOnceHandle()

func Script() templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<!-- Load Chart.js library --> <script defer nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/chart/chart.templ`, Line: 117, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\" src=\"https://cdn.jsdelivr.net/npm/chart.js@4.4.8/dist/chart.umd.min.js\"></script> <!-- Initialize charts --> <script nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/chart/chart.templ`, Line: 119, Col: 37}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\">\n\t\t\tdocument.addEventListener('DOMContentLoaded', function() {\n\t\t\t\tconst chartInstances = {};\n\t\t\t\tfunction getThemeColors() {\n\t\t\t\t\tconst computedStyle = getComputedStyle(document.documentElement);\n\t\t\t\t\treturn {\n\t\t\t\t\t\tforeground: computedStyle.getPropertyValue('--foreground').trim(),\n\t\t\t\t\t\tbackground: computedStyle.getPropertyValue('--background').trim(),\n\t\t\t\t\t\tmutedForeground: computedStyle.getPropertyValue('--muted-foreground').trim(),\n\t\t\t\t\t\tborder: computedStyle.getPropertyValue('--border').trim()\n\t\t\t\t\t};\n\t\t\t\t}\n\t\t\t\tfunction initCharts() {\n\t\t\t\t\tconst colors = getThemeColors();\n\t\t\t\t\tdocument.querySelectorAll('.chart-container').forEach(container => {\n\t\t\t\t\t\tconst canvas = container.querySelector('canvas');\n\t\t\t\t\t\tif (!canvas) return;\n\t\t\t\t\t\tconst dataId = canvas.getAttribute('data-chart-id');\n\t\t\t\t\t\tconst dataElement = document.getElementById(dataId);\n\t\t\t\t\t\tif (!dataElement) return;\n\t\t\t\t\t\tif (chartInstances[canvas.id]) {\n\t\t\t\t\t\t\tchartInstances[canvas.id].destroy();\n\t\t\t\t\t\t}\n\t\t\t\t\t\tconst chartConfig = JSON.parse(dataElement.textContent);\n\t\t\t\t\t\tChart.defaults.elements.point.radius = 0;\n\t\t\t\t\t\tChart.defaults.elements.point.hoverRadius = 5;\n\t\t\t\t\t\tconst chart = new Chart(canvas, {\n\t\t\t\t\t\t\t...chartConfig,\n\t\t\t\t\t\t\toptions: {\n\t\t\t\t\t\t\t\tinteraction: {\n\t\t\t\t\t\t\t\t\tintersect: false,\n\t\t\t\t\t\t\t\t\taxis: \"xy\",\n\t\t\t\t\t\t\t\t\tmode: chartConfig.type === \"pie\" || \n\t\t\t\t\t\t\t\t\t\tchartConfig.type === \"doughnut\" ||\n\t\t\t\t\t\t\t\t\t\tchartConfig.type === \"bar\" ||\n\t\t\t\t\t\t\t\t\t\tchartConfig.type === \"radar\" ? \"nearest\" : \"index\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\tindexAxis: chartConfig.horizontal ? \"y\" : \"x\",\n\t\t\t\t\t\t\t\tplugins: {\n\t\t\t\t\t\t\t\t\tlegend: {\n\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showLegend || false,\n\t\t\t\t\t\t\t\t\t\tlabels: {\n\t\t\t\t\t\t\t\t\t\t\tcolor: colors.foreground\n\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\ttooltip: {\n\t\t\t\t\t\t\t\t\t\tbackgroundColor: colors.background,\n\t\t\t\t\t\t\t\t\t\tbodyColor: colors.mutedForeground,\n\t\t\t\t\t\t\t\t\t\ttitleColor: colors.foreground,\n\t\t\t\t\t\t\t\t\t\tborderColor: colors.border,\n\t\t\t\t\t\t\t\t\t\tborderWidth: 1\n\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\tscales: chartConfig.type === \"radar\" ? {\n\t\t\t\t\t\t\t\t\tr: {\n\t\t\t\t\t\t\t\t\t\tgrid: {\n\t\t\t\t\t\t\t\t\t\t\tcolor: colors.border,\n\t\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showYGrid || true\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tticks: {\n\t\t\t\t\t\t\t\t\t\t\tcolor: colors.mutedForeground,\n\t\t\t\t\t\t\t\t\t\t\tbackdropColor: 'transparent',\n\t\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showYLabels || true\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tangleLines: {\n\t\t\t\t\t\t\t\t\t\t\tcolor: colors.border,\n\t\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showXGrid || true\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tpointLabels: {\n\t\t\t\t\t\t\t\t\t\t\tcolor: colors.foreground,\n\t\t\t\t\t\t\t\t\t\t\tfont: {\n\t\t\t\t\t\t\t\t\t\t\t\tsize: 12\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tborder: {\n\t\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showYAxis || false,\n\t\t\t\t\t\t\t\t\t\t\tcolor: colors.border\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tbeginAtZero: true\n\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t} : {\n\t\t\t\t\t\t\t\t\tx: {\n\t\t\t\t\t\t\t\t\t\tbeginAtZero: true,\n\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showXLabels || chartConfig.showXGrid || chartConfig.showXAxis,\n\t\t\t\t\t\t\t\t\t\tborder: {\n\t\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showXAxis || false\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tticks: {                                      \n\t\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showXLabels || false,   \n\t\t\t\t\t\t\t\t\t\t\tcolor: colors.mutedForeground\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tgrid: {\n\t\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showXGrid || false,\n\t\t\t\t\t\t\t\t\t\t\tcolor: colors.border\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tstacked: chartConfig.stacked\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\ty: {\n\t\t\t\t\t\t\t\t\t\toffset: true,\n\t\t\t\t\t\t\t\t\t\tbeginAtZero: true,\n\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showYLabels || chartConfig.showYGrid || chartConfig.showYAxis,\n\t\t\t\t\t\t\t\t\t\tborder: {\n\t\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showYAxis || false\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tticks: {                                      \n\t\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showYLabels || false,   \n\t\t\t\t\t\t\t\t\t\t\tcolor: colors.mutedForeground\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tgrid: {\n\t\t\t\t\t\t\t\t\t\t\tdisplay: chartConfig.showYGrid || false,\n\t\t\t\t\t\t\t\t\t\t\tcolor: colors.border\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\tstacked: chartConfig.stacked\n\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t});\n\t\t\t\t\t\tchartInstances[canvas.id] = chart;\n\t\t\t\t\t});\n\t\t\t\t}\n\t\t\t\tinitCharts();\n\t\t\t\twindow.addEventListener('resize', function() {\n\t\t\t\t\tfor (let id in chartInstances) {\n\t\t\t\t\t\tchartInstances[id].resize();\n\t\t\t\t\t}\n\t\t\t\t});\n\t\t\t\tconst observer = new MutationObserver(() => {\n\t\t\t\t\tsetTimeout(initCharts, 50);\n\t\t\t\t});\n\t\t\t\tobserver.observe(document.documentElement, { attributes: true });\n\t\t\t\tdocument.addEventListener('alpine:initialized', () => {\n\t\t\t\t\tdocument.addEventListener('theme-changed', initCharts);\n\t\t\t\t});\n\t\t\t});\n\t\t</script>")
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

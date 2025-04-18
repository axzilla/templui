// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package datepicker

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/axzilla/templui/component/button"
	"github.com/axzilla/templui/icon"
	"github.com/axzilla/templui/util"
	"strconv"
	"time"
)

type Format string

const (
	FormatISO  Format = "iso"
	FormatEU   Format = "eu"
	FormatUK   Format = "uk"
	FormatUS   Format = "us"
	FormatLONG Format = "long"
)

var formatMapping = map[Format]string{
	FormatISO:  "2006-01-02",
	FormatEU:   "02.01.2006",
	FormatUK:   "02/01/2006",
	FormatUS:   "01/02/2006",
	FormatLONG: "January 2, 2006",
}

type Locale struct {
	MonthNames []string
	DayNames   []string
}

var (
	LocaleDefault = Locale{
		MonthNames: []string{"January", "February", "March", "April", "May", "June",
			"July", "August", "September", "October", "November", "December"},
		DayNames: []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
	}

	LocaleSpanish = Locale{
		MonthNames: []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
			"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"},
		DayNames: []string{"Lu", "Ma", "Mi", "Ju", "Vi", "Sa", "Do"},
	}

	LocaleGerman = Locale{
		MonthNames: []string{"Januar", "Februar", "März", "April", "Mai", "Juni",
			"Juli", "August", "September", "Oktober", "November", "Dezember"},
		DayNames: []string{"Mo", "Di", "Mi", "Do", "Fr", "Sa", "So"},
	}

	LocaleFrench = Locale{
		MonthNames: []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin",
			"Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"},
		DayNames: []string{"Lu", "Ma", "Me", "Je", "Ve", "Sa", "Di"},
	}

	LocaleItalian = Locale{
		MonthNames: []string{"Gennaio", "Febbraio", "Marzo", "Aprile", "Maggio", "Giugno",
			"Luglio", "Agosto", "Settembre", "Ottobre", "Novembre", "Dicembre"},
		DayNames: []string{"Lu", "Ma", "Me", "Gi", "Ve", "Sa", "Do"},
	}

	LocaleJapanese = Locale{
		MonthNames: []string{"1月", "2月", "3月", "4月", "5月", "6月",
			"7月", "8月", "9月", "10月", "11月", "12月"},
		DayNames: []string{"日", "月", "火", "水", "木", "金", "土"},
	}
)

var (
	ConfigISO = Config{
		Format: FormatISO,
		Locale: LocaleDefault,
	}

	ConfigEU = Config{
		Format: FormatEU,
		Locale: LocaleDefault,
	}

	ConfigUK = Config{
		Format: FormatUK,
		Locale: LocaleDefault,
	}

	ConfigUS = Config{
		Format: FormatUS,
		Locale: LocaleDefault,
	}

	ConfigLONG = Config{
		Format: FormatLONG,
		Locale: LocaleDefault,
	}
)

func NewConfig(format Format, locale Locale) Config {
	return Config{
		Format: format,
		Locale: locale,
	}
}

type Config struct {
	Format Format
	Locale Locale
}

type Props struct {
	ID          string
	Class       string
	Attributes  templ.Attributes
	Value       time.Time
	Config      Config
	Placeholder string
	Disabled    bool
	Required    bool
	HasError    bool
	Name        string
}

func DatePicker(props ...Props) templ.Component {
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
		var p Props
		if len(props) > 0 {
			p = props[0]
		}
		if p.ID == "" {
			p.ID = util.RandomID()
		}
		if p.Placeholder == "" {
			p.Placeholder = "Select a date"
		}
		var templ_7745c5c3_Var2 = []any{util.TwMerge("relative", p.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID + "-container")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 136, Col: 26}
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
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.Value != (time.Time{}) {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, " data-value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(p.Value.Format(p.Config.getGoFormat()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 139, Col: 54}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, " data-format=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(string(p.Config.Format))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 141, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\" data-monthnames=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(p.Config.Locale.MonthNames))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 142, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\" data-daynames=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(p.Config.Locale.DayNames))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 143, Col: 60}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\" data-placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(p.Placeholder)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 144, Col: 34}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" x-data=\"datePicker\" @resize.window=\"updatePosition\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, p.Attributes)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "><input type=\"hidden\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID + "-hidden")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 151, Col: 24}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(p.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 152, Col: 16}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\" x-ref=\"hiddenInput\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.Disabled {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if p.Required {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, " required")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "><div class=\"relative\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var12 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "<span x-ref=\"displayText\" class=\"text-left grow\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if p.Value != (time.Time{}) {
				var templ_7745c5c3_Var13 string
				templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(p.Value.Format(p.Config.getGoFormat()))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 182, Col: 46}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "<span class=\"text-muted-foreground\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var14 string
				templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(p.Placeholder)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 184, Col: 57}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "</span><!-- Calendar icon --> <span class=\"text-muted-foreground flex items-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = icon.Calendar().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = button.Button(button.Props{
			ID:      p.ID,
			Type:    "button",
			Variant: button.VariantOutline,
			Class: util.TwMerge(
				"w-full select-trigger flex items-center justify-between focus:ring-2 focus:ring-offset-2",
				util.If(p.HasError, "border-destructive ring-destructive"),
				p.Class,
			),
			Disabled: p.Disabled,
			Attributes: util.MergeAttributes(
				templ.Attributes{
					"@click":   "toggleDatePicker",
					"x-ref":    "triggerButton",
					"required": strconv.FormatBool(p.Required),
				},
				p.Attributes,
			),
		}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var12), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "</div><!-- DatePicker Popup -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 = []any{
			util.TwMerge(
				"absolute left-0 z-50 w-64 p-4",
				"rounded-lg border bg-popover shadow-md",
			),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var15...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "<div x-show=\"open\" x-ref=\"datePickerPopup\" @click.away=\"closeDatePicker\" x-transition.opacity class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var15).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "\" style=\"display: none;\"><div class=\"flex items-center justify-between mb-4\"><span x-ref=\"monthDisplay\" class=\"text-sm font-medium\"></span><div class=\"flex gap-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 = []any{
			util.TwMerge(
				"inline-flex",
				"items-center",
				"justify-center",
				"rounded-md",
				"text-sm",
				"font-medium",
				"ring-offset-background",
				"transition-colors",
				"focus-visible:outline-hidden",
				"focus-visible:ring-2",
				"focus-visible:ring-ring",
				"focus-visible:ring-offset-2",
				"disabled:pointer-events-none",
				"disabled:opacity-50",
				"hover:bg-accent",
				"hover:text-accent-foreground",
				"h-7",
				"w-7",
			),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var17...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "<button type=\"button\" @click=\"prevMonth\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var18 string
		templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var17).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icon.ChevronLeft().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "</button> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var19 = []any{
			util.TwMerge(
				"inline-flex",
				"items-center",
				"justify-center",
				"rounded-md",
				"text-sm",
				"font-medium",
				"ring-offset-background",
				"transition-colors",
				"focus-visible:outline-hidden",
				"focus-visible:ring-2",
				"focus-visible:ring-ring",
				"focus-visible:ring-offset-2",
				"disabled:pointer-events-none",
				"disabled:opacity-50",
				"hover:bg-accent",
				"hover:text-accent-foreground",
				"h-7",
				"w-7",
			),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var19...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "<button type=\"button\" @click=\"nextMonth\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var20 string
		templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var19).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icon.ChevronRight().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "</button></div></div><!-- Weekday container --><div x-ref=\"weekdaysContainer\" class=\"grid grid-cols-7 gap-1 mb-2\"></div><!-- Calendar days container --><div x-ref=\"daysContainer\" class=\"grid grid-cols-7 gap-1\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func (c Config) getGoFormat() string {
	if format, ok := formatMapping[c.Format]; ok {
		return format
	}
	return formatMapping[FormatISO]
}

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
		templ_7745c5c3_Var21 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var21 == nil {
			templ_7745c5c3_Var21 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		handle := templ.NewOnceHandle()
		templ_7745c5c3_Var22 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "<script defer nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var23 string
			templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/datepicker/date_picker.templ`, Line: 286, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 32, "\">\n\t\t\tdocument.addEventListener('alpine:init', () => {\n\t\t\t\tAlpine.data('datePicker', () => ({\n\t\t\t\t\topen: false,\n\t\t\t\t\tvalue: null,\n\t\t\t\t\tformat: null,\n\t\t\t\t\tcurrentMonth: new Date().getMonth(),\n\t\t\t\t\tcurrentYear: new Date().getFullYear(),\n\t\t\t\t\tmonthNames: [],\n\t\t\t\t\tdayNames: [],\n\t\t\t\t\tposition: 'bottom',\n\n\t\t\t\t\tinit() {\n\t\t\t\t\t\t\t// Parse Month and Day names\n\t\t\t\t\t\ttry {\n\t\t\t\t\t\t\tthis.monthNames = JSON.parse(this.$el.dataset.monthnames) || [];\n\t\t\t\t\t\t\tthis.dayNames = JSON.parse(this.$el.dataset.daynames) || [];\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\tif (!this.monthNames.length) {\n\t\t\t\t\t\t\t\tthis.monthNames = [\n\t\t\t\t\t\t\t\t\t'January',\n\t\t\t\t\t\t\t\t\t'February', \n\t\t\t\t\t\t\t\t\t'March',\n\t\t\t\t\t\t\t\t\t'April',\n\t\t\t\t\t\t\t\t\t'May',\n\t\t\t\t\t\t\t\t\t'June',\n\t\t\t\t\t\t\t\t\t'July', \n\t\t\t\t\t\t\t\t\t'August',\n\t\t\t\t\t\t\t\t\t'September',\n\t\t\t\t\t\t\t\t\t'October',\n\t\t\t\t\t\t\t\t\t'November',\n\t\t\t\t\t\t\t\t\t'December'\n\t\t\t\t\t\t\t\t];\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\tif (!this.dayNames.length) {\n\t\t\t\t\t\t\t\tthis.dayNames = [\n\t\t\t\t\t\t\t\t\t'Su',\n\t\t\t\t\t\t\t\t\t'Mo',\n\t\t\t\t\t\t\t\t\t'Tu', \n\t\t\t\t\t\t\t\t\t'We',\n\t\t\t\t\t\t\t\t\t'Th',\n\t\t\t\t\t\t\t\t\t'Fr',\n\t\t\t\t\t\t\t\t\t'Sa'\n\t\t\t\t\t\t\t\t];\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t// Format and Placeholder\n\t\t\t\t\t\t\tthis.format = this.$el.dataset.format || 'iso';\n\t\t\t\t\t\t\tthis.placeholder = this.$el.dataset.placeholder || 'Select a date';\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t// Process initial value\n\t\t\t\t\t\t\tconst initialValue = this.$el.dataset.value;\n\t\t\t\t\t\t\tif (initialValue) {\n\t\t\t\t\t\t\t\tconst initialDate = new Date(this.parseDate(initialValue));\n\t\t\t\t\t\t\t\tthis.currentMonth = initialDate.getMonth();\n\t\t\t\t\t\t\t\tthis.currentYear = initialDate.getFullYear();\n\t\t\t\t\t\t\t\tthis.value = this.formatDate(initialDate);\n\t\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t\t// Update hidden input\n\t\t\t\t\t\t\t\tthis.$refs.hiddenInput.value = this.value;\n\t\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t\t// Update display text\n\t\t\t\t\t\t\t\tthis.updateDisplayText();\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t// Initialize UI\n\t\t\t\t\t\t\tthis.updateMonthDisplay();\n\t\t\t\t\t\t\tthis.renderWeekdays();\n\t\t\t\t\t\t} catch(e) {\n\t\t\t\t\t\t\tconsole.error('DatePicker initialization error:', e);\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tupdateDisplayText() {\n\t\t\t\t\t\tconst displayTextEl = this.$refs.displayText;\n\t\t\t\t\t\tif (!displayTextEl) return;\n\t\t\t\t\t\t\n\t\t\t\t\t\tdisplayTextEl.innerHTML = '';\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (this.value) {\n\t\t\t\t\t\t\tdisplayTextEl.textContent = this.value;\n\t\t\t\t\t\t\t// Remove muted style\n\t\t\t\t\t\t\tconst muted = displayTextEl.querySelector('.text-muted-foreground');\n\t\t\t\t\t\t\tif (muted) {\n\t\t\t\t\t\t\t\tmuted.classList.remove('text-muted-foreground');\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tconst span = document.createElement('span');\n\t\t\t\t\t\t\tspan.className = 'text-muted-foreground';\n\t\t\t\t\t\t\tspan.textContent = this.placeholder;\n\t\t\t\t\t\t\tdisplayTextEl.appendChild(span);\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\ttoggleDatePicker() {\n\t\t\t\t\t\tif (this.$refs.triggerButton.disabled) return;\n\t\t\t\t\t\t\n\t\t\t\t\t\tthis.open = !this.open;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (this.open) {\n\t\t\t\t\t\t\tthis.$nextTick(() => {\n\t\t\t\t\t\t\t\tthis.updatePosition();\n\t\t\t\t\t\t\t\tthis.renderCalendar();\n\t\t\t\t\t\t\t});\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t   closeDatePicker(event) {\n\t\t\t\t\t\tconst trigger = this.$refs.triggerButton;\n                        if (event.target.tagName === 'LABEL' && event.target.htmlFor === trigger.id) {\n                            return; \n\t\t\t\t\t\t}\n                        this.open = false;\n                    },\t\n\t\t\t\t\t\n\t\t\t\t\tupdateMonthDisplay() {\n\t\t\t\t\t\tif (this.$refs.monthDisplay) {\n\t\t\t\t\t\t\tthis.$refs.monthDisplay.textContent = this.monthNames[this.currentMonth] + ' ' + this.currentYear;\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tupdatePosition() {\n\t\t\t\t\t\tconst trigger = this.$refs.triggerButton;\n\t\t\t\t\t\tconst popup = this.$refs.datePickerPopup;\n\t\t\t\t\t\tif (!trigger || !popup) return;\n\t\t\t\t\t\t\n\t\t\t\t\t\tsetTimeout(() => {\n\t\t\t\t\t\t\tconst rect = trigger.getBoundingClientRect();\n\t\t\t\t\t\t\tconst popupRect = popup.getBoundingClientRect();\n\t\t\t\t\t\t\tconst viewportHeight = window.innerHeight;\n\t\t\t\t\t\t\tif (rect.bottom + popupRect.height > viewportHeight && rect.top > popupRect.height) {\n\t\t\t\t\t\t\t\tpopup.classList.add('bottom-full', 'mb-1');\n\t\t\t\t\t\t\t\tpopup.classList.remove('top-full', 'mt-1');\n\t\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\t\tpopup.classList.add('top-full', 'mt-1');\n\t\t\t\t\t\t\t\tpopup.classList.remove('bottom-full', 'mb-1');\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t}, 0); // 0ms is enough to wait for the next tick\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\trenderWeekdays() {\n\t\t\t\t\t\tconst container = this.$refs.weekdaysContainer;\n\t\t\t\t\t\tif (!container) return;\n\t\t\t\t\t\t\n\t\t\t\t\t\tcontainer.innerHTML = '';\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Day headers\n\t\t\t\t\t\tthis.dayNames.forEach(day => {\n\t\t\t\t\t\t\tconst el = document.createElement('div');\n\t\t\t\t\t\t\tel.className = 'text-center text-xs text-muted-foreground';\n\t\t\t\t\t\t\tel.textContent = day;\n\t\t\t\t\t\t\tcontainer.appendChild(el);\n\t\t\t\t\t\t});\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\trenderCalendar() {\n\t\t\t\t\t\t// Calculate first day and days in month\n\t\t\t\t\t\tlet firstDay = new Date(this.currentYear, this.currentMonth, 1).getDay();\n\t\t\t\t\t\tfirstDay = firstDay === 0 ? 6 : firstDay - 1; // Sunday = 6\n\t\t\t\t\t\t\n\t\t\t\t\t\tconst daysInMonth = new Date(this.currentYear, this.currentMonth + 1, 0).getDate();\n\t\t\t\t\t\tconst today = new Date();\n\t\t\t\t\t\tconst container = this.$refs.daysContainer;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (!container) return;\n\t\t\t\t\t\tcontainer.innerHTML = '';\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Empty days at the beginning\n\t\t\t\t\t\tfor (let i = 0; i < firstDay; i++) {\n\t\t\t\t\t\t\tconst blank = document.createElement('div');\n\t\t\t\t\t\t\tblank.className = 'h-8 w-8';\n\t\t\t\t\t\t\tcontainer.appendChild(blank);\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Days in the month\n\t\t\t\t\t\tfor (let day = 1; day <= daysInMonth; day++) {\n\t\t\t\t\t\t\tconst button = document.createElement('button');\n\t\t\t\t\t\t\tbutton.type = 'button';\n\t\t\t\t\t\t\tbutton.className = 'inline-flex h-8 w-8 items-center justify-center rounded-md text-sm font-medium';\n\t\t\t\t\t\t\tbutton.textContent = day;\n\t\t\t\t\t\t\tbutton.dataset.day = day;\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t// Click handler\n\t\t\t\t\t\t\tbutton.addEventListener('click', e => this.selectDate(e));\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t// Highlight today or selected\n\t\t\t\t\t\t\tconst isToday = today.getDate() === day && \n\t\t\t\t\t\t\t\t\t\t   today.getMonth() === this.currentMonth && \n\t\t\t\t\t\t\t\t\t\t   today.getFullYear() === this.currentYear;\n\t\t\t\t\t\t\t\t\t\t   \n\t\t\t\t\t\t\tconst isSelected = this.isSelectedDate(day);\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\tif (isSelected) {\n\t\t\t\t\t\t\t\tbutton.classList.add('bg-primary', 'text-primary-foreground');\n\t\t\t\t\t\t\t} else if (isToday) {\n\t\t\t\t\t\t\t\tbutton.classList.add('text-red-500');\n\t\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\t\tbutton.classList.add('hover:bg-accent', 'hover:text-accent-foreground');\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\tcontainer.appendChild(button);\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tprevMonth() {\n\t\t\t\t\t\tthis.currentMonth--;\n\t\t\t\t\t\tif (this.currentMonth < 0) {\n\t\t\t\t\t\t\tthis.currentMonth = 11;\n\t\t\t\t\t\t\tthis.currentYear--;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tthis.updateMonthDisplay();\n\t\t\t\t\t\tthis.renderCalendar();\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tnextMonth() {\n\t\t\t\t\t\tthis.currentMonth++;\n\t\t\t\t\t\tif (this.currentMonth > 11) {\n\t\t\t\t\t\t\tthis.currentMonth = 0;\n\t\t\t\t\t\t\tthis.currentYear++;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tthis.updateMonthDisplay();\n\t\t\t\t\t\tthis.renderCalendar();\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tparseDate(dateStr) {\n\t\t\t\t\t\tconst parts = dateStr.split(/[-/.]/);\n\t\t\t\t\t\tswitch(this.format) {\n\t\t\t\t\t\t\tcase 'eu':\n\t\t\t\t\t\t\t\treturn `${parts[2]}-${parts[1]}-${parts[0]}`;\n\t\t\t\t\t\t\tcase 'us':\n\t\t\t\t\t\t\t\treturn `${parts[2]}-${parts[0]}-${parts[1]}`;\n\t\t\t\t\t\t\tcase 'uk':\n\t\t\t\t\t\t\t\treturn `${parts[2]}-${parts[1]}-${parts[0]}`;\n\t\t\t\t\t\t\tcase 'long':\n\t\t\t\t\t\t\tcase 'iso':\n\t\t\t\t\t\t\tdefault:\n\t\t\t\t\t\t\t\treturn dateStr;\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tformatDate(date) {\n\t\t\t\t\t\tconst d = date.getDate().toString().padStart(2, '0');\n\t\t\t\t\t\tconst m = (date.getMonth() + 1).toString().padStart(2, '0');\n\t\t\t\t\t\tconst y = date.getFullYear();\n\t\t\t\t\t\t\n\t\t\t\t\t\tswitch(this.format) {\n\t\t\t\t\t\t\tcase 'eu':\n\t\t\t\t\t\t\t\treturn `${d}.${m}.${y}`;\n\t\t\t\t\t\t\tcase 'uk':\n\t\t\t\t\t\t\t\treturn `${d}/${m}/${y}`;\n\t\t\t\t\t\t\tcase 'us':\n\t\t\t\t\t\t\t\treturn `${m}/${d}/${y}`;\n\t\t\t\t\t\t\tcase 'long':\n\t\t\t\t\t\t\t\treturn `${this.monthNames[date.getMonth()]} ${d}, ${y}`;\n\t\t\t\t\t\t\tcase 'iso':\n\t\t\t\t\t\t\tdefault:\n\t\t\t\t\t\t\t\treturn `${y}-${m}-${d}`;\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tisSelectedDate(day) {\n\t\t\t\t\t\tif (!this.value) return false;\n\t\t\t\t\t\t\n\t\t\t\t\t\ttry {\n\t\t\t\t\t\t\tconst date = new Date(this.currentYear, this.currentMonth, parseInt(day));\n\t\t\t\t\t\t\tconst selected = new Date(this.parseDate(this.value));\n\t\t\t\t\t\t\treturn date.toDateString() === selected.toDateString();\n\t\t\t\t\t\t} catch(e) {\n\t\t\t\t\t\t\treturn false;\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tselectDate(e) {\n\t\t\t\t\t\tconst day = e.target.dataset.day;\n\t\t\t\t\t\tif (!day) return;\n\t\t\t\t\t\t\n\t\t\t\t\t\tconst date = new Date(this.currentYear, this.currentMonth, parseInt(day));\n\t\t\t\t\t\tthis.value = this.formatDate(date);\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Update hidden input\n\t\t\t\t\t\tthis.$refs.hiddenInput.value = this.value;\n\t\t\t\t\t\tthis.$refs.hiddenInput.dispatchEvent(new Event('change', { bubbles: true }));\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Update display text\n\t\t\t\t\t\tthis.updateDisplayText();\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Close DatePicker\n\t\t\t\t\t\tthis.open = false;\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Focus on button\n\t\t\t\t\t\tthis.$refs.triggerButton.focus();\n\t\t\t\t\t}\n\t\t\t\t}));\n\t\t\t});\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = handle.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var22), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/axzilla/templui/internal/ui/showcase"
)

func writeHTML(filename string, c templ.Component) error {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// ✅ Use context.Background() instead of nil
	return c.Render(context.Background(), f)
}

func main() {
	showcases := map[string]templ.Component{
		"out/showcase/accordion_default.html": showcase.AccordionDefault(),

		// Alert
		"out/showcase/alert_default.html":     showcase.AlertDefault(),
		"out/showcase/alert_destructive.html": showcase.AlertDestructive(),

		// Aspect Ratio
		"out/showcase/aspect_ratio_default.html": showcase.AspectRatioDefault(),

		// Avatar
		"out/showcase/avatar_default.html":   showcase.AvatarDefault(),
		"out/showcase/avatar_fallback.html":  showcase.AvatarFallback(),
		"out/showcase/avatar_group.html":     showcase.AvatarGroup(),
		"out/showcase/avatar_sizes.html":     showcase.AvatarSizes(),
		"out/showcase/avatar_with_icon.html": showcase.AvatarWithIcon(),

		// Badge
		"out/showcase/badge_default.html":     showcase.BadgeDefault(),
		"out/showcase/badge_destructive.html": showcase.BadgeDestructive(),
		"out/showcase/badge_outline.html":     showcase.BadgeOutline(),
		"out/showcase/badge_secondary.html":   showcase.BadgeSecondary(),
		"out/showcase/badge_with_icon.html":   showcase.BadgeWithIcon(),

		// Breadcrumb
		"out/showcase/breadcrumb_custom_separator.html": showcase.BreadcrumbCustomSeparator(),
		"out/showcase/breadcrumb_default.html":          showcase.BreadcrumbDefault(),
		"out/showcase/breadcrumb_responsive.html":       showcase.BreadcrumbResponsive(),
		"out/showcase/breadcrumb_with_icons.html":       showcase.BreadcrumbWithIcons(),

		// Button
		"out/showcase/button_default.html":      showcase.ButtonDefault(),
		"out/showcase/button_destructive.html":  showcase.ButtonDestructive(),
		"out/showcase/button_ghost.html":        showcase.ButtonGhost(),
		"out/showcase/button_htmx_loading.html": showcase.ButtonHtmxLoading(),
		"out/showcase/button_icon.html":         showcase.ButtonIcon(),
		"out/showcase/button_link.html":         showcase.ButtonLink(),
		"out/showcase/button_loading.html":      showcase.ButtonLoading(),
		"out/showcase/button_outline.html":      showcase.ButtonOutline(),
		"out/showcase/button_primary.html":      showcase.ButtonPrimary(),
		"out/showcase/button_secondary.html":    showcase.ButtonSecondary(),
		"out/showcase/button_with_icon.html":    showcase.ButtonWithIcon(),

		// Card
		"out/showcase/card_default.html":      showcase.CardDefault(),
		"out/showcase/card_image_bottom.html": showcase.CardImageBottom(),
		"out/showcase/card_image_left.html":   showcase.CardImageLeft(),
		"out/showcase/card_image_right.html":  showcase.CardImageRight(),
		"out/showcase/card_image_top.html":    showcase.CardImageTop(),

		// Carousel
		"out/showcase/carousel_autoplay.html":    showcase.CarouselAutoplay(),
		"out/showcase/carousel_default.html":     showcase.CarouselDefault(),
		"out/showcase/carousel_minimal.html":     showcase.CarouselMinimal(),
		"out/showcase/carousel_with_images.html": showcase.CarouselWithImages(),

		// Chart
		"out/showcase/chart_area.html":             showcase.ChartArea(),
		"out/showcase/chart_area_linear.html":      showcase.ChartAreaLinear(),
		"out/showcase/chart_area_stacked.html":     showcase.ChartAreaStacked(),
		"out/showcase/chart_area_step.html":        showcase.ChartAreaStep(),
		"out/showcase/chart_bar_horizontal.html":   showcase.ChartBarHorizontal(),
		"out/showcase/chart_bar_multiple.html":     showcase.ChartBarMultiple(),
		"out/showcase/chart_bar_negative.html":     showcase.ChartBarNegative(),
		"out/showcase/chart_bar_stacked.html":      showcase.ChartBarStacked(),
		"out/showcase/chart_default.html":          showcase.ChartDefault(),
		"out/showcase/chart_doughnut.html":         showcase.ChartDoughnut(),
		"out/showcase/chart_doughnut_legend.html":  showcase.ChartDoughnutLegend(),
		"out/showcase/chart_doughnut_stacked.html": showcase.ChartDoughnutStacked(),
		"out/showcase/chart_line.html":             showcase.ChartLine(),
		"out/showcase/chart_line_linear.html":      showcase.ChartLineLinear(),
		"out/showcase/chart_line_multiple.html":    showcase.ChartLineMultiple(),
		"out/showcase/chart_line_step.html":        showcase.ChartLineStep(),
		"out/showcase/chart_pie.html":              showcase.ChartPie(),
		"out/showcase/chart_pie_legend.html":       showcase.ChartPieLegend(),
		"out/showcase/chart_pie_stacked.html":      showcase.ChartPieStacked(),
		"out/showcase/chart_radar.html":            showcase.ChartRadar(),
	}

	for path, comp := range showcases {
		fmt.Println("Rendering:", path)
		if err := writeHTML(path, comp); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error rendering %s: %v\n", path, err)
		}
	}
}

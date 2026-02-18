package tui

import (
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/utils"
)

// Breakpoint constants for responsive design.
const (
	// MinWidth is the minimum supported terminal width.
	MinWidth = 40
	// MinHeight is the minimum supported terminal height.
	MinHeight = 15

	// CompactWidth is the threshold for compact mode.
	CompactWidth = 80
	// StandardWidth is the threshold for standard mode.
	StandardWidth = 120

	// SidebarWidth is the typical sidebar width.
	SidebarWidth = 30
	// ContentWidth is the typical content area width.
	ContentWidth = 80
)

// ResponsiveConfig holds configuration for responsive behavior.
type ResponsiveConfig struct {
	// Compact adaptations
	CompactPadding     int
	CompactMargin      int
	CompactLabelWidth  int
	CompactShowHelp    bool

	// Standard adaptations
	StandardPadding    int
	StandardMargin     int
	StandardLabelWidth int
	StandardShowHelp   bool

	// Wide adaptations
	WidePadding     int
	WideMargin      int
	WideLabelWidth  int
	WideShowHelp    bool
	WideExtraInfo   bool
}

// DefaultResponsiveConfig returns the default responsive configuration.
func DefaultResponsiveConfig() ResponsiveConfig {
	return ResponsiveConfig{
		CompactPadding:     1,
		CompactMargin:      0,
		CompactLabelWidth:  12,
		CompactShowHelp:    false,

		StandardPadding:    2,
		StandardMargin:     1,
		StandardLabelWidth: 20,
		StandardShowHelp:   true,

		WidePadding:     3,
		WideMargin:      2,
		WideLabelWidth:  25,
		WideShowHelp:    true,
		WideExtraInfo:   true,
	}
}

// Responsive provides responsive layout functionality.
type Responsive struct {
	config     ResponsiveConfig
	width      int
	height     int
	breakpoint styles.Breakpoint
	theme      *styles.Theme

	// Cached values
	padding    int
	margin     int
	labelWidth int
	showHelp   bool
	extraInfo  bool
}

// NewResponsive creates a new Responsive instance.
func NewResponsive(config ResponsiveConfig) *Responsive {
	return &Responsive{
		config: config,
		theme:  styles.GetTheme(),
	}
}

// Update updates the responsive layout with new dimensions.
func (r *Responsive) Update(width, height int) {
	r.width = width
	r.height = height
	r.breakpoint = styles.CalculateBreakpoint(width)

	// Update cached values based on breakpoint
	switch r.breakpoint {
	case styles.BreakpointCompact:
		r.padding = r.config.CompactPadding
		r.margin = r.config.CompactMargin
		r.labelWidth = r.config.CompactLabelWidth
		r.showHelp = r.config.CompactShowHelp
		r.extraInfo = false
	case styles.BreakpointWide:
		r.padding = r.config.WidePadding
		r.margin = r.config.WideMargin
		r.labelWidth = r.config.WideLabelWidth
		r.showHelp = r.config.WideShowHelp
		r.extraInfo = r.config.WideExtraInfo
	default: // Standard
		r.padding = r.config.StandardPadding
		r.margin = r.config.StandardMargin
		r.labelWidth = r.config.StandardLabelWidth
		r.showHelp = r.config.StandardShowHelp
		r.extraInfo = false
	}
}

// Width returns the current width.
func (r *Responsive) Width() int {
	return r.width
}

// Height returns the current height.
func (r *Responsive) Height() int {
	return r.height
}

// Breakpoint returns the current breakpoint.
func (r *Responsive) Breakpoint() styles.Breakpoint {
	return r.breakpoint
}

// IsCompact returns true in compact mode.
func (r *Responsive) IsCompact() bool {
	return r.breakpoint == styles.BreakpointCompact
}

// IsStandard returns true in standard mode.
func (r *Responsive) IsStandard() bool {
	return r.breakpoint == styles.BreakpointStandard
}

// IsWide returns true in wide mode.
func (r *Responsive) IsWide() bool {
	return r.breakpoint == styles.BreakpointWide
}

// Padding returns the appropriate padding.
func (r *Responsive) Padding() int {
	return r.padding
}

// Margin returns the appropriate margin.
func (r *Responsive) Margin() int {
	return r.margin
}

// LabelWidth returns the appropriate label width.
func (r *Responsive) LabelWidth() int {
	return r.labelWidth
}

// ShowHelp returns whether to show help.
func (r *Responsive) ShowHelp() bool {
	return r.showHelp
}

// ExtraInfo returns whether to show extra information.
func (r *Responsive) ExtraInfo() bool {
	return r.extraInfo
}

// ContentWidth returns the available content width.
func (r *Responsive) ContentWidth() int {
	return r.width - (r.padding * 2) - (r.margin * 2)
}

// AvailableHeight returns the available height for content.
func (r *Responsive) AvailableHeight(reserveForHeader, reserveForFooter int) int {
	return r.height - reserveForHeader - reserveForFooter - (r.padding * 2)
}

// GridColumns returns the optimal number of columns for a grid.
func (r *Responsive) GridColumns(itemWidth int) int {
	available := r.ContentWidth()
	if available <= 0 || itemWidth <= 0 {
		return 1
	}

	cols := available / itemWidth
	if cols < 1 {
		return 1
	}
	return cols
}

// CanShowSidebar returns true if there's room for a sidebar.
func (r *Responsive) CanShowSidebar(sidebarWidth int) bool {
	return r.width >= SidebarWidth+sidebarWidth+ContentWidth
}

// SplitForSidebar returns widths for main content and sidebar.
func (r *Responsive) SplitForSidebar(sidebarWidth int) (mainWidth, sideWidth int) {
	if !r.CanShowSidebar(sidebarWidth) {
		return r.ContentWidth(), 0
	}

	available := r.ContentWidth()
	sideWidth = sidebarWidth
	mainWidth = available - sideWidth - r.margin
	return
}

// Truncate truncates text to fit within available width.
func (r *Responsive) Truncate(text string, maxLen int) string {
	available := r.ContentWidth()
	if maxLen > 0 && maxLen < available {
		available = maxLen
	}
	return utils.TruncateText(text, available)
}

// Wrap wraps text to fit within available width.
func (r *Responsive) Wrap(text string) string {
	return styles.TextWrap(text, r.ContentWidth())
}

// ResponsiveModel is a base for models that need responsive behavior.
type ResponsiveModel struct {
	responsive *Responsive
	mu         sync.RWMutex
}

// NewResponsiveModel creates a new responsive model.
func NewResponsiveModel(config ResponsiveConfig) *ResponsiveModel {
	return &ResponsiveModel{
		responsive: NewResponsive(config),
	}
}

// UpdateSize updates the responsive dimensions.
func (m *ResponsiveModel) UpdateSize(width, height int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.responsive.Update(width, height)
}

// GetResponsive returns the responsive configuration.
func (m *ResponsiveModel) GetResponsive() *Responsive {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.responsive
}

// HandleResize handles window resize messages.
func (m *ResponsiveModel) HandleResize(msg tea.WindowSizeMsg) {
	m.UpdateSize(msg.Width, msg.Height)
}

// ScreenSizer calculates sizes for screen components.
type ScreenSizer struct {
	width       int
	height      int
	headerLines int
	footerLines int
}

// NewScreenSizer creates a new screen sizer.
func NewScreenSizer(width, height int) *ScreenSizer {
	return &ScreenSizer{
		width:  width,
		height: height,
	}
}

// SetHeader sets the number of header lines.
func (s *ScreenSizer) SetHeader(lines int) *ScreenSizer {
	s.headerLines = lines
	return s
}

// SetFooter sets the number of footer lines.
func (s *ScreenSizer) SetFooter(lines int) *ScreenSizer {
	s.footerLines = lines
	return s
}

// ContentHeight returns available height for content.
func (s *ScreenSizer) ContentHeight() int {
	return s.height - s.headerLines - s.footerLines
}

// ContentWidth returns available width for content.
func (s *ScreenSizer) ContentWidth() int {
	return s.width
}

// MaxVisibleItems returns the maximum number of visible items.
func (s *ScreenSizer) MaxVisibleItems(itemHeight int) int {
	if itemHeight <= 0 {
		itemHeight = 1
	}
	return s.ContentHeight() / itemHeight
}

// CalculateVisibleRange calculates the visible range for a scrollable list.
func CalculateVisibleRange(total, selected, visible, offset int) (start, end int) {
	if total <= visible {
		return 0, total
	}

	// Ensure selected is in bounds
	if selected < 0 {
		selected = 0
	}
	if selected >= total {
		selected = total - 1
	}

	// If offset is valid, use it
	if offset >= 0 && offset+visible <= total {
		start = offset
	} else {
		// Center selected item
		start = selected - visible/2
		if start < 0 {
			start = 0
		}
		if start+visible > total {
			start = total - visible
		}
	}

	end = start + visible
	if end > total {
		end = total
	}

	return start, end
}

// IsVisible checks if an index is within the visible range.
func IsVisible(index, start, end int) bool {
	return index >= start && index < end
}

// EnsureVisible ensures an item is visible, returning the new offset.
func EnsureVisible(index, currentOffset, visibleCount, totalCount int) int {
	if totalCount <= visibleCount {
		return 0
	}

	// If item is before visible range
	if index < currentOffset {
		return index
	}

	// If item is after visible range
	if index >= currentOffset+visibleCount {
		return index - visibleCount + 1
	}

	// Item is already visible
	return currentOffset
}

// ColumnLayout handles multi-column responsive layouts.
type ColumnLayout struct {
	columns    int
	gap        int
	width      int
	breakpoint styles.Breakpoint
}

// NewColumnLayout creates a new column layout.
func NewColumnLayout(width int) *ColumnLayout {
	return &ColumnLayout{
		width:      width,
		breakpoint: styles.CalculateBreakpoint(width),
	}
}

// SetGap sets the gap between columns.
func (c *ColumnLayout) SetGap(gap int) *ColumnLayout {
	c.gap = gap
	return c
}

// Calculate determines the number of columns and their widths.
func (c *ColumnLayout) Calculate(minColumnWidth int) (columns int, columnWidths []int) {
	available := c.width
	gapWidth := 0

	// Determine number of columns based on breakpoint and available width
	switch c.breakpoint {
	case styles.BreakpointCompact:
		columns = 1
	case styles.BreakpointWide:
		columns = (available + c.gap) / (minColumnWidth + c.gap)
		if columns < 2 {
			columns = 2
		}
		if columns > 4 {
			columns = 4
		}
	default:
		columns = (available + c.gap) / (minColumnWidth + c.gap)
		if columns < 1 {
			columns = 1
		}
		if columns > 3 {
			columns = 3
		}
	}

	// Calculate gap width
	gapWidth = c.gap * (columns - 1)
	if gapWidth < 0 {
		gapWidth = 0
	}

	// Distribute width among columns
	usableWidth := available - gapWidth
	baseWidth := usableWidth / columns
	remainder := usableWidth % columns

	columnWidths = make([]int, columns)
	for i := 0; i < columns; i++ {
		columnWidths[i] = baseWidth
		if i < remainder {
			columnWidths[i]++
		}
	}

	return columns, columnWidths
}

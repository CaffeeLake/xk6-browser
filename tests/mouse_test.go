package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMouseActions(t *testing.T) {
	t.Parallel()

	t.Run("click", func(t *testing.T) {
		t.Parallel()

		tb := newTestBrowser(t)
		p := tb.NewPage(nil)
		m := p.GetMouse()

		// Set up a page with a button that changes text when clicked
		buttonHTML := `
			<button onclick="this.innerHTML='Clicked!'">Click me</button>
		`
		p.SetContent(buttonHTML, nil)
		button, err := p.Query("button")
		require.NoError(t, err)

		// Simulate a click at the button coordinates
		box := button.BoundingBox()
		m.Click(box.X, box.Y, nil)

		// Verify the button's text changed
		assert.Equal(t, "Clicked!", button.TextContent())
	})

	t.Run("double_click", func(t *testing.T) {
		t.Parallel()

		tb := newTestBrowser(t)
		p := tb.NewPage(nil)
		m := p.GetMouse()

		// Set up a page with a button that changes text on double click and also counts clicks
		buttonHTML := `
			<script>window.clickCount = 0;</script>
			<button
				onclick="document.getElementById('clicks').innerHTML = ++window.clickCount;"
				ondblclick="this.innerHTML='Double Clicked!';">Click me</button>
			<div id="clicks"></div>
		`
		p.SetContent(buttonHTML, nil)
		button, err := p.Query("button")
		require.NoError(t, err)

		// Get the button's bounding box for accurate clicking
		box := button.BoundingBox()

		// Simulate a double click at the button coordinates
		m.DblClick(box.X, box.Y, nil)

		// Verify the button's text changed
		assert.Equal(t, "Double Clicked!", button.TextContent())

		// Also verify that the element was clicked twice
		clickCountDiv, err := p.Query("div#clicks")
		require.NoError(t, err)
		assert.Equal(t, "2", clickCountDiv.TextContent())
	})

	t.Run("move", func(t *testing.T) {
		t.Parallel()

		tb := newTestBrowser(t)
		p := tb.NewPage(nil)
		m := p.GetMouse()

		// Set up a page with an area that detects mouse move
		areaHTML := `
			<div
				onmousemove="this.innerHTML='Mouse Moved';"
				style="width:100px;height:100px;"
			></div>
		`
		p.SetContent(areaHTML, nil)
		area, err := p.Query("div")
		require.NoError(t, err)

		// Simulate mouse move within the div
		box := area.BoundingBox()
		m.Move(box.X+50, box.Y+50, nil) // Move to the center of the div
		assert.Equal(t, "Mouse Moved", area.TextContent())
	})

	t.Run("move_down_up", func(t *testing.T) {
		t.Parallel()

		tb := newTestBrowser(t)
		p := tb.NewPage(nil)
		m := p.GetMouse()

		// Set up a page with a button that tracks mouse down and up
		buttonHTML := `
			<button
				onmousedown="this.innerHTML='Mouse Down';"
				onmouseup="this.innerHTML='Mouse Up';"
			>Mouse</button>
		`
		p.SetContent(buttonHTML, nil)
		button, err := p.Query("button")
		require.NoError(t, err)

		box := button.BoundingBox()
		m.Move(box.X, box.Y, nil)
		m.Down(nil)
		assert.Equal(t, "Mouse Down", button.TextContent())
		m.Up(nil)
		assert.Equal(t, "Mouse Up", button.TextContent())
	})
}

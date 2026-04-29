import "github.com/XenomorphingTV/waymark/parser"

type Engine struct {
	Story     *parser.Story
	Vars      map[string]any
	Callbacks []callback
	Pos       cursor
}

type callback struct {
	scene     *parser.SceneNode
	pos       int
	localvars map[string]any
}

type cursor struct {
	scene *parser.SceneNode
	pos   int
}

type State struct {
	Lines   []Line
	Choices []Choice
}

type Line struct {
	Content    string
	IsDialogue bool
}

type Choice struct {
	Index int
	Label string
}

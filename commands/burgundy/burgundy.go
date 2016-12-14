package burgundy

import (
	"fmt"
	"io"
)

const (
	ron = `
                              ,,▄▄▓▄▄▄▄,,,,
                       ,,▄╣▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▄▄,
                  ╓▄▄▓▓▓▓▓▓▓▓▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▄▄
               ▄▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▄╕
             ▄▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▄
           ╗▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓╕
          ╣▓▓▓▓▓▓▓▓▓▒Θ░░╣▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▓▓▓▓▓▓▓▓▓▓▄
         ╣▓▓▓▓▓▓▓█Θ▒╣╣▒▒▓▓▓▓▓▓▓▓▓██████ΘΘΘΘ░δΘ▒▓▓▒╣▓▓▓▓▓▓▓▓▓▄
        ╢▓▓▓▓▓▓▓▒╣╣╣▒▓▓▓▓█▀▀╨░ . ░ ░░░░░░░░░░╢╢╣╣▒▒▒▓▓▓▓▓▓▓▓▓▄
       ╥▓▓▓▓▓▓▓▓▓▒ΘΘΘΘ░░░░           ░░░░░░░░░░╢╣░╢▒▓▓▓▓▓▓▓▓▓▓Q
       ╫▓▓▓▓▒╣▓▓█░░░░░░               ░░░░░░░░░░░░░╢▓▓▓▓▓▓▓▓▓▓▓╕
       ╫▓▓▓▒╣▓▓█ ░░                  ░░░░░░░░░░░░░░╬▓▓▓▓▓▓▓▓▓▓▓▓
       ╫▒▒▒▒▒▓▓∩                   ░░░░░░░░░░░░░░░░░╫▓▓▓▓▓▓▓▓▓▓▓▌
       ║▒▒▒╬▓▓▓                   ░░░░░░░░░░░░░░╣╢╣╣╢▓▓▓▓▓▓▓▓▓▓▓▓▄
       ╠▒▓▓▓▓▓▓░                    ░░░░░░░░░░░╣╣╣╣╣╣▒▓▓▓▓▓▓▓▓▓▓▓▓
       ╠▒▓▓▓▓▓▓▌                      ⌐░░░╣╣▓▓▓▓▓▓▓▒▒╣▓▓▓▓▓▓▓▓▓▓▓▓▌
       ╙▒▒▓▓▓▓▓▓ |            ░░ ⌐░░░░╣╣▒▓▓▓▓▓▓▓▓▓▓▓▒▒▒▓▓▓▓▓▓▓▓▓▓▓Θ
        ░▒▒▒▓▓▓▓░  ⌐╔░╣╣╣╣╣╣╣╣░░░░░░╣▓▓▓▓██▓▒░▒▓▒▒Θ▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓
         ▓▓╣▒▓▓▓⌐ ░▒▒╣▒▓▓▓▓▓█▒╣░  ░░▒▓▒▓░╖░ΘΘ╣▒Θ░╢╣░╢▒▒▒▓▓▓▓▓▓▓▓▓⌐
         ╫▓▓▓▓▓▓░ ░░░ÑΘ░░ΘΘ░░░░░⌐ ░░╫▒░░░░░░░░░░░░░╢╢▒▒▒▓▓▓▓▓▓▓▓▓
         ╚▓▓▓▓▓▓░⌐     ░░░░░   ░  ░╢╣╣░░░░░░░░░░░░░░╢▒▒╣▓▓▓▓▓▓▓▓▓
          '█▓▓▓▓░⌐             ░  ░╢╣╣╣░░░░░░░░░░░░╢╣▒▒▒▓▓▓▓▓▓▓▓▓╕
             '▀▓▓░⌐           ░░  ░░╢╣▒╣░░░░░░░░░░╢╣▒▒▒▓▓▓▓▓▓▓▓▓▓░
                ▒░░⌐         ░░   "░╢╣╣Θ▒╣░░░░░░░╢╣▒▒▒▒▓▓▓▓▓▓▓▓▓▀
                 ░░░⌐    -░░░░╓⌐  -░╣╣╣╣▒▒░░░░░░╢╣▒▒▒▒▓▓▓▓▓▓▓▓█'
                  ░░░⌐ ⌐░░░░  ░░░╢╣▒▓▓▒δ░░░╣░░░░╢╣▒▒▒▒▓▓▓▓▓▓▓▀
                   ░░░░░╢Θ ╓░╣▓▒╣▓▓▒▒▓▓▓▓▓▓▓╣╣╣░░╢╣╣▒▓▓▓▓▓▓█
                   ░░░░░╢╥╣▒▓▓▒▓▓▓▒▒▒█▒▓▓▓▓▓▓▓▓▓▓╣╣▓▓▓▓▓▓▓█
                    ░░░░▓▓▓▓▓Θ╜''    ░░░╣╣▓▓▓▓▓▓▒▒▒▓▓▓▓▓Γ
                     ░░░░Θδ░░░░╣╣▒▒▒▒▒▓▓▓▓▓▒╣╣╣╣▒▒▓▓▓▓▓▓
                      "╢░░░░░░░░░░░░░░░░Θ▒░░░╢╣▒▓▓▓▓▓▓▓░
                       ╠╫░░░░░░░░░░░░░░░╣╢╣╣░╢╣▓▓▓▓▓▓▓▓
                      -░░░╢░░░░░░░░░░░╣╣╣╣╣╣╬╣▓▓▓▓▓▓▓▓▓▌
                       ╙░░░╢░╣░░░░░░╣╢╣╣╣▒▒╣▓▓▓▓▓▓▓▓▓▒╣▒▓
                   ╥╣    ╙░░░░╢╣╣╣╣╣▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓Θ░╣╣╣▓▓▄
                 ╓╣▓⌐       ░░░░╢▒▒▒▒▓▓▓▓▓▒▒▒▒▒▒Θ░░░░╢╣╣╣▓▓▓▄
               ╓▒▓▓▒          "ò░░░╢╢Θ▒ΘΘΘδ╢╣▒δ░░░░░░░╢╣╣▓▓▓▓▓▄╦
            ╖╣▓▓▒▒▓Θ             ╚░░░░░░░░░δ╨ ░░░░░░░░╢╣╣▓▓▓▓▓▓▓▓▓▓▄▄▄-µ
       ,╥▄▓▓▓▓▒▒▒▒▓░               "░░░░░╨       ░░░░░╢╢▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▄▄
  .╗▄▓▓▓▓▓▓▓▓▓▓▒▒▓▓░                  ╙'          ░░░░╢╣▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓

`
)

// Command ...
type Command struct {
	Name string
	W    io.Writer
	WErr io.Writer
}

// Key returns the commands name for sorting.
func (c *Command) Key() string {
	return c.Name
}

// Run ...
func (c *Command) Run(args []string) (int, error) {
	fmt.Fprintf(c.W, ron)
	return 0, nil
}

// Aliases are the aliases and name for the command. For instance
// a command can have a long form and short form.
func (c *Command) Aliases() map[string]struct{} {
	return map[string]struct{}{
		"burgundy": struct{}{},
	}
}

// Description is what is printed in Usage.
func (c *Command) Description() string {
	return "Stay classy."
}

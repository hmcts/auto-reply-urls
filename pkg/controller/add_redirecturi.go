package controller

import (
	"github.com/hmcts/auto-reply-urls/pkg/controller/redirecturi"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, redirecturi.Add)
}

package tools

import (
	"context"
	"fmt"
	"os"

	customsearch "google.golang.org/api/customsearch/v1"
	option "google.golang.org/api/option"
)

// GoogleSearch returns the results of a Google search for the given query
package templates

import "embed"

// Tell the embed package that we want to embed files at compile-time
// The * is what is known as a glob pattern. It helps us define which types of files
// to match. We could have used something like *.gohtml which would have
// only included files with the .gohtml extension

//go:embed *
var FS embed.FS

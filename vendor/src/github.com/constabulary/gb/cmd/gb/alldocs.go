// DO NOT EDIT THIS FILE.
//go:generate gb help documentation

/*
gb, a project based build tool for the Go programming language.

Usage:

        gb command [arguments]

The commands are:

        build       build a package
        doc         show documentation for a package or symbol
        env         print project environment variables
        generate    generate Go files by processing source
        info        info returns information about this project
        list        list the packages named by the importpaths
        test        test packages

Use "gb help [command]" for more information about a command.

Additional help topics:

        plugin      plugin information
        project     gb project layout

Use "gb help [topic]" for more information about that topic.


Build a package

Usage:

        gb build [build flags] [packages]

Build compiles the packages named by the import paths, along with their
dependencies.

Flags:

	-f
		ignore cached packages if present, new packages built will overwrite
		any cached packages. This effectively disables incremental
		compilation.
	-F
		do not cache packages, cached packages will still be used for
		incremental compilation. -f -F is advised to disable the package
		caching system.
	-q
		decreases verbosity, effectively raising the output level to ERROR.
		In a successful build, no output will be displayed.
	-P
		The number of build jobs to run in parallel, including test execution.
		By default this is the number of CPUs visible to gb.
	-R
		sets the base of the project root search path from the current working
		directory to the value supplied. Effectively gb changes working
		directory to this path before searching for the project root.
	-dotfile
		if provided, gb will output a dot formatted file of the build steps to
		be performed.
	-ldflags 'flag list'
		arguments to pass on each linker invocation.
	-gcflags 'arg list'
		arguments to pass on each compile invocation.
        -race
                enable data race detection.
                Supported only on linux/amd64, freebsd/amd64, darwin/amd64 and windows/amd64.
	-tags 'tag list'
		additional build tags.

The list flags accept a space-separated list of strings. To embed spaces in an
element in the list, surround it with either single or double quotes.

For more about where packages and binaries are installed, run 'gb help project'.


Show documentation for a package or symbol

Usage:

        gb doc <pkg> <sym>[.<method>]

Doc shows documentation for a package or symbol.

See 'go help doc'.


Print project environment variables

Usage:

        gb env [var ...]

Env prints project environment variables. If one or more variable names is
given as arguments, env prints the value of each named variable on its own line.


Generate Go files by processing source

Usage:

        gb generate [-run regexp] [file.go... | packages]

Generate runs commands described by directives within existing files.

Those commands can run any process, but the intent is to create or update Go
source files, for instance by running yacc.

See 'go help generate'.


Info returns information about this project

Usage:

        gb info [var ...]

info prints gb environment information.

Values:

	GB_PROJECT_DIR
		The root of the gb project.
	GB_SRC_PATH
		The list of gb project source directories.
	GB_PKG_DIR
		The path of the gb project's package cache.
	GB_BIN_SUFFIX
		The suffix applied any binary written to $GB_PROJECT_DIR/bin
	GB_GOROOT
		The value of runtime.GOROOT for the Go version that built this copy of gb.

info returns 0 if the project is well formed, and non zero otherwise.
If one or more variable names is given as arguments, info prints the
value of each named variable on its own line.


List the packages named by the importpaths

Usage:

        gb list [-s] [-f format] [-json] [packages]

List lists packages imported by the project.

The default output shows the package import paths:

	% gb list github.com/constabulary/...
	github.com/constabulary/gb
	github.com/constabulary/gb/cmd
	github.com/constabulary/gb/cmd/gb
	github.com/constabulary/gb/cmd/gb-env
	github.com/constabulary/gb/cmd/gb-list

Flags:
	-f
		alternate format for the list, using the syntax of package template.
		The default output is equivalent to -f '{{.ImportPath}}'. The struct
		being passed to the template is currently an instance of gb.Package.
		This structure is under active development and it's contents are not
		guaranteed to be stable.
	-s
		read format template from STDIN.
	-json
		prints output in structured JSON format. WARNING: gb.Package
		structure is not stable and will change in the future!


Plugin information

gb supports git style plugins.

A gb plugin is anything in the $PATH with the prefix gb-. In other words
gb-something, becomes gb something.

gb plugins are executed from the parent gb process with the environment
variable, GB_PROJECT_DIR set to the root of the current project.

gb plugins can be executed directly but this is rarely useful, so authors
should attempt to diagnose this by looking for the presence of the
GB_PROJECT_DIR environment key.


Gb project layout

A gb project is defined as any directory that contains a src/ subdirectory.
gb automatically detects the root of the project by looking at the current
working directory and walking backwards until it finds a directory that
contains a src/ subdirectory.

In the event you wish to override this auto detection mechanism, the -R flag
can be used to supply a project root.

See http://getgb.io/docs/project for details


Test packages

Usage:

        gb test [build flags] -v [packages] [flags for test binary]

Test automates testing the packages named by the import paths.

'gb test' recompiles each package along with any files with names matching
the file pattern "*_test.go".

Flags:

        -v
                print output from test subprocess.


*/
package main
def wire(**kwargs):
    native.genrule(
        outs = ["wire_gen.go"],

        # our command to run, we need to use execpath to get the path to the binary
        # we saw in the original go:generate directive, that it takes genql.yaml as its first argument
        # so by using $(location ...), we ask Bazel to get its path.
        # cmd = "echo HEEEEEEEEEEREEEEEE; echo $(location wire.go); $(execpath @com_github_google_wire//:wire) help",
        cmd = """\
env GOCACHE=$$(mktemp -d)\
    $(execpath @com_github_google_wire//cmd/wire:wire) $(location wire.go)
""",

        # we need to inform Bazel, that we need this binary for our cmd, otherwise, it won't find it.
        tools = [
            "@com_github_google_wire//cmd/wire:wire",
        ],
        **kwargs
    )

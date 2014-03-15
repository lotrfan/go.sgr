package sgr

var blockAlises = make(map[string]string)

func AddAlias(name, block string) {
    blockAlises[name] = block
}

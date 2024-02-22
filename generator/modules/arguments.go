package modules

type argument struct {
	name string
	typ  string
}

func makeLineArgumentsTemplate(args ...argument) string {
	if len(args) == 0 {
		return ""
	}

	var argsTMP string

	if hasContextArgument(args) {
		argsTMP += "ctx context.Context, "

		for i, arg := range args {
			if arg.typ == "context.Context" {
				args = append(args[:i], args[i+1:]...)
				break
			}
		}
	}

	groupedArgs := groupArgumentsByType(args)

	for typ, arguments := range groupedArgs {
		for _, arg := range arguments {
			argsTMP += arg.name + ", "
		}

		argsTMP = argsTMP[:len(argsTMP)-2] + " " + typ + ", "
	}

	for i := len(argsTMP) - 1; i >= 0; i-- {
		if argsTMP[i] != ' ' && argsTMP[i] != ',' {
			break
		}

		argsTMP = argsTMP[:i]
	}

	return argsTMP
}

func makeWithNewLinesArguments(args ...argument) string {
	if len(args) == 0 {
		return ""
	}

	var argsTMP string

	if hasContextArgument(args) {
		argsTMP += "\n\tctx context.Context,"

		for i, arg := range args {
			if arg.typ == "context.Context" {
				args = append(args[:i], args[i+1:]...)
				break
			}
		}
	}

	groupedArgs := groupArgumentsByType(args)

	for typ, arguments := range groupedArgs {
		argsTMP += "\n\t"

		for _, arg := range arguments {
			argsTMP += arg.name + ", "
		}

		argsTMP = argsTMP[:len(argsTMP)-2] + " " + typ + ","
	}

	return argsTMP + "\n"
}

func hasContextArgument(args []argument) bool {
	for _, arg := range args {
		if arg.typ == "context.Context" {
			return true
		}
	}

	return false
}

func groupArgumentsByType(args []argument) map[string][]argument {
	grouped := make(map[string][]argument)

	for _, arg := range args {
		grouped[arg.typ] = append(grouped[arg.typ], arg)
	}

	return grouped
}

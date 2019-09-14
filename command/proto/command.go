package proto


type ICommand interface {
	DoCommand(cs string) ([]byte, error)
}

type ICommandItem interface {
	Run(opt string)([]byte, error)
}
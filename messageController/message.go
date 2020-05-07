package messageController

type TaskManagerMessage struct {
	Add    addMessage
	Help   helpMessage
	Remove removeMessage
	List   listMessage
}

type addMessage struct {
	Err      addError
	Template string
}

type addError struct {
	NotEnoughArgs      string
	InvalidChannel     string
	InvalidDatePatarn  string
	InvalidSubjectName string
	DuplicateName      string
}

type helpMessage struct {
	General  string
	Subjects subjectsTemplate
}

type subjectsTemplate struct {
	EachSubject string
	EachCourse  string
}

type listMessage struct {
	Err      listError
	Template string
}

type listError struct {
	NotCheckTaskChannel  string
	SubjectIsNotInCourse string
	GetValueError        string
	TaskNotFound         string
}

type removeMessage struct {
	Success       string
	NotEnoughArgs string
	TaskNotFound  string
}

func CreateTaskManagerMessage() TaskManagerMessage {
	var messages TaskManagerMessage

	addMessages := createTaskAddMessage()

	return messages
}

func createTaskAddMessage() addMessage {
	addMessages := addMessage{}
	addMessages.Template = "```name: %s\nlimit: %d/%d\nsubject: %s```\nで新しい課題を作成しました。"
	addMessages.Err = createAddErrorMessage()

	return addMessages
}

func createAddErrorMessage() addError {
	addErrors := addError{}
	addErrors.NotEnoughArgs = "引数が足りません\n最低でも2個は必要です"
	addErrors.InvalidChannel = "課題を登録する際は<#%s>で行ってください"
	addErrors.InvalidDatePatarn = "日付の指定は n/m で指定してください"
	addErrors.InvalidSubjectName = "データの作成に失敗しました\n有効な教科の名前を指定してください"
	addErrors.DuplicateName = "データの作成に失敗しました\n課題の名前の重複などが無いか確認してください"

	return addErrors
}

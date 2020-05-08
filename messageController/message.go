package messageController

// GeneralError is Error in root of taskManager
type GeneralError struct {
	NotSetupped string
}

// AddError is Error while task add of taskManager
type AddError struct {
	NotEnoughArgs      string
	InvalidChannel     string
	InvalidDatePatarn  string
	InvalidSubjectName string
	DuplicateName      string
}

// HelpMessage is taskManager help message's template
type HelpMessage struct {
	General  string
	Subjects subjectsTemplate
}

type subjectsTemplate struct {
	EachSubject string
	EachCourse  string
}

// ListError is error while show taskList of taskManager
type ListError struct {
	NotCheckTaskChannel  string
	SubjectIsNotInCourse string
	GetValueError        string
	TaskNotFound         string
}

// RemoveError is a error while removeTask of taskManager
type RemoveError struct {
	NotEnoughArgs string
	TaskNotFound  string
}

// CreateTaskAddMessage is a function that creates a message when taskAdd succeeds.
func CreateTaskAddMessage() (addMessageTemplate string) {
	addMessageTemplate = "```name: %s\nlimit: %d/%d\nsubject: %s```\nで新しい課題を作成しました。"
	return
}

// CreateAddErrorMessage is a function that creates error messages while taskAdd.
func CreateAddErrorMessage() (addErrors AddError) {
	addErrors.NotEnoughArgs = "引数が足りません\n最低でも2個は必要です"
	addErrors.InvalidChannel = "課題を登録する際は<#%s>で行ってください"
	addErrors.InvalidDatePatarn = "日付の指定は n/m で指定してください"
	addErrors.InvalidSubjectName = "データの作成に失敗しました\n有効な教科の名前を指定してください"
	addErrors.DuplicateName = "データの作成に失敗しました\n課題の名前の重複などが無いか確認してください"
	return
}

// CreateTaskHelpMessage is a function that creates a message when execute taskHelp
func CreateTaskHelpMessage() (helpMessages HelpMessage) {
	helpMessages.General = "***課題管理BOT***\n```!task add <task> <limit> <subject>```\ntask: 課題名\nlimit: 締め切り(初期値=翌日)\nsubject: 教科(初期値='')\n教科は省略できる\n```!task list <subject>```\n課題一覧を表示します\n<subject>を指定すると教科ごとの絞り込みが可能です\n```!task remove <task>```\n課題を課題名から検索して削除します```!task help (subject)```\n使い方を表示します\nsubjectを付けると利用可能な教科を表示します"
	helpMessages.Subjects = subjectsTemplate{
		EachSubject: ", %s",
		EachCourse:  "%s\n```%s```\n",
	}
	return
}

// CreateTaskListMessageTemplate is a function that creates a template when show task list
func CreateTaskListMessageTemplate() (listMessageTemplate string) {
	listMessageTemplate = "```task: %s\nlimit: %s\nsubject: %s```"
	return
}

// CreateListErrorMessage is a function that creates error messages while show task list
func CreateListErrorMessage() (listErr ListError) {
	listErr.NotCheckTaskChannel = "このチャンネルは課題確認用に設定されていません"
	listErr.SubjectIsNotInCourse = "指定された教科は現在のチャンネルの系に存在しません"
	listErr.GetValueError = "値の取り出しでエラーが発生しました"
	listErr.TaskNotFound = "このチャンネル向けに作成された課題はありません"
	return
}

// CreateTaskRemoveMessage is a function that creates messsages when removeTask succeeds
func CreateTaskRemoveMessage() (removeSuccess string) {
	removeSuccess = "%sを削除しました"
	return
}

// CreateRemoveErrorMessage is a function that creates error messages while execute removeTask
func CreateRemoveErrorMessage() (removeErrors RemoveError) {
	removeErrors.NotEnoughArgs = "引数が足りません\n削除する課題の名前を指定してください"
	removeErrors.TaskNotFound = "指定された名前の課題は存在しません"
	return
}

// CreateGeneralErrorMessage is a function that creates error messages while root of taskManager
func CreateGeneralErrorMessage() (generalErrors GeneralError) {
	generalErrors.NotSetupped = "授業一覧が登録されていません"
	return
}

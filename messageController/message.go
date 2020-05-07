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
	Success string
	Err     removeError
}

type removeError struct {
	NotEnoughArgs string
	TaskNotFound  string
}

func CreateTaskManagerMessage() TaskManagerMessage {
	messages := TaskManagerMessage{}

	messages.Add = createTaskAddMessage()
	messages.Help = createTaskHelpMessage()
	messages.List = createTaskListMessage()
	messages.Remove = createTaskRemoveMessage()

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

func createTaskHelpMessage() helpMessage {
	helpMessages := helpMessage{}

	helpMessages.General = "***課題管理BOT***\n```!task add <task> <limit> <subject>```\ntask: 課題名\nlimit: 締め切り(初期値=翌日)\nsubject: 教科(初期値='')\n教科は省略できる\n```!task list <subject>```\n課題一覧を表示します\n<subject>を指定すると教科ごとの絞り込みが可能です\n```!task remove <task>```\n課題を課題名から検索して削除します```!task help (subject)```\n使い方を表示します\nsubjectを付けると利用可能な教科を表示します"
	helpMessages.Subjects = subjectsTemplate{
		EachSubject: ", %s",
		EachCourse:  "%s\n```%s```\n",
	}

	return helpMessages
}

func createTaskListMessage() listMessage {
	listMessages := listMessage{}

	listMessages.Err = createListErrorMessage()
	listMessages.Template = "```task: %s\nlimit: %s\nsubject: %s```"

	return listMessages
}

func createListErrorMessage() listError {
	listErr := listError{}

	listErr.NotCheckTaskChannel = "このチャンネルは課題確認用に設定されていません"
	listErr.SubjectIsNotInCourse = "指定された教科は現在のチャンネルの系に存在しません"
	listErr.GetValueError = "値の取り出しでエラーが発生しました"
	listErr.TaskNotFound = "このチャンネル向けに作成された課題はありません"

	return listErr
}

func createTaskRemoveMessage() removeMessage {
	removeMessages := removeMessage{}

	removeMessages.Success = "%sを削除しました"
	removeMessages.Err = removeError{
		NotEnoughArgs: "引数が足りません\n削除する課題の名前を指定してください",
		TaskNotFound:  "指定された名前の課題は存在しません",
	}

	return removeMessages
}

package repository

import (
	"testing"

	"github.com/lirlia/go-mypj/todo/model/task"
)

func TestAddTask(t *testing.T) {
	rep := New()
	rep.Add(task.Task{
		Title: "new task",
	})

	if len(rep.(*instance).tasks) != 3 {
		t.Errorf("タスクが追加されていること %d", len(rep.(*instance).tasks))
	} else {
		addedTask := rep.(*instance).tasks[2]
		if addedTask.Title != "new task" {
			t.Errorf("追加したタスクが末尾に追加されていること %s", addedTask.Title)

			if addedTask.ID <= 2 {
				t.Errorf("タスクに新しいIDが振られること %d", addedTask.ID)
			}
		}
	}
}

func TestDone(t *testing.T) {
	rep := New()
	rep.Done(0)

	if !rep.(*instance).tasks[0].Done {
		t.Errorf("タスクが完了していること")
	} else {
		rep.Done(1)
		if !rep.(*instance).tasks[1].Done {
			t.Errorf("完了済みのタスクをDoneしてもDone状態であること")
		}
	}
}

func TestDelete(t *testing.T) {
	rep := New()
	if len(rep.List()) != 2 {
		t.Errorf("タスクの数が2であること %d", len(rep.(*instance).tasks))
	}
}

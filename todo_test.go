package todocli_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Franklynoble/todocli"
)

//TestAcc test th Add method of the List type

func TestAdd(t *testing.T) {

	l := todocli.List{}

	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
	fmt.Printf("%+v\n", l[0].Task)
}

//TestComplete tests the Complete method of the List type
func TestComplete(t *testing.T) {

	l := todocli.List{}

	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New Task should Not be Completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed")
	}

}

func TestDelete(t *testing.T) {

	l := todocli.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, v := range tasks {

		l.Add(v)
	}
	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Task)
	}
	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead,", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)
	}

}

//TestSaveGet tests the save and get methods of he List type

func TestSaveGet(t *testing.T) {
	l1 := todocli.List{}
	l2 := todocli.List{}

	taskName := "New Task"

	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
	}

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file %s", err)
	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving to file: %s", err)

	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("error getting list from file %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Fatalf("error saving List to file: %s", err)
	}
}

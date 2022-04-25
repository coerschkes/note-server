package api

import (
	"co/note-server/src/main/adapter/api"
	"co/note-server/src/main/adapter/persistence/db"
	"co/note-server/src/main/config"
	"co/note-server/src/main/domain/model"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var controller api.NoteController
var repository model.NoteRepository
var baseUrl string

func TestMain(m *testing.M) {
	repository = db.MakeInMemoryNoteRepository()
	controller = api.NewNoteController(repository)
	baseUrl = "http://localhost:" + config.MakeConfigProvider().GetProperty("server.port")
	go controller.InitServer()
	os.Exit(m.Run())
}

func Test__post_note_should_create_new_note_in_the_repository(t *testing.T) {
	defer cleanup()
	id, title, content := "1", "Test", "Test content"

	resp := doRequest(http.MethodPost, "/notes", buildNoteJson(id, title, content))
	repoGet, _ := repository.GetById(id)
	assert.Equal(t, "201 Created", resp.Status)
	assert.Equal(t, id, repoGet.ID)
	assert.Equal(t, title, repoGet.Title)
	assert.Equal(t, content, repoGet.Content)
}

func Test__post_note_should_return_BadRequest_if_note_is_already_present(t *testing.T) {
	defer cleanup()
	id, title, content := "1", "Test", "Test content"
	repository.Add(model.Note{ID: id, Title: title, Content: content})

	resp := doRequest(http.MethodPost, "/notes", buildNoteJson(id, title, content))

	assert.Equal(t, "400 Bad Request", resp.Status)
	assert.Equal(t, readResponseBody(resp), "{\n    \"message\": \"Note with id '"+id+"' already exists.\"\n}")
}

func Test__delete_note_should_delete_note_from_the_repository(t *testing.T) {
	defer cleanup()
	id, title, content := "1", "Test", "Test content"
	repository.Add(model.Note{ID: id, Title: title, Content: content})

	resp := doRequest(http.MethodDelete, "/notes/"+id, "")

	_, err := repository.GetById(id)
	assert.Equal(t, "200 OK", resp.Status)
	assert.True(t, err != nil)
}

func Test__delete_note_should_return_BadRequest_if_note_is_not_present(t *testing.T) {
	defer cleanup()
	id := "1"

	resp := doRequest(http.MethodDelete, "/notes/"+id, "")

	assert.Equal(t, "400 Bad Request", resp.Status)
	assert.Equal(t, readResponseBody(resp), "{\n    \"message\": \"Note with id '"+id+"' not found\"\n}")
}

func Test__get_by_id_should_return_note_from_repository(t *testing.T) {
	defer cleanup()
	id, title, content := "1", "Test", "Test content"
	repository.Add(model.Note{ID: id, Title: title, Content: content})

	resp := doRequest(http.MethodGet, "/notes/"+id, buildNoteJson(id, title, content))

	body := readResponseBody(resp)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, buildNoteJson(id, title, content), body)
}

func Test__get_by_id_should_return_BadRequest_if_note_is_not_present(t *testing.T) {
	defer cleanup()
	id := "1"

	resp := doRequest(http.MethodGet, "/notes/"+id, "")

	assert.Equal(t, "400 Bad Request", resp.Status)
	assert.Equal(t, readResponseBody(resp), "{\n    \"message\": \"Note with id '"+id+"' not found\"\n}")
}

func Test__get_all_should_return_all_notes_in_repository(t *testing.T) {
	defer cleanup()
	id1, title1, content1 := "1", "Test1", "Test1 content"
	id2, title2, content2 := "2", "Test2", "Test2 content"
	repository.Add(model.Note{ID: id1, Title: title1, Content: content1})
	repository.Add(model.Note{ID: id2, Title: title2, Content: content2})

	resp := doRequest(http.MethodGet, "/notes", "")

	body := readResponseBody(resp)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Contains(t, body, id1)
	assert.Contains(t, body, title1)
	assert.Contains(t, body, content1)
	assert.Contains(t, body, id2)
	assert.Contains(t, body, title2)
	assert.Contains(t, body, content2)
}

func Test__get_all_should_return_empty_array_if_no_notes_exist(t *testing.T) {
	defer cleanup()

	resp := doRequest(http.MethodGet, "/notes", "")

	body := readResponseBody(resp)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, "[]", body)
}

func doRequest(method string, path string, requestBody string) *http.Response {
	req, err := http.NewRequest(method, baseUrl+path, strings.NewReader(requestBody))
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	return resp
}

func buildNoteJson(id string, title string, content string) string {
	return "{\n    \"id\": \"" + id + "\",\n    \"title\": \"" + title + "\",\n    \"content\": \"" + content + "\"\n}"
}

func readResponseBody(response *http.Response) string {
	b, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func cleanup() {
	if all, err := repository.GetAll(); err != nil {
		panic(err)
	} else {
		for _, note := range all {
			if err := repository.DeleteById(note.ID); err != nil {
				panic(err)
			}
		}
	}
}

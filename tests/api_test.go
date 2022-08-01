package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"rollic/pkg/routes"
	"strings"
	"testing"
)

func Test_AddUser(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user", routes.Add)
	ts := httptest.NewServer(r)
	defer ts.Close()
	body := `{"email":"admin@gmail.com","password":"123456"}`
	req, err := http.NewRequest(http.MethodPut, ts.URL+"/user", strings.NewReader(body))
	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("got error on api request %s", err)
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		t.Errorf("Got unexpected resonse %s", resp.Status)
	}
}

func Test_Login(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/login", routes.Login)
	ts := httptest.NewServer(r)
	defer ts.Close()
	body := `{"email":"admin@gmail.com","password":"123456"}`
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/user/login", strings.NewReader(body))
	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("got error on api request %s", err)
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		t.Errorf("Got unexpected resonse %s", resp.Status)
	}
}

func Test_DeleteUser(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", routes.Delete)
	ts := httptest.NewServer(r)
	defer ts.Close()
	req, err := http.NewRequest(http.MethodDelete, ts.URL+"/user/62c576ea438084dd21bd0e31", nil)
	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("got error on api request %s", err)
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		t.Errorf("Got unexpected resonse %s", resp.Status)
	}
}

func Test_GetUser(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", routes.Get)
	ts := httptest.NewServer(r)
	defer ts.Close()
	req, err := http.NewRequest(http.MethodGet, ts.URL+"/user/62c576ea438084dd21bd0e32", nil)
	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("got error on api request %s", err)
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		t.Errorf("Got unexpected resonse %s", resp.Status)
	}
}

func Test_GetAll(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users/all", routes.All)
	ts := httptest.NewServer(r)
	defer ts.Close()
	req, err := http.NewRequest(http.MethodGet, ts.URL+"/users/all", nil)
	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("got error on api request %s", err)
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		t.Errorf("Got unexpected resonse %s", resp.Status)
	}
}

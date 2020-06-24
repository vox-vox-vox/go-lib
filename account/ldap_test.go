package main

import (
	"testing"
    _ "errors"
    "fmt"
    "github.com/go-ldap/ldap/v3"
)

func TestUserPass(t *testing.T){
    ldapServer := "ldap://192.168.1.40:389"
    departmentDN := "开发部"
    departmentPasswd := "开发部password"
    personDN := "OU=All Users,DC=ahuigo,DC=com"
    personAttribute := []string{ "accountName"}
    username := "username"
    password := "user_password"

	l, err := ldap.DialURL(ldapServer)
    if err != nil {
		t.Fatal(err)
    }
    // bind  BaseDN
    err = l.Bind(departmentDN, departmentPasswd)
    if err != nil {
		t.Fatal(err)
    }
    // search person
    searchRequest := ldap.NewSearchRequest(
		personDN,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
        fmt.Sprintf("(&(objectclass=person)(CN=%s))", username),
		personAttribute,
		nil,
    )
	sr, err := l.Search(searchRequest)
	if err != nil {
		t.Error(err)
	}

    if len(sr.Entries) == 0 {
        t.Error("user not exists")
    }

    entry_dn := sr.Entries[0].DN
    t.Logf("Search: %s -> num of entries = %d", searchRequest.Filter, len(sr.Entries))
    t.Logf("%#v",  sr.Entries[0])
    t.Logf("dn:%s",  sr.Entries[0].DN)

    // verify password
    err = l.Bind(entry_dn, password)
    if err != nil {
		t.Fatal("invalid password")
		t.Fatal(err)
    }

}


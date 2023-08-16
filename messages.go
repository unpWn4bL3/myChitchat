package main

import "github.com/nicksnyder/go-i18n/v2/i18n"

var messsages = []i18n.Message{
	i18n.Message{
		ID:          "thread_not_found",
		Description: "Thread not exsists in db",
		Other:       "Cannot read thread",
	},
}

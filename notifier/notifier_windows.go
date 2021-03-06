package notifier

import "gopkg.in/toast.v1"

func Notify(appName string, title string, message string, icon string) error {
	notification := toast.Notification{
		AppID: appName,
		Title: title,
		Message: message,
		Icon: icon,
	};

	err := notification.Push();
	return err;
}
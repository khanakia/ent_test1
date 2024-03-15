// Code generated from Pkl module `AppConfig`. DO NOT EDIT.
package appconfig

type Database struct {
	// The username for this database.
	Username string `pkl:"username"`

	// The password for this database.
	Password string `pkl:"password"`

	// The remote host for this database.
	Host string `pkl:"host"`

	// The remote port for this database.
	Port uint16 `pkl:"port"`

	// The name of the database.
	DbName string `pkl:"dbName"`

	Sslmode string `pkl:"sslmode"`
}
